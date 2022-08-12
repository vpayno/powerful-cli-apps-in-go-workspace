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

	lineMode := fs.Bool("l", false, "count newlines")
	wordMode := fs.Bool("w", false, "count words")
	runeMode := fs.Bool("r", false, "count runes/characters")
	byteMode := fs.Bool("b", false, "count bytes")
	verboseMode := fs.Bool("v", false, "verbose mode")
	versionMode := fs.Bool("V", false, "show the app version")

	fs.Usage = Usage

	err := fs.Parse(os.Args[1:])

	if err != nil {
		return config{}, err
	}

	conf := config{
		verboseMode: *verboseMode,
		versionMode: *versionMode,
	}

	usingDefaults := true

	// order for flag overrides: newline, word, character, byte.
	if *byteMode {
		usingDefaults = false

		conf.byteMode = true
		conf.lineMode = false
		conf.runeMode = false
		conf.wordMode = false
	}

	if *runeMode {
		usingDefaults = false

		conf.byteMode = false
		conf.lineMode = false
		conf.runeMode = true
		conf.wordMode = false
	}

	if *wordMode {
		usingDefaults = false

		conf.byteMode = false
		conf.lineMode = false
		conf.runeMode = false
		conf.wordMode = true
	}

	if *lineMode {
		usingDefaults = false

		conf.byteMode = false
		conf.lineMode = true
		conf.runeMode = false
		conf.wordMode = false
	}

	if usingDefaults {
		conf.byteMode = true
		conf.lineMode = true
		conf.runeMode = false
		conf.wordMode = true
	}

	conf.modes = map[string]bool{
		"byte": conf.byteMode,
		"line": conf.lineMode,
		"rune": conf.runeMode,
		"word": conf.wordMode,
	}

	return conf, nil
}

func getCounts(r io.Reader, conf config) results {
	counts := results{
		"byte": 0,
		"rune": 0,
		"line": 0,
		"word": 0,
	}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		text := scanner.Text() + "\n"

		if conf.byteMode {
			counts["byte"] += len(text)
		}

		if conf.runeMode {
			counts["rune"] += utf8.RuneCountInString(text)
		}

		if conf.wordMode {
			counts["word"] += len(strings.Fields(text))
		}

		if conf.lineMode {
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
