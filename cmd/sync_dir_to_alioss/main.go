package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"app/pkg/alioss"
	"app/pkg/task"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var srcPath string
var dstPath string
var bucketName string

func main() {
	flag.StringVar(&srcPath, "src", "", "src path")
	flag.StringVar(&dstPath, "dst", "", "dst path")
	flag.StringVar(&bucketName, "bucket", "", "bucket name")
	flag.Parse()

	if srcPath == "" || dstPath == "" || bucketName == "" {
		flag.Usage()
		log.Fatalf("invalid args")
	}

	client, err := alioss.NewClient("oss-cn-hangzhou.aliyuncs.com", "", "")
	if err != nil {
		log.Fatalf("failed to new client: %v", err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Fatalf("failed to get bucket: %v", err)
	}
	if err := doSync(bucket, srcPath, dstPath); err != nil {
		log.Fatalf("failed to sync: %v", err)
	}
}

func doSync(bucket *oss.Bucket, srcPath, dstPath string) error {
	w := task.NewWorker(100)
	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("failed to walk dir: %v", err)
		}
		if info.IsDir() {
			return nil
		}
		w.Go(func() {
			relPath, err := filepath.Rel(srcPath, path)
			if err != nil {
				log.Fatalf("failed to get rel path: %v", err)
			}
			dst := strings.TrimPrefix(filepath.Join(dstPath, relPath), "/")
			if err := bucket.PutObjectFromFile(dst, path); err != nil {
				log.Fatalf("failed to put object: %v", err)
			}
			log.Printf("sync %s to %s", path, dst)
		})
		return nil
	})
	w.Wait()
	return err
}
