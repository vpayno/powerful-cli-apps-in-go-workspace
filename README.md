# "Powerful Cli Applications in Go" Workspace

[![Go Report Card](https://goreportcard.com/badge/github.com/vpayno/powerful-cli-apps-in-go-workspace)](https://goreportcard.com/report/github.com/vpayno/powerful-cli-apps-in-go-workspace)
[![CodeFactor](https://www.codefactor.io/repository/github/vpayno/powerful-cli-apps-in-go-workspace/badge)](https://www.codefactor.io/repository/github/vpayno/powerful-cli-apps-in-go-workspace)
[![Maintainability](https://api.codeclimate.com/v1/badges/43c8f7b58097ca3fa1ec/maintainability)](https://codeclimate.com/github/vpayno/powerful-cli-apps-in-go-workspace/maintainability)
[![codebeat badge](https://codebeat.co/badges/84438134-87ef-48d5-b388-0673c7e20db4)](https://codebeat.co/projects/github-com-vpayno-powerful-cli-apps-in-go-workspace-main)

[![Test Coverage](https://api.codeclimate.com/v1/badges/43c8f7b58097ca3fa1ec/test_coverage)](https://codeclimate.com/github/vpayno/powerful-cli-apps-in-go-workspace/test_coverage)
![Code2Test Ratio](./reports/.octocov-ratio.svg?branch=main)

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![Go version in go.mod](https://img.shields.io/github/go-mod/go-version/vpayno/powerful-cli-apps-in-go-workspace.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/blob/main/go.mod)
[![Version](https://badge.fury.io/gh/vpayno%2Fpowerful-cli-apps-in-go-workspace.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/tags)

[![Go Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/go.yml)
[![CodeQL Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/codeql-analysis-go.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/codeql-analysis-go.yml)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B33315%2Fgithub.com%2Fvpayno%2Fpowerful-cli-apps-in-go-workspace.svg?branch=main&type=shield)](https://app.fossa.com/projects/custom%2B33315%2Fgithub.com%2Fvpayno%2Fpowerful-cli-apps-in-go-workspace?ref=badge_shield)

[![Bash Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/bash.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/bash.yml)
[![Git Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/git.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/git.yml)
[![Link Check Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/links.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/links.yml)
[![Woke Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/woke.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/woke.yml)

[![Spelling Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/misspell.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/misspell.yml)
[![Json Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/json.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/json.yml)
[![Yaml Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/yaml.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/yaml.yml)
[![Markdown Workflow](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/markdown.yml/badge.svg?branch=main)](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/actions/workflows/markdown.yml)

This isn't a "real" project in the sense that it accepts PRs or should be used or forked by anyone as a real application.

This is my workspace for learing concepts from the book.

This is also my "notebook" on how to do things in Go or how to manage a Go project.

## Book Info

- [Website](https://pragprog.com/titles/rggo/powerful-command-line-applications-in-go/)

## Versions

Versions tags have two formats:

- `v[0-9]+`
- `v[0-9]+.[0-9]+.[0-9]+`

The short version is used by some tools to mean any X version.
The short tag is applied at the same time when the major number is incremented.

The long version has three parts:

- major: it's the chapter number minus 1
- minor: bumped when a new feature of the chapter's application is added, sometimes bumped when significant changes are made during a chapter.
- patch: application fixes and other changes

Once a new chapter is started, changes to a previous chapter need to be hot-fixed to the last tag of that chapter and to the HEAD of the repo.
If it's a new feature or fix of the chapter's application, it should be tagged in both locations.
It may seem like unnecessary work but it's good practice for real applications that have to be maintained over long periods of time.
The hot-fix can be skipped if it's using processes learned in the current chapter since the changes are related to only that chapter.

## Commit Messages

Commit messages are used to create change logs.

For this and many other reasons, commit messages:
    - must only contain a single type of change.
    - must never be squashed together.
    - must have properly formatted commit messages.
    - should be in the present tense.

As for the single type of change, it can include a lot of related changes to the same form:
    - formatting changes
    - linter recommended fixes
    - documentation (docs, comments, examples, etc.)
    - adding/removing a feature
    - fixing a feature

That way when a PR is reviewed, one can click on each commit to see what changed.
If you mix code formatting changes with a fix to an existing feature, it makes it a lot harder to see what changed.

Including multiple changes into a commit or squashing them also makes it more difficult to bisect the code at a later date.
What do you do when you find that the bug was introduced in a commit that includes hundreds or thousands of formatting changes and a bug fix?
If they were two separate commits, it would be easy to tell if it was the formatting changes or the bug fix that broke the code.

If you find a bug that tests missed while updating documentation, make that bug fix and doc/test update a separate commit instead of hiding it in another commit.

If in doubt about what categories may apply to an existing file or what commit messages look like, you can use

```text
git log ./path/to/file
```

or

```text
tig ./path/to/file
```

to browse all the commit messages associated with that file.

It's always best to check instead of just guessing.

Commit messages can also include more information in their bodies help others understand the changes.
    - What error was prompted the fix?
    - How to test for the error.

<details>
    <summary><h3>Commit message formats:</h3></summary>

```text
category: short message stating what changed
```

```text
category(subcategory): short message stating what changed
```

#### Categories:

- build: things related to the build system
- chore: catch all for project chores that don't fit in the other categories
- ci: things related to the ci system
- doc: things related to project documentation
- feat: new features
- fix: fixes to features
- release: release related changes
- site: GitHub pages related changes
- test: things related to project testing

#### Subcategories:

Categories that can be further subdivided, like build and ci, can have many subcategories.

- build(deps): dependency changes (usually version bumps)
- build(go): build system updates (native to the language)
- build(make): make related changes
- build(tools): build system tooling updates

- ci(bash): CI workflow for BASH checks
- ci(codeclimate): CI workflow for CodeClimate checks
- ci(codeql): CI workflow for GitHub CodeQL checks
- ci(fossa): CI workflow for FOSSA (license) checks
- ci(gha): CI workflow for GitHub Actions checks
- ci(git): CI workflow for Git related checking
- ci(goreleaser): CI workflow for Go release bulding and publishing
- ci(go): CI workflow for Golang checking, testing
- ci(hugo): CI workflow for GitHub Pages with Hugo build and deploy
- ci(links): CI workflow for the link checker
- ci(security): CI workflow for security checking
- ci(tools): CI workflow for generic CI tooling
- ci(woke): CI workflow for running the Woke checking

</details>

## Change Log

The script `./scripts/git-release-notes` can be used to extract the change log from the commit message summaries/titles.

Right now only one type of change log is supported, the developer version. It's composed of all the commit messages in the repo.

The user change log, will only include `doc`, `feat`, `fix` and `release` commits.

The script `././scripts/git-release-notes` can be used to create change logs for the whole repo or for just a single tag.

- Single tag:

```text
./scripts/git-release-notes v1.2.3
```

- All tags:

```text
./scripts/git-release-notes all
```

[Change Log](./CHANGELOG.md)

## Creating a Release

The script `./tag-release` can be used to create a release commit which contains

- updates to CHANGELOG.md
- version bump to ./cmd/*/.version.txt

The commit message and the annotated tag also contain the change log for the new release tag.

To create a new release, from the *main* branch or from a hot-fix branch, run:

```text
./tag-release v1.2.3 "added feature x"
```

These changes get a patch version bump:

- When a feature fix is merged into main.

These changes get a minor version bump:

- When a new feature is merged into main.
- When build system updates are added.
- When dependency or Go version bumps are made.

These changes get a major version bump:

- When a new chapter starts (n-1).

Artifact releases, to save space, will only be created at the end of a chapter.

## Go Version Auto Updates

The project as a scheduled job that runs every Wednesday that looks for a new major version of Golang.
It won't update to the next version until after the first patch release exists.

For example, when Golang `1.20` (`1.20.0`) is released, the script won't update the version in `go.mod` to `1.20` until version `1.20.1` is released.

When it successfully creates a Golang version bump commit, with the commit message `build(go): bump golang version to major.minor`,
it will also create a PR titled `Bump Golang Version to major.minor`.

If the `Go Workflow` checks pass (which will be triggered by the change to `go.mod`), the pull-request will be auto-approved.
Still working on auto-merging.

For this to work with a high degree of confidence, we need

- 100% [test coverage](https://codeclimate.com/github/vpayno/powerful-cli-apps-in-go-workspace) and
- dependencies need to be tested without mocking them out of existence.

To trigger a run manually run:

```text
gh workflow run 'Golang Version Bump Workflow'
```

To watch the workflow run:

```text
gh run watch $(gh run list --workflow "Golang Version Bump Workflow" --limit 1 --json 'databaseId' --jq '.[].databaseId')
```

## Dependabot

Dependabot configuration: [.github/dependabot.yml](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/blob/main/.github/dependabot.yml)

[Dependabot Status](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/network/updates)

Dependabot runs once a week, early Monday mornings, and updates dependencies as needed.
When dependencies are updated, Dependabot will open a new PR with the updates.
GitHub actions/workflows that run in PRs when Go files change will test the new dependencies automatically.

For this to work with a high degree of confidence, we need

- 100% [test coverage](https://codeclimate.com/github/vpayno/powerful-cli-apps-in-go-workspace) and
- dependencies need to be tested without mocking them out of existence.

Note: When CI secrets are added, they also need to be added as [Dependabot secrets](https://github.com/vpayno/powerful-cli-apps-in-go-workspace/settings/secrets/dependabot) for any workflow that will run for a Dependabot PR.

After a Dependabot PR is opened, two human actions are required:

- Start a review and approve the PR.
- In the review comment add the string `@dependabot merge` to automatically merge the PR after the CI checks have passed.

Only use `@dependabot merge`, in a comment in the PR, to automatically merge the PR after the CI checks have passed.
Don't merge Dependabot PRs manually using the `Merge Pull Request` button on the PR.
Only GitHub can sign Dependabot commits and if you merge the PR, the commits will be unsigned and unverified.
Commits can be signed and merged by hand on the CLI; However, it's easier to just ask Dependabot to merge the PR.

<details>
    <summary><hr4>Dependabot commands and options</hr4></summary>

You can trigger Dependabot actions by commenting on the PR:

- `@dependabot rebase` will rebase this PR
- `@dependabot recreate` will recreate this PR, overwriting any edits that have been made to it
- `@dependabot merge` will merge this PR after your CI passes on it
- `@dependabot squash and merge` will squash and merge this PR after your CI passes on it
- `@dependabot cancel merge` will cancel a previously requested merge and block automerging
- `@dependabot reopen` will reopen this PR if it is closed
- `@dependabot close` will close this PR and stop Dependabot recreating it. You can achieve the same result by closing it manually
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)

</details>

## Build Releases

[Install GoReleaser](https://goreleaser.com/install/):

```text
go install github.com/goreleaser/goreleaser@latest
```

Build Linux, OSX and Windows binaries:

```text
make build-all
```

For instructions on how to install individual applications, check the applications section for `go install url@tag` command you need.

## Chapters|Apps

<details id=1>
    <summary><h3>Chapter 01: Your First Command-Line Program in Go - wordcount/wc</h3></summary>

#### How to Install *cli-wc*

I'm adding `cli-` prefix to the binaries so I don't replace the system version of `wc` with this one in my `PATH`.

Using `go install`

```text
go install github.com/vpayno/powerful-cli-apps-in-go-workspace/cmd/cli-wc@latest
```

or

```text
git clone https://github.com/vpayno/powerful-cli-apps-in-go-workspace.git
cd powerful-cli-apps-in-go-workspace
go generate
go install ./cmd/cli-wc/main.go
```

#### *cli-wc* Usage

```text
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

  <details><summary><h4><em>cli-wc</em> Examples</h4></summary>

- Show Version

```text
$ cli-wc --version
v0.3.4
```

```text
$ cli-wc --version --verbose

Word Count Version: v0.3.4
```

- Default Counts

```text
$ printf "%s\n" one two 😂 four five | cli-wc
      5       5      23

```

```text
$ printf "%s\n" one two 😂 four five | cli-wc --verbose
Word Count Version v0.2.1

      5 (line)       5 (word)      23 (byte)
```

- Count Words

```text
$ printf "%s\n" one two 😂 four five | cli-wc --words
5
```

```text
$ printf "%s\n" one two 😂 four five | cli-wc --words --verbose
Word Count Version v0.2.1

5 (word)
```

- Count Lines

```text
$ printf "%s\n" one two 😂 four five | cli-wc --lines
5
```

```text
$ printf "%s\n" one two 😂 four five | cli-wc -lines --verbose
Word Count Version v0.2.1

5 (line)
```

- Count Bytes

```text
$ printf "%s\n" one two 😂 four five | cli-wc --bytes
23
```

```text
$ printf "%s\n" one two 😂 four five | cli-wc --bytes --verbose
Word Count Version v0.2.1

23 (byte)
```

- Count Chars

```text
$ printf "%s\n" one two 😂 four five | cli-wc --chars
20
```

```text
$ printf "%s\n" one two 😂 four five | cli-wc --chars --verbose
Word Count Version v0.2.1

20 (char)
```

- Max Line Length

```text
$ printf "0123456\n0123456789\n😂\n01234\n1234567890123\n\n" | ./cli-wc --bytes --chars --lines --words --max-line-length
      6       5      42      45      13
```

```text
$ printf "0123456\n0123456789\n😂\n01234\n1234567890123\n\n" | ./cli-wc --bytes --chars --lines --words --max-line-length --verbose
Word Count Version v0.3.0

      6 (line)       5 (word)      42 (char)      45 (byte)      13 (length)
```

- All Counts

```text
$ printf "\n%s\n" one two 😂 four five | cli-wc --bytes --chars --lines --words --max-line-length
      2       1      12      12      10
```

```text
$ printf "\n%s\n" one two 😂 four five | cli-wc --bytes --chars --lines --words --max-line-length --verbose
Word Count Version v0.3.0

      2 (line)       1 (word)      12 (char)      12 (byte)      10 (length)
```

  </details>

</details>
