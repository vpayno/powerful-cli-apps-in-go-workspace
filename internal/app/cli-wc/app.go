package appwc

import (
	"flag"
	"fmt"
)

func showBanner() {
	fmt.Println(metadata.name + " Version " + metadata.version)
	fmt.Println()
}

// RunApp is called my the main function. It's basically the main function of the app.
func RunApp() {
	c := setup()

	fmt.Printf("config: %#v\n", c)

	showBanner()
}

func setup() config {
	lineMode := flag.Bool("l", defaults.lineMode, "count lines instead of words")
	version := flag.Bool("V", false, "show the app version")
	flag.Parse()

	if *version {
		showVersion()
		Exit(0, "")
	}

	return config{
		lineMode: *lineMode,
	}
}
