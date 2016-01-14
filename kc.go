package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yuuki1/gokc/parser"
)

func main() {
	var filepath string
	flag.StringVar(&filepath, "f", "", "keepalived.conf file path")
	flag.Parse()

	if filepath == "" {
		fmt.Println("filepath required")
		os.Exit(1)
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("--> Parsing %s...\n", filepath)

	parser.Parse(file)
}
