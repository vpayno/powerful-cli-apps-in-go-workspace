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
	name      string
	version   string
	gitHash   string
	buildTime string
}

var metadata = appInfo{
	name:    "Word Count",
	version: "0.0.0",
}

// Using uint64 because uint switches between 32-bit and 64-bit depending on the architecture instead of being consistent.
const (
	byteFlag uint64 = 1
	charFlag uint64 = 1 << iota
	wordFlag
	lineFlag
	lengthFlag
)

// default modes: byte, word, line
var defaultFlags = byteFlag | wordFlag | lineFlag

type config struct {
	flags       uint64
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
  -L, --max-line-length  print the maximum display width
  -l, --lines            print the newline counts
  -w, --words            print the word counts
  -h, --help             display this help and exit
  -v, --version          output version information and exit
  -V, --verbose          verbose mode
`
)

var defaultModes = map[string]bool{
	"byte":   defaultFlags&byteFlag != 0,
	"char":   defaultFlags&charFlag != 0,
	"length": defaultFlags&lengthFlag != 0,
	"line":   defaultFlags&lineFlag != 0,
	"word":   defaultFlags&wordFlag != 0,
}

// Printing order: newline, word, character, byte, max-line-length.
var printOrder = []string{
	"line",
	"word",
	"char",
	"byte",
	"length",
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
			metadata.gitHash = slice[1]
		}
		if slice[2] != "" {
			metadata.buildTime = slice[2]
		}
	}
}

func showVersion(conf config) {
	if !conf.verboseMode {
		fmt.Printf("%s\n", metadata.version)
		return
	}

	fmt.Println()
	fmt.Printf("%s Version: %s\n\n", metadata.name, metadata.version)

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
