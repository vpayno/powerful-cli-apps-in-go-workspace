package main

import (
	"os"
	"testing"

	appwc "github.com/vpayno/powerful-cli-apps-in-go-workspace/internal/app/cli-wc"
)

// This is the main test function. This is the gatekeeper of all the tests in the main package.
func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

// The functions in main() are already tested. Just running them together with zero test questions.
func TestMain_app(t *testing.T) {
	os.Args = []string{"test", "-V"}

	appwc.SetVersion(version)
	main()
}
