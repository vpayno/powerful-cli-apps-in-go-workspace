---
title: "Ch01 Wc"
date: 2022-09-05T14:15:31-07:00
draft: true
---

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

- Max Line Length

```
$ printf "0123456\n0123456789\n😂\n01234\n1234567890123\n\n" | ./cli-wc --bytes --chars --lines --words --max-line-length
      6       5      42      45      13
```

```
$ printf "0123456\n0123456789\n😂\n01234\n1234567890123\n\n" | ./cli-wc --bytes --chars --lines --words --max-line-length --verbose
Word Count Version v0.3.0

      6 (line)       5 (word)      42 (char)      45 (byte)      13 (length)
```

- All Counts

```
$ printf "\n%s\n" one two 😂 four five | cli-wc --bytes --chars --lines --words --max-line-length
      2       1      12      12      10
```

```
$ printf "\n%s\n" one two 😂 four five | cli-wc --bytes --chars --lines --words --max-line-length --verbose
Word Count Version v0.3.0

      2 (line)       1 (word)      12 (char)      12 (byte)      10 (length)
```
