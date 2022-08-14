package appwc

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// This is used to disable side-effects like flag calling os.Exit() during tests when it encounters an error or the --help flag.
var flagExitErrorBehavior = flag.ExitOnError
var flagSet = flag.NewFlagSet(os.Args[0], flagExitErrorBehavior)

type appInfo struct {
	name       string
	version    string
	gitVersion string
	gitHash    string
	buildTime  string
}

var metadata = appInfo{
	name:    "Word Count",
	version: "0.0.0",
}

type config struct {
	modes       map[string]bool
	verboseMode bool
	versionMode bool
}

type results map[string]int

const (
	usage = `Usage: %s [OPTION]...

Print newline, word, and byte counts for stdin input.

A word is a non-zero-length sequence of characters delimited by white space.

The options below may be used to select which counts are printed, always in
the following order: newline, word, character, byte.

Options:
  -c, --bytes            print the byte counts
  -m, --chars            print the character counts
  -l, --lines            print the newline counts
  -w, --words            print the word counts
  -h, --help             display this help and exit
  -v, --version          output version information and exit
  -V, --verbose          verbose mode
`
)

var flagDefaults = map[string]bool{
	"byte": true,
	"char": false,
	"line": true,
	"word": true,
}

// Printing order: line, word, byte, char
var printOrder = []string{
	"line",
	"word",
	"byte",
	"char",
}

// SetVersion is used my the main package to pass version information to the app package.
func SetVersion(b []byte) {
	slice := strings.Split(string(b), "\n")
	slice = slice[:len(slice)-1]

	if slice[0] != "" {
		metadata.version = slice[0]
	}

	if len(slice) > 1 {
		if slice[1] != "" {
			metadata.gitVersion = slice[1]
		}
		if slice[2] != "" {
			metadata.gitHash = slice[2]
		}
		if slice[3] != "" {
			metadata.buildTime = slice[3]
		}
	}
}

func showVersion() {
	fmt.Println()
	fmt.Printf("%s Version: %s\n\n", metadata.name, metadata.version)

	if metadata.gitVersion != "" {
		fmt.Printf("git version: %s\n", metadata.gitVersion)
	}

	if metadata.gitHash != "" {
		fmt.Printf("   git hash: %s\n", metadata.gitHash)
	}

	if metadata.buildTime != "" {
		fmt.Printf(" build time: %s\n", metadata.buildTime)
	}

	fmt.Println()
}

// OSExit is used to Money Patch the Exit function during testing.
var OSExit = os.Exit

// Exit is used to prematurely end the application with an exit code and message to stdout.
func Exit(code int, msg string) {
	fmt.Println(msg)
	OSExit(code)
}
