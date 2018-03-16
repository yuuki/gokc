package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/mitchellh/colorstring"

	"github.com/yuuki/gokc/log"
	"github.com/yuuki/gokc/parser"
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
	var (
		filepath  string
		isVersion bool
		printJSON bool
	)
	flag.StringVar(&filepath, "f", "/etc/keepalived/keepalived.conf", "keepalived.conf file path")
	flag.BoolVar(&isVersion, "v", false, "print the version")
	flag.BoolVar(&log.IsVerbose, "V", false, "verbose log mode")
	flag.BoolVar(&printJSON, "json", false, "print configuration as json")
	flag.Parse()

	if isVersion {
		log.Infof("gokc version %s", Version)
		os.Exit(0)
	}

	if filepath == "" {
		log.Error("filepath required")
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Error(err)
	}
	defer file.Close()

	result, err := parser.Parse(file, filepath)
	if err != nil {
		if e, ok := err.(*parser.Error); ok {
			msg := colorstring.Color(fmt.Sprintf("[white]%s:%d:%d: [red]%s[reset]", e.Filename, e.Line, e.Column, e.Message))
			log.Error(msg)
		} else {
			log.Error(err)
		}
	}
	log.Debugf("%#v\n", result)

	if printJSON {
		j, err := json.Marshal(result)
		if err != nil {
			log.Error(err)
		}
		fmt.Println(string(j))
	} else {
		log.Infof("gokc: the configuration file %s syntax is ok", filepath)
	}
}
