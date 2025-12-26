package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		defer f.Close()
		data := make([]byte, 8)
		var str strings.Builder
		for {
			n, err := f.Read(data)
			if n > 0 {
				//fmt.Print(string(data[:n]))
				chunk := data[:n]
				for {
					i := bytes.IndexByte(chunk, '\n')
					if i == -1 {
						str.Write(chunk)
						break
					}
					str.Write(chunk[:i])
					chunk = chunk[i+1:]
				}

			}
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
		}
		out <- str.String()
	}()
	return out
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	var filedata string = <-getLinesChannel(file)
	fmt.Println(filedata)

}
