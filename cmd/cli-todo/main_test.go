package main

import (
	"os"
	"testing"

	apptodo "github.com/vpayno/powerful-cli-apps-in-go-workspace/internal/app/cli-todo"
)

// This is the main test function. This is the gatekeeper of all the tests in the main package.
func TestMain(m *testing.M) {
	exitCode := m.Run()

	os.Exit(exitCode)
}

// The functions in main() are already tested. Just running them together with zero test questions.
func TestMain_app(_ *testing.T) {
	os.Args = []string{"test", "-V"}

	apptodo.SetVersion(version)
	main()
}
