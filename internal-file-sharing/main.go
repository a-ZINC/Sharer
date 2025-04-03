package main

import (
	"github.com/a-ZINC/sharer/config"
	"github.com/a-ZINC/sharer/service"
)

func main() {
	srcPath, dstPath := config.Load()
	share := service.New(srcPath, dstPath)
	share.Read();
	

}
