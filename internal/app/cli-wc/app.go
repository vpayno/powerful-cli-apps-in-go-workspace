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
func RunApp() error {
	conf, err := setup()

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		return nil
	}

	if conf.versionMode {
		showVersion(conf)
		return nil
	}

	if conf.verboseMode {
		showBanner()
	}

	if conf.updateMode {
		return updateApp()
	}

	counts := getCounts(os.Stdin, conf)

	showCount(counts, conf)

	return nil
}

// Usage prints the command-line usage help message.
func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
}

func setup() (config, error) {

	var byteFlagPtr bool
	var charFlagPtr bool
	var lengthFlagPtr bool
	var lineFlagPtr bool
	var wordFlagPtr bool

	flagSet.BoolVar(&byteFlagPtr, "c", false, "print the byte counts")
	flagSet.BoolVar(&byteFlagPtr, "bytes", false, "print the byte counts")
	flagSet.BoolVar(&charFlagPtr, "m", false, "print the character counts")
	flagSet.BoolVar(&charFlagPtr, "chars", false, "print the character counts")
	flagSet.BoolVar(&lengthFlagPtr, "L", false, "print the maximum display width")
	flagSet.BoolVar(&lengthFlagPtr, "max-line-length", false, "print the maximum display width")
	flagSet.BoolVar(&lineFlagPtr, "l", false, "print the newline counts")
	flagSet.BoolVar(&lineFlagPtr, "lines", false, "print the newline counts")
	flagSet.BoolVar(&wordFlagPtr, "w", false, "print the word counts")
	flagSet.BoolVar(&wordFlagPtr, "words", false, "print the word counts")

	var updateFlagPtr bool
	var verboseFlagPtr bool
	var versionFlagPtr bool

	flagSet.BoolVar(&updateFlagPtr, "u", false, "update mode")
	flagSet.BoolVar(&updateFlagPtr, "update", false, "update mode")
	flagSet.BoolVar(&verboseFlagPtr, "v", false, "verbose mode")
	flagSet.BoolVar(&verboseFlagPtr, "verbose", false, "verbose mode")
	flagSet.BoolVar(&versionFlagPtr, "V", false, "output version information and exit")
	flagSet.BoolVar(&versionFlagPtr, "version", false, "output version information and exit")

	flagSet.Usage = Usage

	err := flagSet.Parse(os.Args[1:])

	if err != nil {
		return config{}, err
	}

	conf := config{
		flags:       defaultFlags,
		updateMode:  updateFlagPtr,
		verboseMode: verboseFlagPtr,
		versionMode: versionFlagPtr,
		modes:       defaultModes,
	}

	if byteFlagPtr || charFlagPtr || lengthFlagPtr || lineFlagPtr || wordFlagPtr {
		conf.flags = byteFlag | charFlag | wordFlag | lineFlag | lengthFlag
	}

	conf.modes = map[string]bool{
		"byte":   conf.flags&byteFlag != 0,
		"char":   conf.flags&charFlag != 0,
		"length": conf.flags&lengthFlag != 0,
		"line":   conf.flags&lineFlag != 0,
		"word":   conf.flags&wordFlag != 0,
	}

	return conf, nil
}

func getCounts(r io.Reader, conf config) results {
	counts := results{
		"byte":   0,
		"char":   0,
		"length": 0,
		"line":   0,
		"word":   0,
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

		if conf.modes["length"] {
			maxLength := utf8.RuneCountInString(strings.TrimSuffix(text, "\n"))

			if maxLength > counts["length"] {
				counts["length"] = maxLength
			}
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
