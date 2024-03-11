// Package apptodo is the module with the cli logic for the wc main application.
package apptodo

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

// item struct represents a ToDo item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of ToDo items
type List []item

// Add creates a new Todo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete method marks add ToDo item as completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func showBanner() {
	fmt.Println(metadata.name + " Version " + metadata.version)
	fmt.Println()
}

// RunApp is called my the main function. It's basically the main function of the app.
func RunApp() {
	conf, err := setup()
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)

		return
	}

	if conf.versionMode {
		showVersion(conf)

		return
	}

	if conf.verboseMode {
		showBanner()
	}

	counts := getCounts(os.Stdin, conf)

	showCount(counts, conf)
}

// Usage prints the command-line usage help message.
func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
}

func setup() (config, error) {
	var byteFlagPtr bool

	var charFlagPtr bool

	var lengthFlagPtr bool

	var lineFlagPtr bool

	var wordFlagPtr bool

	flagSet.BoolVar(&byteFlagPtr, "c", false, "print the byte counts")
	flagSet.BoolVar(&byteFlagPtr, "bytes", false, "print the byte counts")
	flagSet.BoolVar(&charFlagPtr, "m", false, "print the character counts")
	flagSet.BoolVar(&charFlagPtr, "chars", false, "print the character counts")
	flagSet.BoolVar(&lengthFlagPtr, "L", false, "print the maximum display width")
	flagSet.BoolVar(&lengthFlagPtr, "max-line-length", false, "print the maximum display width")
	flagSet.BoolVar(&lineFlagPtr, "l", false, "print the newline counts")
	flagSet.BoolVar(&lineFlagPtr, "lines", false, "print the newline counts")
	flagSet.BoolVar(&wordFlagPtr, "w", false, "print the word counts")
	flagSet.BoolVar(&wordFlagPtr, "words", false, "print the word counts")

	var verboseFlagPtr bool

	var versionFlagPtr bool

	flagSet.BoolVar(&verboseFlagPtr, "v", false, "verbose mode")
	flagSet.BoolVar(&verboseFlagPtr, "verbose", false, "verbose mode")
	flagSet.BoolVar(&versionFlagPtr, "V", false, "output version information and exit")
	flagSet.BoolVar(&versionFlagPtr, "version", false, "output version information and exit")

	flagSet.Usage = Usage

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		return config{}, err
	}

	conf := config{
		flags:       defaultFlags,
		verboseMode: verboseFlagPtr,
		versionMode: versionFlagPtr,
		modes:       defaultModes,
	}

	if byteFlagPtr || charFlagPtr || lengthFlagPtr || lineFlagPtr || wordFlagPtr {
		conf.flags = byteFlag | charFlag | wordFlag | lineFlag | lengthFlag
	}

	conf.modes = map[string]bool{
		"byte":   conf.flags&byteFlag != 0,
		"char":   conf.flags&charFlag != 0,
		"length": conf.flags&lengthFlag != 0,
		"line":   conf.flags&lineFlag != 0,
		"word":   conf.flags&wordFlag != 0,
	}

	return conf, nil
}

func getCounts(r io.Reader, conf config) results {
	counts := results{
		"byte":   0,
		"char":   0,
		"length": 0,
		"line":   0,
		"word":   0,
	}

	reader := bufio.NewReader(r)

	for {
		text, err := reader.ReadString('\n')

		if conf.modes["byte"] {
			counts["byte"] += len(text)
		}

		if conf.modes["char"] {
			counts["char"] += utf8.RuneCountInString(text)
		}

		if conf.modes["word"] {
			counts["word"] += len(strings.Fields(text))
		}

		if conf.modes["length"] {
			maxLength := utf8.RuneCountInString(strings.TrimSuffix(text, "\n"))

			if maxLength > counts["length"] {
				counts["length"] = maxLength
			}
		}

		if conf.modes["line"] {
			counts["line"]++

			if err == io.EOF {
				counts["line"]--
			}
		}

		if err != nil {
			break
		}
	}

	return counts
}

func showCount(counts results, conf config) {
	first := true
	fieldSize := "7"

	var modeCount int

	for _, v := range conf.modes {
		if v {
			modeCount++
		}
	}

	if modeCount == 1 {
		fieldSize = "0"
	}

	for _, mode := range printOrder {
		if !conf.modes[mode] {
			continue
		}

		if !first {
			fieldSize = "8"
		}

		first = false

		var label string

		if conf.verboseMode {
			label = " (" + mode + ")"
		}

		fmt.Printf("%"+fieldSize+"d%s", counts[mode], label)
	}

	fmt.Println()
}
