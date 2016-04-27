package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data/foo.txt")
	if err != nil {
		log.Fatal(err)
	}

	bf := bufio.NewReader(f)

	for {
		// reader.ReadLine does a buffered read up to a line terminator,
		// handles either /n or /r/n, and returns just the line without
		// the /r or /r/n.
		line, isPrefix, err := bf.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", f.Name())
		}

		fmt.Println(string(line))
	}
}
