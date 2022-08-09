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
$ ./scripts/go-cmd run ./... -h
go generate ./...
go generate: creating /home/vpayno/git_vpayno/powerful-cli-apps-in-go-workspace/cmd/cli-wc/./.version.txt

v0.0.0
v0.0.0-8-g275e726
275e72657a4241eaddc5ce90bbb1aaa0c6a0289b
2022-08-08T07:00:12Z


real	0m0.144s
user	0m0.102s
sys	0m0.140s

Usage of /tmp/go-build1277001378/b001/exe/cli-wc:
  -V	show the app version
  -b	count bytes instead of words
  -l	count lines instead of words
  -v	verbose mode

real	0m0.314s
user	0m0.301s
sys	0m0.290s
go run ./... -h

git restore ./cmd/cli-wc/.version.txt

real	0m0.004s
user	0m0.003s
sys	0m0.003s
```

#### Examples

- Show Version

```
$ go generate ./...
go generate: creating /home/vpayno/git_vpayno/powerful-cli-apps-in-go-workspace/cmd/cli-wc/./.version.txt

v0.0.0
v0.0.0-8-g275e726
275e72657a4241eaddc5ce90bbb1aaa0c6a0289b
2022-08-08T07:25:49Z

$ go build ./cmd/cli-wc/cli-wc.go

$ printf "%s\n" one two three four five | ./cli-wc -V

Word Count Version: v0.0.0

git version: v0.0.0-8-g275e726
   git hash: 275e72657a4241eaddc5ce90bbb1aaa0c6a0289b
 build time: 2022-08-08T07:25:49Z
```

- Count Words

```
$ printf "%s\n" one two three four five | ./cli-wc
5
```

```
$ printf "%s\n" one two three four five | ./cli-wc -v
Word Count Version v0.0.0

word count: 5
```

- Count Lines

```
$ printf "%s\n" one two three four five | ./cli-wc -l
5
```

```
$ printf "%s\n" one two three four five | ./cli-wc -l -v
Word Count Version v0.0.0

line count: 5
```

- Count Bytes

```
$ printf "%s\n" one two three four five | ./cli-wc -b
19
```

```
$ printf "%s\n" one two three four five | ./cli-wc -b -v
Word Count Version v0.0.0

byte count: 24
```
