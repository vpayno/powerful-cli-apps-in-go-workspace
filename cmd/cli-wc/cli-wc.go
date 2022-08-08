// Package main is the main module for the wc application.
package main

import (
	appwc "github.com/vpayno/powerful-cli-apps-in-go-workspace/internal/app/cli-wc"

	_ "embed"
)

//go:generate bash ../../scripts/go-generate-helper-git-version-info
//go:embed .version.txt
var version []byte

func init() {
	appwc.SetVersion(version)
}

func main() {
	appwc.RunApp()
}
