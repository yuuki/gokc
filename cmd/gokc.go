package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	log.Infof("gokc: --> Parsing %s ...\n", filepath)

	p := parser.NewParser(file)
	if err := p.Parse(); err != nil {
		var msgs []string
		for _, e := range p.Errors() {
			msgs = append(msgs, fmt.Sprintf("gokc: %s, line %d, pos %d", e.Message, e.Line, e.Column))
		}
		log.Errorf("%s", strings.Join(msgs, "\n"))
	}

	log.Infof("gokc: the configuration file %s syntax is ok", filepath)
}
