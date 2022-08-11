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

func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
	flag.PrintDefaults()
}

func setup() (config, error) {
	byteMode := flag.Bool("b", defaults.byteMode, "count bytes instead of words")
	lineMode := flag.Bool("l", defaults.lineMode, "count lines instead of words")
	runeMode := flag.Bool("r", defaults.byteMode, "count runes instead of words")
	wordMode := flag.Bool("w", defaults.wordMode, "count words (default)")
	verboseMode := flag.Bool("v", false, "verbose mode")
	versionMode := flag.Bool("V", false, "show the app version")

	flag.Usage = Usage

	flag.Parse()

	conf := config{
		byteMode:    *byteMode,
		lineMode:    *lineMode,
		runeMode:    *runeMode,
		wordMode:    *wordMode,
		verboseMode: *verboseMode,
		versionMode: *versionMode,
	}

	// If line mode is set, disable the other modes.
	if conf.lineMode {
		conf.byteMode = false
		conf.runeMode = false
	}

	// fail if both byte and rune modes are set.
	if conf.byteMode && conf.runeMode {
		err := errors.New("-b (byte count mode) and -r (rune count mode) can't be used at the same time")
		return config{}, err
	}

	// byte, rune and line mode can override word mode since it's enabled by default
	if conf.byteMode || conf.runeMode || conf.lineMode {
		conf.wordMode = false
	}

	return conf, nil
}

func getCount(r io.Reader, conf config) int {
	if conf.byteMode || conf.runeMode {
		return getCountBytes(r, conf)
	}

	scanner := bufio.NewScanner(r)

	var count int

	if conf.wordMode {
		scanner.Split(bufio.ScanWords)
	}

	for scanner.Scan() {
		count++
	}

	return count
}

func getCountBytes(input io.Reader, conf config) int {
	r := bufio.NewReader(input)

	var err error
	var count int
	var str string

	// assuing err is io.EOF
	for err == nil {
		// instead of using /n, using null
		str, err = r.ReadString('\x00')

		switch {
		case conf.byteMode:
			count += len(str)
		case conf.runeMode:
			count += utf8.RuneCountInString(str)
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
		case conf.runeMode:
			prompt = "rune count: "
		case conf.wordMode:
			prompt = "word count: "
		}
	}

	fmt.Print(prompt)
	fmt.Printf("%d\n", n)
}
