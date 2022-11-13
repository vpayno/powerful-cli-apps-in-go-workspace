// Package appwc is the module with the cli logic for the wc main application.
package appwc

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// OSExit is used to Money Patch the Exit function during testing.
var OSExit = os.Exit

// This is used to disable side-effects like flag calling os.Exit() during tests when it encounters an error or the --help flag.
var flagExitErrorBehavior = flag.ExitOnError
var flagSet = flag.NewFlagSet(os.Args[0], flagExitErrorBehavior)

type appInfo struct {
	name      string
	version   string
	gitHash   string
	buildTime string
	cliName   string
	appData   string
}

var metadata = appInfo{
	cliName: "cli-wc",
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
	updateMode  bool
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
  -u, --update           update mode
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

// SetAppInfo is used by the main package to pass application information th the app package.
func SetAppInfo(appData string) {
	metadata.appData = appData
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

func getInstallURL() string {
	// reStrMatch := `^module .*$`
	reStrMatch := `^.*$`
	var srcURL string

	for _, line := range strings.Split(metadata.appData, "\n") {
		match, _ := regexp.MatchString(reStrMatch, line)

		if match {
			// reStrURL := `^(module )([a-z].*)(|/v[0-9]+)?$`
			reStrURL := `^([a-z].*)(|/v[0-9]+)$`
			r, _ := regexp.Compile(reStrURL)

			// srcURL = r.ReplaceAllString(line, `$2`)
			srcURL = r.ReplaceAllString(line, `$1`)

			break
		}
	}

	appName := metadata.cliName

	return srcURL + "/cmd/" + appName + "@latest"
}

func updateApp() error {
	installURL := getInstallURL()

	fmt.Println("Running:", "go", "install", installURL)
	fmt.Println()

	out, err := exec.Command("go", "install", installURL).Output()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("No errors detected during installation.")
	fmt.Println()
	if len(string(out)) > 0 {
		fmt.Println(string(out))
		fmt.Println()
	}

	// TODO: add checks for Go environment variables

	appName := metadata.cliName

	fmt.Printf("Running '%s --version': ", appName)

	out, err = exec.Command(appName, "--version").Output()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(string(out)) > 0 {
		fmt.Println(string(out))
		fmt.Println()
	}

	return nil
}

// Exit is used to end the application with an exit code and message to stderr.
func Exit(err error) {
	var code int

	if err != nil {
		code += 1
		fmt.Println(err)
	}

	OSExit(code)
}
