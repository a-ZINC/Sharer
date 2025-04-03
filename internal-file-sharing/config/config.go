package config

import (
	"log"
	"os"
)

func Load() (string, string) {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a config file")
	}

	srcPath := os.Args[1]
	dstPath := os.Args[2]

	if srcPath == "" || dstPath == "" {
		log.Fatal("Please provide a source and destination path")
	}
	return srcPath, dstPath
}