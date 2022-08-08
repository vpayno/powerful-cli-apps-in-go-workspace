package appwc

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

/*
  Stdout testing code borrowed from Jon Calhoun's FizzBuzz example.
  https://courses.calhoun.io/lessons/les_algo_m01_08
  https://github.com/joncalhoun/algorithmswithgo.com/blob/master/module01/fizz_buzz_test.go
*/

func TestShowBanner(t *testing.T) {
	testStdout, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe() err %v; want %v", err, nil)
	}

	osStdout := os.Stdout // keep backup of the real stdout
	os.Stdout = writer

	defer func() {
		// Undo what we changed when this test is done.
		os.Stdout = osStdout
	}()

	// It's a silly test but I need the practice.
	want := metadata.name + " Version " + metadata.version + "\n\n"

	// Run the function who's output we want to capture.
	showBanner()

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("showBanner(); want %q, got %q", want, got)
	}
}

func TestSetupFlags(t *testing.T) {
	want := config{
		lineMode: true,
	}

	// -l
	os.Args = []string{"test", "-l"}

	got := setup()

	if want.lineMode != got.lineMode {
		t.Errorf("setup() returned the wrong line mode value. want: %v, got %v", want.lineMode, got.lineMode)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagVersion(t *testing.T) {
	testStdout, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe() err %v; want %v", err, nil)
	}

	osStdout := os.Stdout // keep backup of the real stdout
	os.Stdout = writer

	defer func() {
		// Undo what we changed when this test is done.
		os.Stdout = osStdout
	}()

	metadata.gitVersion = "1.2.3-456-abcdef"
	metadata.gitHash = "abcdefabcdefabcdefabcdefabcdefabcdef"
	metadata.buildTime = "date-time"

	want := "\n"
	want += fmt.Sprintf("%s Version: %s\n\n", metadata.name, metadata.version)
	want += fmt.Sprintf("git version: %s\n", metadata.gitVersion)
	want += fmt.Sprintf("   git hash: %s\n", metadata.gitHash)
	want += fmt.Sprintf(" build time: %s\n", metadata.buildTime)
	want += "\n"
	want += "\n"

	// -V
	os.Args = []string{"test", "-V"}
	OSExitBackup := OSExit
	OSExit = func(code int) { _ = code }
	// It's not going to exit, it will return a value we don't want.
	_ = setup()
	OSExit = OSExitBackup

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("Exit(); want %q, got %q", want, got)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestRunApp(t *testing.T) {
	os.Args = []string{"test", "-l"}

	RunApp()
}
