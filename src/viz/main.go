package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var inputFile string

// Parse command line options
func parse_command_line() {
	flag.StringVar(&inputFile, "file", "", "Use file as input source")
	flag.Parse()
}

// Parse input file and print to the output
func parse_file(input string) {
	f, err := os.Open(input)
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

func main() {
	parse_command_line()

	if inputFile != "" {
		parse_file(inputFile)
	}
}
