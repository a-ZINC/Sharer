package service

import (
	"log"
	"os"
)

func Create(src string, n int) {
	file, err := os.Create(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < n; i++ {

		file.WriteString(string(i))
	}
}
