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

	got, err := setup()

	if err != nil {
		t.Error(err)
	}

	if want.lineMode != got.lineMode {
		t.Errorf("setup() returned the wrong line mode value. want: %v, got %v", want.lineMode, got.lineMode)
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
		lineMode: false,
	}

	want := 5
	got := getCount(b, conf)

	if want != got {
		t.Errorf("Expected word count %d, got %d.\n", want, got)
	}
}

func TestGetLineCount(t *testing.T) {
	b := bytes.NewBufferString("one\ntwo\nthree\nfour\nfive\n")

	conf := config{
		lineMode: true,
	}

	want := 5
	got := getCount(b, conf)

	if want != got {
		t.Errorf("Expected line count %d, got %d.\n", want, got)
	}
}

func TestGetByteCount(t *testing.T) {
	b := bytes.NewBufferString("0123456789")

	conf := config{
		byteMode: true,
	}

	want := 10
	got := getCount(b, conf)

	if want != got {
		t.Errorf("Expected byte count %d, got %d.\n", want, got)
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

	count := 5

	conf := config{
		lineMode:    false,
		verboseMode: true,
	}

	logVerbose = true

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("word count: %d\n", count)

	// Run the function who's output we want to capture.
	showCount(count, conf)

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

	logVerbose = false
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

	count := 5

	conf := config{
		lineMode:    true,
		verboseMode: true,
	}

	logVerbose = true

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("line count: %d\n", count)

	// Run the function who's output we want to capture.
	showCount(count, conf)

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

	logVerbose = false
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

	count := 5

	conf := config{
		byteMode:    true,
		verboseMode: true,
	}

	logVerbose = true

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("byte count: %d\n", count)

	// Run the function who's output we want to capture.
	showCount(count, conf)

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

	logVerbose = false
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

	count := 5

	conf := config{
		lineMode:    false,
		verboseMode: true,
	}

	logVerbose = false

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%d\n", count)

	// Run the function who's output we want to capture.
	showCount(count, conf)

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

	count := 5

	conf := config{
		lineMode:    true,
		verboseMode: true,
	}

	logVerbose = false

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%d\n", count)

	// Run the function who's output we want to capture.
	showCount(count, conf)

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

	count := 5

	conf := config{
		byteMode:    true,
		verboseMode: true,
	}

	logVerbose = false

	// It's a silly test but I need the practice.
	want := fmt.Sprintf("%d\n", count)

	// Run the function who's output we want to capture.
	showCount(count, conf)

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

func TestRunAppFlagByteAndLine(t *testing.T) {
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

	want := "Error: -b (byte count mode) and -l (line count mode) can't be used at the same time\n"

	os.Args = []string{"test", "-b", "-l"}
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
