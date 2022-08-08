package appwc

import (
	"bufio"
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
	conf := setup()

	if conf.versionMode {
		showVersion()
		return
	}

	showBanner()

	var count int

	count = getCount(os.Stdin, conf)

	showCount(count, conf)
}

func setup() config {
	byteMode := flag.Bool("b", defaults.lineMode, "count bytes instead of words")
	lineMode := flag.Bool("l", defaults.lineMode, "count lines instead of words")
	versionMode := flag.Bool("V", false, "show the app version")
	flag.Parse()

	return config{
		byteMode:    *byteMode,
		lineMode:    *lineMode,
		versionMode: *versionMode,
	}
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
	switch {
	case conf.byteMode:
		fmt.Printf("byte count: %d\n", n)
	case conf.lineMode:
		fmt.Printf("line count: %d\n", n)
	default:
		fmt.Printf("word count: %d\n", n)
	}
}
