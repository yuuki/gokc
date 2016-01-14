package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var filepath string
	flag.StringVar(&filepath, "f", "", "keepalived.conf file path")
	flag.Parse()

	if filepath == "" {
		fmt.Println("filepath required")
		os.Exit(1)
	}

	fmt.Println(filepath)
}
