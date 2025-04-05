package main

import (
	"fmt"
	"os"

	"github.com/a-ZINC/sharer/server"
)

// "github.com/a-ZINC/sharer/config"
// "github.com/a-ZINC/sharer/service"

func main() {
	// srcPath, dstPath := config.Load()
	// share := service.New(srcPath, dstPath)
	// share.Read();

	err := server.Run()
	if err != nil {
		fmt.Printf("error in connection")
		os.Exit(1)
	}
}
