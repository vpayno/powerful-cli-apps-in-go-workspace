// Package main is the main module for the wc application.
package main

import (
	apptodo "github.com/vpayno/powerful-cli-apps-in-go-workspace/internal/app/cli-todo"

	_ "embed"
)

//go:generate bash ../../scripts/go-generate-helper-git-version-info
//go:embed .version.txt
var version []byte

func init() {
	apptodo.SetVersion(version)
}

func main() {
	apptodo.RunApp()
}
