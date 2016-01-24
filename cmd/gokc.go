package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/colorstring"

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
		log.Error("filepath required")
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Error(err)
	}

	p := parser.NewParser(file)
	if err := p.Parse(); err != nil {
		var msgs []string
		for _, e := range p.Errors() {
			msg := colorstring.Color(fmt.Sprintf("[white]%s:%d:%d: [red]%s[reset]", filepath, e.Line, e.Column, e.Message))
			msgs = append(msgs, msg)
		}
		log.Errorf("%s", strings.Join(msgs, "\n"))
	}

	log.Infof("gokc: the configuration file %s syntax is ok", filepath)
}
