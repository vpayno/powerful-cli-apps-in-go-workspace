package appwc

import (
	"fmt"
	"os"
	"strings"
)

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
	lineMode    bool
	versionMode bool
}

var defaults = config{
	lineMode:    false,
	versionMode: false,
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
