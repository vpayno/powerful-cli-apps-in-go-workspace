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
	flagExitErrorBehavior = flag.ContinueOnError
	flagSet = flag.NewFlagSet(os.Args[0], flagExitErrorBehavior)
}

// Use this to undo things you did in setupTestEnv()
func teardownTestEnv() {
	flagExitErrorBehavior = flag.ExitOnError
	flagSet = flag.NewFlagSet(os.Args[0], flagExitErrorBehavior)
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

	want1 := "Usage: cli [OPTION]..."
	want2 := "output version information and exit"

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

	got1 := strings.Split(got, "\n")[0]
	if got1 != want1 {
		t.Errorf("Usage(); want %q, got %q", want1, got1)
	}

	if !strings.Contains(got, want2) {
		t.Errorf("Usage(); %q doesn't contain %q", got, want2)
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
}

func TestSetupFlagsWordMode(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"word": true,
		},
	}

	os.Args = []string{"test", "--words"}

	got, err := setup()

	if err != nil {
		t.Error(err)
	}

	if want.modes["word"] != got.modes["word"] {
		t.Errorf("setup() returned the wrong word mode value. want: %v, got %v", want.modes["word"], got.modes["word"])
	}
}

func TestSetupFlagsLineMode(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"line": true,
		},
	}

	os.Args = []string{"test", "--lines"}

	got, err := setup()

	if err != nil {
		t.Error(err)
	}

	if want.modes["line"] != got.modes["line"] {
		t.Errorf("setup() returned the wrong line mode value. want: %v, got %v", want.modes["line"], got.modes["line"])
	}
}

func TestSetupFlagsByteMode(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"byte": true,
		},
	}

	os.Args = []string{"test", "--bytes"}

	got, err := setup()

	if err != nil {
		t.Error(err)
	}

	if want.modes["byte"] != got.modes["byte"] {
		t.Errorf("setup() returned the wrong byte mode value. want: %v, got %v", want.modes["byte"], got.modes["byte"])
	}
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

func TestGetCharCount(t *testing.T) {
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
	want := fmt.Sprintf("%d (word)\n", counts["word"])

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
	want := fmt.Sprintf("%d (line)\n", counts["line"])

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
	want := fmt.Sprintf("%d (byte)\n", counts["byte"])

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

func TestShowCharCountVerbose(t *testing.T) {
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
	want := fmt.Sprintf("%d (char)\n", counts["char"])

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

func TestShowCharCount(t *testing.T) {
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
}

func TestRunAppFlagByteAndChar(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"char": true,
			"byte": true,
		},
	}

	os.Args = []string{"test", "--bytes", "--chars"}
	got, _ := setup()

	if got.modes["byte"] != want.modes["byte"] {
		t.Errorf("setup flags --chars & --bytes: want %v, got %v", want.modes["byte"], got.modes["byte"])
	}

	if got.modes["char"] != want.modes["char"] {
		t.Errorf("setup flags --chars & --bytes: want %v, got %v", want.modes["char"], got.modes["char"])
	}
}

func TestRunAppFlagCharAndWord(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"char": true,
			"word": true,
		},
	}

	os.Args = []string{"test", "--chars", "--words"}
	got, _ := setup()

	if got.modes["word"] != want.modes["word"] {
		t.Errorf("setup flags --chars & --bytes: want %v, got %v", want.modes["word"], got.modes["word"])
	}

	if got.modes["char"] != want.modes["char"] {
		t.Errorf("setup flags --chars & --bytes: want %v, got %v", want.modes["char"], got.modes["char"])
	}
}

func TestRunAppFlagByteCharWordAndLine(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"byte": true,
			"char": true,
			"line": true,
			"word": true,
		},
	}

	os.Args = []string{"test", "--bytes", "--chars", "--lines", "--words"}
	got, _ := setup()

	if got.modes["word"] != want.modes["word"] {
		t.Errorf("setup flags --chars & --bytes: want %v, got %v", want.modes["word"], got.modes["word"])
	}

	if got.modes["char"] != want.modes["char"] {
		t.Errorf("setup flags --chars & --bytes: want %v, got %v", want.modes["char"], got.modes["char"])
	}
}
