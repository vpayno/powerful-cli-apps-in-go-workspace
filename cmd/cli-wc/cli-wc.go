// Package main is the main module for the wc application.
package main

import (
	appwc "github.com/vpayno/powerful-cli-apps-in-go-workspace/internal/app/cli-wc"

	_ "embed"
)

//go:generate bash ../../scripts/go-generate-helper-git-version-info
//go:embed .version.txt
var version []byte

//go:generate bash ../../scripts/go-generate-helper-git-app-info
//go:embed .app_info.txt
var goModData string

func init() {
	appwc.SetVersion(version)
	appwc.SetAppInfo(goModData)
}

func main() {
	var err error
	defer func() {
		appwc.Exit(err)
	}()

	err = appwc.RunApp()
}
