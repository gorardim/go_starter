package logger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"app/pkg/utils"
)

var _ Logger = (*FileLogger)(nil)

type FileLogger struct {
	storagePath string
	w           *os.File
	old         *os.File
	mu          sync.Mutex
	appName     string
	currentDate string
	buf         *bytes.Buffer
}

func NewFileLogger(storagePath string, appName string) *FileLogger {
	// 检查目录是否存在
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		log.Fatalf("file log storage path not exist: %s", storagePath)
	}
	l := &FileLogger{
		storagePath: storagePath,
		appName:     appName,
		buf:         &bytes.Buffer{},
	}
	l.init()
	return l
}

func (f *FileLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	f.Output(ctx, 3, map[string]string{"msg": fmt.Sprintf(format, v...)})
}

func (f *FileLogger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	f.Output(ctx, 3, map[string]string{"msg": fmt.Sprintf(format, v...)})
	os.Exit(1)
}

func (f *FileLogger) PrintMap(ctx context.Context, m map[string]string) {
	f.Output(ctx, 3, m)
}

func (f *FileLogger) Output(ctx context.Context, skip int, m map[string]string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	var ok bool
	var file string
	var line int

	for i := 0; i < 5; i++ {
		_, file, line, ok = runtime.Caller(skip)
		if !ok {
			file = "???"
			line = 0
			break
		}
		// /pkg/gormx/
		if strings.Contains(file, "/pkg/gormx/") || strings.Contains(file, "gorm.io") || strings.HasSuffix(file, "component/pagenate.go") {
			skip++
			continue
		}
	}
	short := f.cutFilename(file)
	m["file"] = fmt.Sprintf("%s:%d", short, line)
	m["trace_id"] = TraceIdFromLogger(ctx)
	m["app_name"] = f.appName
	m["@timestamp"] = time.Now().Format("2006-01-02T15:04:05.000Z")
	encodeByte := utils.JsonEncodeByte(m)
	// 写入buffer
	f.buf.Write(encodeByte)
	f.buf.WriteByte('\n')
}

func (f *FileLogger) cutFilename(file string) string {
	if index := strings.Index(file, "internal/"); index >= 0 {
		return file[index+len("internal/"):]
	}
	if index := strings.Index(file, "/mod/"); index >= 0 {
		return file[index+len("/mod/"):]
	}
	// pkg 目录
	if index := strings.Index(file, "/pkg/"); index >= 0 {
		return file[index+len("/pkg/"):]
	}
	return file
}

func (f *FileLogger) init() {
	// 每天创建一个文件
	ticker := time.NewTicker(time.Minute * 1)
	f.checkAndCreateFile(time.Now())
	go func() {
		for v := range ticker.C {
			f.checkAndCreateFile(v)
		}
	}()

	// 没1s写一次
	ticker2 := time.NewTicker(time.Second)
	go func() {
		for range ticker2.C {
			if f.w == nil || f.buf.Len() == 0 {
				continue
			}
			f.mu.Lock()
			_, err := io.Copy(f.w, f.buf)
			if err != nil {
				log.Printf("write file log error: %s", err)
				log.Println(f.buf.String())
			}
			f.buf.Reset()
			f.mu.Unlock()
		}
	}()
}

func (f *FileLogger) checkAndCreateFile(nextTime time.Time) {
	date := nextTime.Format("20060102")
	if f.currentDate == date {
		_, err := os.Stat(f.w.Name())
		if err == nil {
			return
		}
		log.Printf("file log not exist: %s", f.w.Name())
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	file, err := f.createFile(date)
	if err != nil {
		log.Printf("create file log error: %s", err)
		return
	}
	f.currentDate = date
	if f.w != nil {
		f.old = f.w
	}
	f.w = file
	// close old file
	if f.old != nil {
		old := f.old
		// 一分钟后关闭
		time.AfterFunc(time.Minute, func() {
			_ = old.Close()
		})
	}
}

func (f *FileLogger) createFile(date string) (*os.File, error) {
	fileName := fmt.Sprintf("%s/%s_%s.log", f.storagePath, f.appName, date)
	// create file
	return os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
}
