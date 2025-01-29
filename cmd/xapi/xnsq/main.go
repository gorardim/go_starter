package main

import (
	"flag"

	"app/cmd/xapi/internal/modules/nsq"
)

var scanDir string

func main() {
	flag.StringVar(&scanDir, "dir", "", "scan directory")
	flag.Parse()
	if scanDir == "" {
		flag.Usage()
		return
	}
	nsq.Run(scanDir)
}
