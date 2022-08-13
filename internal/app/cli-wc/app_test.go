package appwc

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
  Stdout testing code borrowed from Jon Calhoun's FizzBuzz example.
  https://courses.calhoun.io/lessons/les_algo_m01_08
  https://github.com/joncalhoun/algorithmswithgo.com/blob/master/module01/fizz_buzz_test.go
*/

// Use this to put modules, functions in testing mode.
func setupTestEnv() {
}

// Use this to undo things you did in setupTestEnv()
func teardownTestEnv() {
}

func TestBadFlag(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
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

	if want.modes["byte"] != got.modes["byte"] {
		t.Errorf("setup() returned the wrong byte mode value. want: %v, got %v", want.modes["byte"], got.modes["byte"])
	}

	if want.modes["line"] != got.modes["line"] {
		t.Errorf("setup() returned the wrong line mode value. want: %v, got %v", want.modes["line"], got.modes["line"])
	}

	if want.modes["word"] != got.modes["word"] {
		t.Errorf("setup() returned the wrong word mode value. want: %v, got %v", want.modes["word"], got.modes["word"])
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagsWordMode(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"word": true,
		},
	}

	os.Args = []string{"test", "-w"}

	got, err := setup()

	if err != nil {
		t.Error(err)
	}

	if want.modes["word"] != got.modes["word"] {
		t.Errorf("setup() returned the wrong word mode value. want: %v, got %v", want.modes["word"], got.modes["word"])
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagsLineMode(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
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

	if want.modes["line"] != got.modes["line"] {
		t.Errorf("setup() returned the wrong line mode value. want: %v, got %v", want.modes["line"], got.modes["line"])
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagsByteMode(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
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

	if want.modes["byte"] != got.modes["byte"] {
		t.Errorf("setup() returned the wrong byte mode value. want: %v, got %v", want.modes["byte"], got.modes["byte"])
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagVersion(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

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

func TestGetCounts(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			b := bytes.NewBufferString(tc.input)
			c := config{
				modes: map[string]bool{
					"byte": true,
					"char": true,
					"word": true,
					"line": true,
				},
			}

			got := getCounts(b, c)

			assert.Equal(t, tc.wantByte, got["byte"], "byte counts aren't equal")
			assert.Equal(t, tc.wantChar, got["char"], "char counts aren't equal")
			assert.Equal(t, tc.wantWord, got["word"], "word counts aren't equal")
			assert.Equal(t, tc.wantLine, got["line"], "line counts aren't equal")
		})
	}
}

func TestGetWordCount(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	b := bytes.NewBufferString("one two three four five\n")

	conf := config{
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
	setupTestEnv()
	defer teardownTestEnv()

	b := bytes.NewBufferString("one\ntwo\nthree\nfour\nfive\n")

	conf := config{
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
	setupTestEnv()
	defer teardownTestEnv()

	b := bytes.NewBufferString("0123456789\n0123456789\n")

	conf := config{
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
	setupTestEnv()
	defer teardownTestEnv()

	b := bytes.NewBufferString("0123456789\n0123456789\n")

	conf := config{
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
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

	os.Args = []string{"test", "-v"}

	RunApp()

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestRunAppFlagVersion(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

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
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"char": true,
			"byte": false,
		},
	}

	os.Args = []string{"test", "-b", "-r"}
	got, _ := setup()

	if got.modes["byte"] != want.modes["byte"] {
		t.Errorf("setup flags -r & -b (byteMode): want %v, got %v", want.modes["byte"], got.modes["byte"])
	}

	if got.modes["char"] != want.modes["char"] {
		t.Errorf("setup flags -r & -b (charMode): want %v, got %v", want.modes["char"], got.modes["char"])
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestRunAppFlagWordAndRune(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"char": false,
			"word": true,
		},
	}

	os.Args = []string{"test", "-b", "-w"}
	got, _ := setup()

	if got.modes["word"] != want.modes["word"] {
		t.Errorf("setup flags -r & -b (wordMode): want %v, got %v", want.modes["word"], got.modes["word"])
	}

	if got.modes["char"] != want.modes["char"] {
		t.Errorf("setup flags -r & -b (charMode): want %v, got %v", want.modes["char"], got.modes["char"])
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}
