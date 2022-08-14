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
}

func setup() (config, error) {

	var byteFlag bool
	var charFlag bool
	var lineFlag bool
	var wordFlag bool

	flagSet.BoolVar(&byteFlag, "c", false, "print the byte counts")
	flagSet.BoolVar(&byteFlag, "byte", false, "print the byte counts")
	flagSet.BoolVar(&charFlag, "m", false, "print the character counts")
	flagSet.BoolVar(&charFlag, "char", false, "print the character counts")
	flagSet.BoolVar(&lineFlag, "l", false, "print the newline counts")
	flagSet.BoolVar(&lineFlag, "line", false, "print the newline counts")
	flagSet.BoolVar(&wordFlag, "w", false, "print the word counts")
	flagSet.BoolVar(&wordFlag, "word", false, "print the word counts")

	var verboseFlag bool
	var versionFlag bool

	flagSet.BoolVar(&verboseFlag, "v", false, "verbose mode")
	flagSet.BoolVar(&verboseFlag, "verbose", false, "verbose mode")
	flagSet.BoolVar(&versionFlag, "V", false, "output version information and exit")
	flagSet.BoolVar(&versionFlag, "version", false, "output version information and exit")

	flagSet.Usage = Usage

	err := flagSet.Parse(os.Args[1:])

	if err != nil {
		return config{}, err
	}

	conf := config{
		verboseMode: verboseFlag,
		versionMode: versionFlag,
	}

	usingDefaults := true

	var byteMode bool
	var charMode bool
	var lineMode bool
	var wordMode bool

	// order for flag overrides: newline, word, character, byte.
	switch {
	case lineFlag:
		usingDefaults = false

		byteMode = false
		charMode = false
		lineMode = true
		wordMode = false

	case wordFlag:
		usingDefaults = false

		byteMode = false
		charMode = false
		lineMode = false
		wordMode = true

	case charFlag:
		usingDefaults = false

		byteMode = false
		charMode = true
		lineMode = false
		wordMode = false

	case byteFlag:
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
