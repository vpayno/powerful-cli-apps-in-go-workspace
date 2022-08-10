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
	wordMode := flag.Bool("w", defaults.wordMode, "count words (default)")
	verboseMode := flag.Bool("v", false, "verbose mode")
	versionMode := flag.Bool("V", false, "show the app version")
	flag.Parse()

	conf := config{
		byteMode:    *byteMode,
		lineMode:    *lineMode,
		wordMode:    *wordMode,
		verboseMode: *verboseMode,
		versionMode: *versionMode,
	}

	if conf.byteMode && conf.lineMode {
		err := errors.New("-b (byte count mode) and -l (line count mode) can't be used at the same time")
		return config{}, err
		conf.wordMode = false // byte and line count can override word count since it's enabled by default
	}

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
	if conf.byteMode {
		return getCountBytes(r)
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

func getCountBytes(input io.Reader) int {
	r := bufio.NewReader(input)

	var err error
	var count int
	var str string

	// assuing err is io.EOF
	for err == nil {
		// instead of using /n, using null
		str, err = r.ReadString('\x00')

		count += len(str)
	}

	return count
}

func getCountRunes(s string) int {
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
