package appwc

import (
	"bytes"
	"flag"
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

func TestBadFlag(t *testing.T) {
	testStderr, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe() err %v; want %v", err, nil)
	}

	osStderr := os.Stderr // keep backup of the real stdout
	os.Stderr = writer

	defer func() {
		// Undo what we changed when this test is done.
		os.Stderr = osStderr
	}()

	// It's a silly test but I need the practice.
	want := "flag provided but not defined: -x"

	// Run the function who's output we want to capture.
	os.Args = []string{"cli", "-x"}
	RunApp()

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStderr)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	got = strings.Split(got, "\n")[0]
	if got != want {
		t.Errorf("Usage(); want %q, got %q", want, got)
	}
}

func TestShowUsage(t *testing.T) {
	testStderr, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe() err %v; want %v", err, nil)
	}

	osStderr := os.Stderr // keep backup of the real stdout
	os.Stderr = writer

	defer func() {
		// Undo what we changed when this test is done.
		os.Stderr = osStderr
	}()

	// It's a silly test but I need the practice.
	want := "Usage: cli [OPTION]..."

	// Run the function who's output we want to capture.
	os.Args = []string{"cli", "-h"}
	RunApp()

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStderr)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	got = strings.Split(got, "\n")[0]
	if got != want {
		t.Errorf("Usage(); want %q, got %q", want, got)
	}
}

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

