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
	flag.PrintDefaults()
}

func setup() (config, error) {
	fs := flag.NewFlagSet("cli", flag.ContinueOnError)

	byteFlag := fs.Bool("b", false, "count bytes")
	charFlag := fs.Bool("r", false, "count characters")
	lineFlag := fs.Bool("l", false, "count newlines")
	wordFlag := fs.Bool("w", false, "count words")
	verboseFlag := fs.Bool("v", false, "verbose mode")
	versionFlag := fs.Bool("V", false, "show the app version")

	fs.Usage = Usage

	err := fs.Parse(os.Args[1:])

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

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		text := scanner.Text() + "\n"

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
		}
	}

	return counts
}

func showCount(counts results, conf config) {
	first := true
	fieldSize := "0"

	if conf.verboseMode {
		fieldSize = "8"
	}

	for _, mode := range printOrder {
		if !conf.modes[mode] {
			continue
		}

		if !first {
			fieldSize = "9"
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
