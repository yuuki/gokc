package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/yuuki1/gokc/log"
	"github.com/yuuki1/gokc/parser"
)

func setDebugOutputLevel() {
	for _, f := range os.Args {
		if f == "-D" || f == "--debug" || f == "-debug" {
			log.IsDebug = true
		}
	}

	debugEnv := os.Getenv("GOKC_DEBUG")
	if debugEnv != "" {
		showDebug, err := strconv.ParseBool(debugEnv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing boolean value from GRABENI_DEBUG: %s\n", err)
			os.Exit(1)
		}
		log.IsDebug = showDebug
	}
}

func init() {
	setDebugOutputLevel()
}

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

	fmt.Printf("gokc: --> Parsing %s ...\n", filepath)

	parser.Parse(file)

	fmt.Println("gokc: the configuration file %s syntax is ok", filepath)
}