func TestSetupFlagsDefaults(t *testing.T) {
	want := config{
		byteMode: true,
		lineMode: true,
		wordMode: true,
		modes: map[string]bool{
			"byte": true,
			"line": true,
			"word": true,
		},
	}

	os.Args = []string{"test"}

	got, err := setup()

	if err != nil {
		t.Error(err)
	}

	if want.byteMode != got.byteMode {
		t.Errorf("setup() returned the wrong byte mode value. want: %v, got %v", want.byteMode, got.byteMode)
	}

	if want.lineMode != got.lineMode {
		t.Errorf("setup() returned the wrong line mode value. want: %v, got %v", want.lineMode, got.lineMode)
	}

	if want.wordMode != got.wordMode {
		t.Errorf("setup() returned the wrong word mode value. want: %v, got %v", want.wordMode, got.wordMode)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagsWordMode(t *testing.T) {
	want := config{
		wordMode: true,
		modes: map[string]bool{
			"word": true,
		},
	}

	os.Args = []string{"test", "-w"}

	got, err := setup()

	if err != nil {
		t.Error(err)
	}

	if want.wordMode != got.wordMode {
		t.Errorf("setup() returned the wrong word mode value. want: %v, got %v", want.wordMode, got.wordMode)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagsLineMode(t *testing.T) {
	want := config{
		lineMode: true,
		modes: map[string]bool{
			"line": true,
		},
	}

	// -l
	os.Args = []string{"test", "-l"}

	got, err := setup()

	if err != nil {
		t.Error(err)
	}

	if want.lineMode != got.lineMode {
		t.Errorf("setup() returned the wrong line mode value. want: %v, got %v", want.lineMode, got.lineMode)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagsByteMode(t *testing.T) {
	want := config{
		byteMode: true,
		modes: map[string]bool{
			"byte": true,
		},
	}

	// -l
	os.Args = []string{"test", "-b"}

	got, err := setup()

	if err != nil {
		t.Error(err)
	}

	if want.byteMode != got.byteMode {
		t.Errorf("setup() returned the wrong byte mode value. want: %v, got %v", want.byteMode, got.byteMode)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagVersion(t *testing.T) {
	// -V
	os.Args = []string{"test", "-V"}
	conf, err := setup()

	if err != nil {
		t.Error(err)
	}

	if !conf.versionMode {
		t.Errorf("versionMode: want %v, got %v", true, conf.versionMode)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestGetWordCount(t *testing.T) {
	b := bytes.NewBufferString("one two three four five\n")

	conf := config{
		wordMode: true,
		modes: map[string]bool{
			"word": true,
		},
	}

	want := 5
	got := getCounts(b, conf)["word"]

	if want != got {
		t.Errorf("Expected word count %d, got %d.\n", want, got)
	}
}

func TestGetLineCount(t *testing.T) {
	b := bytes.NewBufferString("one\ntwo\nthree\nfour\nfive\n")

	conf := config{
		lineMode: true,
		modes: map[string]bool{
			"line": true,
		},
	}

	want := 5
	got := getCounts(b, conf)["line"]

	if want != got {
		t.Errorf("Expected line count %d, got %d.\n", want, got)
	}
}

func TestGetByteCount(t *testing.T) {
	b := bytes.NewBufferString("0123456789\n0123456789\n")

	conf := config{
		byteMode: true,
		modes: map[string]bool{
			"byte": true,
		},
	}

	want := 22
	got := getCounts(b, conf)["byte"]

	if want != got {
		t.Errorf("Expected byte count %d, got %d.\n", want, got)
	}
}

func TestGetRuneCount(t *testing.T) {
	b := bytes.NewBufferString("0123456789\n0123456789\n")

	conf := config{
		charMode: true,
		modes: map[string]bool{
			"char": true,
		},
	}

	want := 22
	got := getCounts(b, conf)["char"]

	if want != got {
		t.Errorf("Expected char count %d, got %d.\n", want, got)
	}
}

func TestShowWordCountVerbose(t *testing.T) {
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

	counts := results{
		"word": 5,
	}

	conf := config{
		wordMode:    true,
		verboseMode: true,
		modes: map[string]bool{
			"word": true,
		},
	}

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%8d (word)\n", counts["word"])

	// Run the function who's output we want to capture.
	showCount(counts, conf)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("show word count: want %q, got %q", want, got)
	}
}

func TestShowLineCountVerbose(t *testing.T) {
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

	counts := results{
		"line": 5,
	}

	conf := config{
		lineMode:    true,
		verboseMode: true,
		modes: map[string]bool{
			"line": true,
		},
	}

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%8d (line)\n", counts["line"])

	// Run the function who's output we want to capture.
	showCount(counts, conf)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("show line count: want %q, got %q", want, got)
	}
}

func TestShowByteCountVerbose(t *testing.T) {
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

	counts := results{
		"byte": 5,
	}

	conf := config{
		byteMode:    true,
		verboseMode: true,
		modes: map[string]bool{
			"byte": true,
		},
	}

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%8d (byte)\n", counts["byte"])

	// Run the function who's output we want to capture.
	showCount(counts, conf)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("show byte count: want %q, got %q", want, got)
	}
}

func TestShowRuneCountVerbose(t *testing.T) {
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

	counts := results{
		"char": 5,
	}

	conf := config{
		charMode:    true,
		verboseMode: true,
		modes: map[string]bool{
			"char": true,
		},
	}

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%8d (char)\n", counts["char"])

	// Run the function who's output we want to capture.
	showCount(counts, conf)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("show byte count: want %q, got %q", want, got)
	}
}

func TestShowWordCount(t *testing.T) {
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

	counts := results{
		"word": 5,
	}

	conf := config{
		wordMode:    true,
		verboseMode: false,
		modes: map[string]bool{
			"word": true,
		},
	}

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%d\n", counts["word"])

	// Run the function who's output we want to capture.
	showCount(counts, conf)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("show word count: want %q, got %q", want, got)
	}
}

func TestShowLineCount(t *testing.T) {
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

	counts := results{
		"line": 5,
	}

	conf := config{
		lineMode:    true,
		verboseMode: false,
		modes: map[string]bool{
			"line": true,
		},
	}

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%d\n", counts["line"])

	// Run the function who's output we want to capture.
	showCount(counts, conf)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("show line count: want %q, got %q", want, got)
	}
}

func TestShowByteCount(t *testing.T) {
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

	counts := results{
		"byte": 5,
	}

	conf := config{
		byteMode:    true,
		verboseMode: false,
		modes: map[string]bool{
			"byte": true,
		},
	}

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%d\n", counts["byte"])

	// Run the function who's output we want to capture.
	showCount(counts, conf)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("show byte count: want %q, got %q", want, got)
	}
}

func TestShowRuneCount(t *testing.T) {
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

	counts := results{
		"char": 5,
	}

	conf := config{
		charMode:    true,
		verboseMode: false,
		modes: map[string]bool{
			"char": true,
		},
	}

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%d\n", counts["char"])

	// Run the function who's output we want to capture.
	showCount(counts, conf)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("show char count: want %q, got %q", want, got)
	}
}

func TestRunApp(t *testing.T) {
	os.Args = []string{"test", "-v"}

	RunApp()

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestRunAppFlagVersion(t *testing.T) {
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

	want := "\nWord Count Version: " + metadata.version + "\n\n\n"

	os.Args = []string{"test", "-V"}
	RunApp()

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("RunApp (Flag -V): want %q, got %q", want, got)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestRunAppFlagByteAndRune(t *testing.T) {
	want := config{
		byteMode: false,
		charMode: true,
	}

	os.Args = []string{"test", "-b", "-r"}
	got, _ := setup()

	if got.byteMode != want.byteMode {
		t.Errorf("setup flags -r & -b (byteMode): want %v, got %v", want.byteMode, got.byteMode)
	}

	if got.charMode != want.charMode {
		t.Errorf("setup flags -r & -b (charMode): want %v, got %v", want.charMode, got.charMode)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestRunAppFlagWordAndRune(t *testing.T) {
	want := config{
		wordMode: true,
		charMode: false,
	}

	os.Args = []string{"test", "-b", "-w"}
	got, _ := setup()

	if got.wordMode != want.wordMode {
		t.Errorf("setup flags -r & -b (wordMode): want %v, got %v", want.wordMode, got.wordMode)
	}

	if got.charMode != want.charMode {
		t.Errorf("setup flags -r & -b (charMode): want %v, got %v", want.charMode, got.charMode)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}
