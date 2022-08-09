// Package appwc is the module with the cli logic for the wc main application.
package appwc

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func showBanner() {
	fmt.Println(metadata.name + " Version " + metadata.version)
	fmt.Println()
}

// RunApp is called my the main function. It's basically the main function of the app.
func RunApp() {
	conf, err := setup()

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		return
	}

	if conf.versionMode {
		showVersion()
		return
	}

	if conf.verboseMode {
		showBanner()
	}

	count := getCount(os.Stdin, conf)

	showCount(count, conf)
}

func setup() (config, error) {
	byteMode := flag.Bool("b", defaults.lineMode, "count bytes instead of words")
	lineMode := flag.Bool("l", defaults.lineMode, "count lines instead of words")
	verboseMode := flag.Bool("v", false, "verbose mode")
	versionMode := flag.Bool("V", false, "show the app version")
	flag.Parse()

	conf := config{
		byteMode:    *byteMode,
		lineMode:    *lineMode,
		verboseMode: *verboseMode,
		versionMode: *versionMode,
	}

	if conf.byteMode && conf.lineMode {
		err := errors.New("-b (byte count mode) and -l (line count mode) can't be used at the same time")
		return config{}, err
	}

	return conf, nil
}

func getCount(r io.Reader, c config) int {
	scanner := bufio.NewScanner(r)

	var count int

	if !c.lineMode {
		scanner.Split(bufio.ScanWords)
	}

	for scanner.Scan() {
		if c.byteMode {
			count += utf8.RuneCountInString(scanner.Text())
		} else {
			count++
		}
	}

	return count
}

func showCount(n int, conf config) {
	var prompt string

	if conf.verboseMode {
		switch {
		case conf.byteMode:
			prompt = "byte count: "
		case conf.lineMode:
			prompt = "line count: "
		default:
			prompt = "word count: "
		}
	}

	fmt.Print(prompt)
	fmt.Printf("%d\n", n)
}
