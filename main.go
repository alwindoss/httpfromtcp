package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	var currentLine string
	var nextLine string
	for {
		buff := make([]byte, 8)
		_, err := f.Read(buff)
		if err == io.EOF {
			fmt.Printf("read: %s\n", currentLine)
			break
		}
		if index := bytes.IndexByte(buff, '\n'); index == -1 {
			currentLine = currentLine + string(buff)
		} else {
			parts := bytes.Split(buff, []byte("\n"))
			currentLine = currentLine + string(parts[0])
			nextLine = string(parts[1])
			fmt.Printf("read: %s\n", currentLine)
			currentLine = ""
			currentLine = nextLine
			nextLine = ""
		}
	}
}
