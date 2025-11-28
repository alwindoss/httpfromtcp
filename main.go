package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	var currentLine string
	var nextLine string
	currentLineCh := make(chan string, 1)
	go func() {
		defer f.Close()
		defer close(currentLineCh)
		for {
			buff := make([]byte, 8)
			_, err := f.Read(buff)
			if err == io.EOF {
				// fmt.Printf("read: %s\n", currentLine)
				currentLineCh <- currentLine
				return
			}
			if index := bytes.IndexByte(buff, '\n'); index == -1 {
				currentLine = currentLine + string(buff)
			} else {
				parts := bytes.Split(buff, []byte("\n"))
				currentLine = currentLine + string(parts[0])
				nextLine = string(parts[1])
				// fmt.Printf("read: %s\n", currentLine)
				currentLineCh <- currentLine
				currentLine = ""
				currentLine = nextLine
				nextLine = ""
			}
		}
	}()
	return currentLineCh
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	currentLineCh := getLinesChannel(f)
	for cl := range currentLineCh {
		fmt.Printf("read: %s\n", string(cl))
	}
}
