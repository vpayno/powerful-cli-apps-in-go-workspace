// Package appwc is the module with the cli logic for the wc main application.
package appwc

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
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
	byteMode := flag.Bool("b", defaults.byteMode, "count bytes instead of words")
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

	// setting this after testing if both -b and -l are set simplifies this step.
	conf.wordMode = !(conf.byteMode || conf.lineMode)

	return conf, nil
}

type readCounter struct {
	io.Reader
	bytesRead int
}

// Read returns the number of bytes counted or an error code.
func (r *readCounter) Read(b []byte) (int, error) {
	count, err := r.Reader.Read(b)
	r.bytesRead += count

	return count, err
}

func getCount(r io.Reader, conf config) int {
	scanner := bufio.NewScanner(r)

	var count int

	if conf.wordMode {
		scanner.Split(bufio.ScanWords)
	}

	for scanner.Scan() {
		if conf.byteMode {
			count += getCountBytes(scanner.Text())
		} else {
			count++
		}
	}

	return count
}

func getCountBytes(s string) int {
	b := &readCounter{Reader: bytes.NewBufferString(s)}

	scanner := bufio.NewScanner(b)

	var count int

	for scanner.Scan() {
		count += b.bytesRead
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
