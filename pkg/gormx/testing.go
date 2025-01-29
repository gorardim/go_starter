package gormx

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"app/pkg/randx"
	"github.com/stretchr/testify/assert"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewTestDb create a test database
// schemaPath: schema file path
// fixturePath is the path of fixture files
func NewTestDb(t *testing.T) *gorm.DB {
	// 约定测试文件名称,不做默认值处理
	fixtures := []string{"./testdata/" + t.Name()}
	dbName := createDatabase(t)
	// 重新打开数据库,避免连接导致数据表不存在
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("root:123456@tcp(localhost:33061)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai", dbName)), &gorm.Config{
		Logger: newTestLog(logger.Silent),
	})

	// set sql mode
	assert.NoError(t, db.Exec("set sql_mode='';").Error)
	assert.NoError(t, err)
	// add clean up hook
	t.Cleanup(func() {
		if db.Raw(fmt.Sprintf("drop database %s", dbName)).Error != nil {
			t.Fatalf("drop database error: %v", err)
		}
		log.Println("drop database:", dbName)
	})

	// use database
	if db.Exec(fmt.Sprintf("use %s", dbName)).Error != nil {
		panic(fmt.Sprintf("use database error: %v", err))
	}

	// read schema file
	log.Printf("read schema file: %s", "./testdata/schema.sql")
	schemaSql, err := os.ReadFile("./testdata/schema.sql")
	if err != nil {
		panic(fmt.Sprintf("read schema file error: %v", err))
	}

	// import schema
	for _, sq := range strings.Split(string(schemaSql), ";") {
		sq = strings.TrimSpace(sq)
		if sq == "" {
			continue
		}
		if err = db.Exec(sq).Error; err != nil {
			t.Fatalf("import schema error: %v", err)
		}
	}

	// import fixtures
	for _, fixture := range fixtures {
		log.Printf("read fixture file: %s", fixture)
		// is dir
		if !isDir(fixture) {
			loadYmlToDB(db, fixture)
			continue
		}
		// read dir
		files, err := os.ReadDir(fixture)
		if err != nil {
			t.Fatalf("read dir error: %v", err)
		}
		// import fixtures
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			loadYmlToDB(db, filepath.Join(fixture, file.Name()))
		}
	}
	return db.Session(&gorm.Session{
		Logger: newTestLog(logger.Info),
	})
}

func createDatabase(t *testing.T) string {
	// open gorm db
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:33061)/mysql?charset=utf8&parseTime=True&loc=Asia%2FShanghai"), &gorm.Config{
		Logger: newTestLog(logger.Silent),
	})
	assert.NoError(t, err)

	// create database
	dbName := "test_" + fmt.Sprintf("%s_%d_%s", time.Now().Format("20060102150405"), time.Now().Nanosecond(), randx.Digit(6))
	log.Printf("create database: %s", dbName)
	if err = db.Exec(fmt.Sprintf("create database %s character set utf8mb4 collate utf8mb4_unicode_ci", dbName)).Error; err != nil {
		panic(fmt.Sprintf("create database error: %v", err))
	}
	sqlDb, err := db.DB()
	assert.NoError(t, err)
	_ = sqlDb.Close()
	return dbName
}

var l = log.New(os.Stdout, "\r\n", log.Llongfile)

type writer struct{}

func (w *writer) Printf(s string, i ...interface{}) {
	i[0] = getOutputFilename(i[0].(string))
	log.Printf(s+"\n\n", i...)
}

func getOutputFilename(f string) string {
	skip := 5
	lastFile := f
	for i := 0; i < 3; i++ {
		if strings.Contains(lastFile, "/pkg/gormx/") {
			_, file, line, ok := runtime.Caller(skip)
			if ok && !strings.Contains(file, "/pkg/gormx/") {
				return fmt.Sprintf("%s:%d", file, line)
			}
			lastFile = file
			skip++
		}
	}
	return lastFile
}

// newTestLog create a test logger
func newTestLog(level logger.LogLevel) logger.Interface {
	return logger.New(
		// log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		&writer{},
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  level,       // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
}

// loadYmlToDB load yml file to db
func loadYmlToDB(db *gorm.DB, fixture string) {
	// read fixture file
	fixtureContent, err := os.ReadFile(fixture)
	if err != nil {
		panic(fmt.Sprintf("read fixture file error: %v", err))
	}

	lines := make([]map[string]interface{}, 0)
	if err := yaml.Unmarshal(fixtureContent, &lines); err != nil {
		panic(fmt.Sprintf("parse fixture file error: %v", err))
	}
	// import fixture
	if err = db.Table(strings.TrimSuffix(filepath.Base(fixture), ".yml")).Create(lines).Error; err != nil {
		panic(fmt.Sprintf("import fixture error: %v", err))
	}
}

func isDir(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return info.IsDir()
}
