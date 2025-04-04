package main

import (
	"github.com/a-ZINC/sharer/config"
	"github.com/a-ZINC/sharer/service"
)

func main() {
	srcPath, dstPath := config.Load()
	// service.Create(srcPath, 10000000)
	share := service.New(srcPath, dstPath)
	share.Read();
}
