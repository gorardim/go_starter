package main

import (
	"flag"

	"app/cmd/xapi/internal/modules/api"
)

var scanDir string
var apiTitle string

func main() {
	flag.StringVar(&scanDir, "dir", "", "scan directory")
	flag.StringVar(&apiTitle, "api-title", "", "apiTitle")
	flag.Parse()
	api.Run(scanDir, apiTitle)
}
