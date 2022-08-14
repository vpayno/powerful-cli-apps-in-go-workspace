[![Go Report Card](https://goreportcard.com/badge/github.com/vpayno/powerful-cli-apps-in-go-workspace)](https://goreportcard.com/report/github.com/vpayno/powerful-cli-apps-in-go-workspace)
[![Go Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/go.yml/badge.svg)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/go.yml)
[![Bash Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/bash.yml/badge.svg)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/bash.yml)
[![Git Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/git.yml/badge.svg)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/git.yml)
[![Link Check Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/links.yml/badge.svg)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/links.yml)
[![Woke Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/woke.yml/badge.svg)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/woke.yml)

![Coverage](./reports/.octocov-coverage.svg)
![Code2Test Ratio](./reports/.octocov-ratio.svg)

# "Powerful Cli Applications in Go" Workspace

This isn't a "real" project in the sense that it accepts PRs or should be used or forked by anyone as a real application.

This is my workspace for learing concepts from the book.

This is also my "notebook" on how to do things in Go or how to manage a Go project.

## Book Info

- [Website](https://pragprog.com/titles/rggo/powerful-command-line-applications-in-go/)

## Chapters|Apps

### Chapter 01 - Your First Command-Line Program in Go - wordcount/wc

#### How to Install

I'm adding `cli-` prefix to the binaries so I don't replace the system version of `wc` with this one in my `PATH`.

Using `go install`

```
$ go install github.com/vpayno/powerful-cli-apps-in-go-workspace/cmd/cli-wc@latest
```

or

```
$ git clone https://github.com/vpayno/powerful-cli-apps-in-go-workspace.git
$ cd powerful-cli-apps-in-go-workspace
$ go generate
# go install ./cmd/cli-wc/cli-wc.go
```

#### Usage

```
$ cli-wc --help
Usage: cli-wc [OPTION]...

Print newline, word, and byte counts for stdin input.

A word is a non-zero-length sequence of characters delimited by white space.

The options below may be used to select which counts are printed, always in
the following order: newline, word, character, byte.

Options:
  -c, --bytes            print the byte counts
  -m, --chars            print the character counts
  -l, --lines            print the newline counts
  -w, --words            print the word counts
  -h, --help             display this help and exit
  -v, --version          output version information and exit
  -V, --verbose          verbose mode
```

#### Examples

- Show Version

```
$ cli-wc --version

Word Count Version: v0.2.1
```

- Default Counts

```
$ printf "%s\n" one two 😂 four five | cli-wc
      5       5      23

```

```
$ printf "%s\n" one two 😂 four five | cli-wc --verbose
Word Count Version v0.2.1

      5 (line)       5 (word)      23 (byte)
```

- Count Words

```
$ printf "%s\n" one two 😂 four five | cli-wc --words
5
```

```
$ printf "%s\n" one two 😂 four five | cli-wc --words --verbose
Word Count Version v0.2.1

5 (word)
```

- Count Lines

```
$ printf "%s\n" one two 😂 four five | cli-wc --lines
5
```

```
$ printf "%s\n" one two 😂 four five | cli-wc -lines --verbose
Word Count Version v0.2.1

5 (line)
```

- Count Bytes

```
$ printf "%s\n" one two 😂 four five | cli-wc --bytes
23
```

```
$ printf "%s\n" one two 😂 four five | cli-wc --bytes --verbose
Word Count Version v0.2.1

23 (byte)
```

- Count Chars

```
$ printf "%s\n" one two 😂 four five | cli-wc --chars
20
```

```
$ printf "%s\n" one two 😂 four five | cli-wc --chars --verbose
Word Count Version v0.2.1

20 (char)
```

- All Counts

```
$ printf "%s\n" one two 😂 four five | cli-wc --bytes --chars --lines --words
      5       5      20      23
```

```
$ printf "%s\n" one two 😂 four five | cli-wc --bytes --chars --lines --words --verbose
Word Count Version v0.2.1

      5 (line)       5 (word)      20 (char)      23 (byte)
```
