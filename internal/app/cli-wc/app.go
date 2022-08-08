package appwc

import (
	"bufio"
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
	lineMode := flag.Bool("l", defaults.lineMode, "count lines instead of words")
	versionMode := flag.Bool("V", false, "show the app version")
	flag.Parse()

	return config{
		lineMode:    *lineMode,
		versionMode: *versionMode,
	}
}

func getCount(r io.Reader, c config) int {
	scanner := bufio.NewScanner(r)

	if !c.lineMode {
		scanner.Split(bufio.ScanWords)
	}

	var count int

	for scanner.Scan() {
		count++
	}

	return count
}

func showCount(n int, conf config) {
	if conf.lineMode {
		fmt.Printf("line count: %d\n", n)
	} else {
		fmt.Printf("word count: %d\n", n)
	}
}
