package service

import (
	"errors"
	"io"
	"log"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

type ShareService struct {
	src             string
	dest            string
	total           int
}

func New(src, dest string, args ...int) *ShareService {
	total := 1024
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

	var copiedSize int64 = 0
	if destStats, err := os.Stat(s.dest); err == nil {
		copiedSize = destStats.Size()
		if copiedSize >= size {
			log.Printf("✅ File already fully copied!, skipping %s copy", s.src)
			return
		}
	}

	destFile, err := os.OpenFile(s.dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	_, err = srcFile.Seek(copiedSize, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	_, err = destFile.Seek(copiedSize, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	bar := progressbar.DefaultBytes(
		size-copiedSize,
		"Copying...",
	)

	for {
		n, err := srcFile.Read(buffer)
		if errors.Is(err, io.EOF) {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		destFile.Write(buffer[:n])
		bar.Add(n)
		time.Sleep(time.Millisecond * 1)
	}
}
