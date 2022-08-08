package main

import (
	"os"
	"testing"

	appwc "github.com/vpayno/powerful-cli-apps-in-go-workspace/internal/app/cli-wc"
)

// The functions in main() are already tested. Just running them together with zero test questions.
func TestMain(t *testing.T) {
	os.Args = []string{"test", "-V"}

	appwc.SetVersion(version)
	main()
}
