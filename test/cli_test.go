package appwc

import (
	"os"
	"testing"
)

// This is the main test function. This is the gatekeeper of all the tests in the appwc root package.
func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestApp(t *testing.T) {
}
