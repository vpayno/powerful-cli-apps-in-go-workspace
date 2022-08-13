// Package appwc is the module with the cli logic for the wc main application.
package appwc

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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

	counts := getCounts(os.Stdin, conf)

	showCount(counts, conf)
}

// Usage prints the command-line usage help message.
func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
	flagSet.PrintDefaults()
}

func setup() (config, error) {

	byteFlag := flagSet.Bool("b", false, "count bytes")
	charFlag := flagSet.Bool("r", false, "count characters")
	lineFlag := flagSet.Bool("l", false, "count newlines")
	wordFlag := flagSet.Bool("w", false, "count words")
	verboseFlag := flagSet.Bool("v", false, "verbose mode")
	versionFlag := flagSet.Bool("V", false, "show the app version")

	flagSet.Usage = Usage

	err := flagSet.Parse(os.Args[1:])

	if err != nil {
		return config{}, err
	}

	conf := config{
		verboseMode: *verboseFlag,
		versionMode: *versionFlag,
	}

	usingDefaults := true

	var byteMode bool
	var charMode bool
	var lineMode bool
	var wordMode bool

	// order for flag overrides: newline, word, character, byte.
	switch {
	case *lineFlag:
		usingDefaults = false

		byteMode = false
		charMode = false
		lineMode = true
		wordMode = false

	case *wordFlag:
		usingDefaults = false

		byteMode = false
		charMode = false
		lineMode = false
		wordMode = true

	case *charFlag:
		usingDefaults = false

		byteMode = false
		charMode = true
		lineMode = false
		wordMode = false

	case *byteFlag:
		usingDefaults = false

		byteMode = true
		charMode = false
		lineMode = false
		wordMode = false
	}

	if usingDefaults {
		byteMode = true
		lineMode = true
		charMode = false
		wordMode = true
	}

	conf.modes = map[string]bool{
		"byte": byteMode,
		"char": charMode,
		"line": lineMode,
		"word": wordMode,
	}

	return conf, nil
}

func getCounts(r io.Reader, conf config) results {
	counts := results{
		"byte": 0,
		"char": 0,
		"line": 0,
		"word": 0,
	}

	reader := bufio.NewReader(r)

	for {
		text, err := reader.ReadString('\n')

		if conf.modes["byte"] {
			counts["byte"] += len(text)
		}

		if conf.modes["char"] {
			counts["char"] += utf8.RuneCountInString(text)
		}

		if conf.modes["word"] {
			counts["word"] += len(strings.Fields(text))
		}

		if conf.modes["line"] {
			counts["line"]++

			if err == io.EOF {
				counts["line"]--
			}
		}

		if err != nil {
			break
		}
	}

	return counts
}

func showCount(counts results, conf config) {
	first := true
	fieldSize := "7"

	var modeCount int
	for _, v := range conf.modes {
		if v {
			modeCount++
		}
	}

	if modeCount == 1 {
		fieldSize = "0"
	}

	for _, mode := range printOrder {
		if !conf.modes[mode] {
			continue
		}

		if !first {
			fieldSize = "8"
		}

		first = false

		var label string

		if conf.verboseMode {
			label = " (" + mode + ")"
		}

		fmt.Printf("%"+fieldSize+"d%s", counts[mode], label)
	}

	fmt.Println()
}
