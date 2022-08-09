package appwc

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

/*
  Stdout testing code borrowed from Jon Calhoun's FizzBuzz example.
  https://courses.calhoun.io/lessons/les_algo_m01_08
  https://github.com/joncalhoun/algorithmswithgo.com/blob/master/module01/fizz_buzz_test.go
*/

func TestExitVerbose(t *testing.T) {
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

	// It's a silly test but I need the practice mocking.
	code := 123
	msg := "testing Exit()"
	want := fmt.Sprintf("%s\nCalling os.Exit(%d)...\n", msg, code)

	OSExitBackup := OSExit
	OSExit = func(code int) { fmt.Printf("Calling os.Exit(%d)...\n", code) }
	Exit(code, msg)
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
}

func TestShowVersion(t *testing.T) {
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

	wantMetadata := appInfo{
		name:       metadata.name,
		version:    "version",
		gitVersion: "gitVersion",
		gitHash:    "gitHash",
		buildTime:  "buildTime",
	}

	strSlice := []string{wantMetadata.version, wantMetadata.gitVersion, wantMetadata.gitHash, wantMetadata.buildTime}
	b := []byte(strings.Join(strSlice, "\n") + "\n")
	SetVersion(b)

	// It's a silly test but I need the practice.
	want := "\n"
	want += fmt.Sprintf("%s Version: %s\n\n", wantMetadata.name, wantMetadata.version)
	want += fmt.Sprintf("git version: %s\n", wantMetadata.gitVersion)
	want += fmt.Sprintf("   git hash: %s\n", wantMetadata.gitHash)
	want += fmt.Sprintf(" build time: %s\n", wantMetadata.buildTime)
	want += "\n"

	// Run the function who's output we want to capture.
	showVersion()

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("showVersion(); want %q, got %q", want, got)
	}
}

func TestSetVersion(t *testing.T) {
	want := appInfo{
		name:       "name",
		version:    "version",
		gitVersion: "gitVersion",
		gitHash:    "gitHash",
		buildTime:  "buildTime",
	}

	strSlice := []string{want.version, want.gitVersion, want.gitHash, want.buildTime}
	b := []byte(strings.Join(strSlice, "\n") + "\n")
	SetVersion(b)

	got := metadata

	if want.version != got.version {
		t.Errorf("expected version to be set to %q, got %q", want.version, got.version)
	}

	if want.gitVersion != got.gitVersion {
		t.Errorf("expected gitVersion to be set to %q, got %q", want.gitVersion, got.gitVersion)
	}

	if want.gitHash != got.gitHash {
		t.Errorf("expected gitHash to be set to %q, got %q", want.gitHash, got.gitHash)
	}

	if want.buildTime != got.buildTime {
		t.Errorf("expected buildTime to be set to %q, got %q", want.buildTime, got.buildTime)
	}
}
