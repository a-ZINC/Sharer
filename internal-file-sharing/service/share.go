package service

import (
	"errors"
	"io"
	"log"
	"os"
)

type ShareService struct {
	src             string
	dest            string
	total           int
	LoadingProgress int
}

func New(src, dest string, args ...int) *ShareService {
	total := 1024 * 1024
	if len(args) > 0 {
		total = args[0]
	}
	return &ShareService{
		src:   src,
		dest:  dest,
		total: total,
	}
}

func (s *ShareService) Read() {
	srcFile, err := os.Open(s.src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	buffer := make([]byte, s.total)

	stats, err := srcFile.Stat()

	if err != nil {
		log.Fatal(err)
	}

	size := stats.Size()
	if size == 0 {
		log.Fatal("Source file is empty")
	}

	destFile, err := os.Create(s.dest)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	for {
		n, err := srcFile.Read(buffer)
		if errors.Is(err, io.EOF) {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		destFile.Write(buffer[:n])
		s.LoadingProgress += n
		log.Printf("Loading progress: %d/%d", s.LoadingProgress, size)
	}
}
