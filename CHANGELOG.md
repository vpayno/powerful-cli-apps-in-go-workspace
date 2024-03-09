# "Powerful Cli Applications in Go" Workspace Release Notes

<details open>
    <summary>
<h2> [2024-03-09] Release v0.7.16: bump golang version to 1.22
</h2>
    </summary>

### build(go)

- bump golang version to 1.22

### ci(tools)

- fix to file detection and version string editing
- update version bump commit title
- use git to get latest golang versions, minor code improvements

</details>

<details>
    <summary>
<h2> [2024-01-07] Release v0.7.15: gitlab release fixes
</h2>
    </summary>

### ci(gitlab)

- dotenv doesn't support multiline variables

</details>

<details>
    <summary>
<h2> [2024-01-07] Release v0.7.13: gitlab release fixes
</h2>
    </summary>

### build(tools)

- make octocov happy in gitlab

### ci(gitlab)

- add git dependency to build job

</details>

<details>
    <summary>
<h2> [2024-01-07] Release v0.7.12: experiment with gitlab releases
</h2>
    </summary>

### ci(gitlab)

- fix release url and notes

</details>

<details>
    <summary>
<h2> [2024-01-07] Release v0.7.11: experiment with gitlab releases
</h2>
    </summary>

### ci(gitlab)

- release fixes

</details>

<details>
    <summary>
<h2> [2024-01-06] Release v0.7.10: experiment with gitlab releases
</h2>
    </summary>

### ci(gitlab)

- allow manual/web execution

</details>

<details>
    <summary>
<h2> [2024-01-06] Release v0.7.9: experiment with go workspaces
</h2>
    </summary>

### chore

- setup go workspaces for the project and the ci code

### ci(dagger-ci)

- show pwd and directory listing in lint, test & build jobs
- stop excluding ./ci workspace member in lint & test jobs

### ci(gitlab)

- add artifact creation and retention to build step
- add release step
- add stages to lint, test & built steps
- define pipeline stages

### ci(go)

- fix revive install error
- format yaml file

</details>

<details>
    <summary>
<h2> [2024-01-03] Release v0.7.8: test new create-release-files wrapper script
</h2>
    </summary>

### ci(dagger-go)

- move release asset creation to a wrapper script

</details>

<details>
    <summary>
<h2> [2024-01-02] Release v0.7.7: add more build targets
</h2>
    </summary>

### chore

- add /ci/*/.version.txt to .gitignore

### ci(dagger-go)

- add more build targets

### ci(git)

- add 'site' commit message type

### site

- change download button to releases

</details>

<details>
    <summary>
<h2> [2023-12-31] Release v0.7.6: fix workflow to download git notes
</h2>
    </summary>

### ci(dagger-go)

- fetch git notes when releasing

</details>

<details>
    <summary>
<h2> [2023-12-31] Release v0.7.5: one more tag-release fix
</h2>
    </summary>

### ci(tools)

- more tag-release fixes

</details>

<details>
    <summary>
<h2> [2023-12-31] Release v0.7.4: fix pushing of git notes when releasing
</h2>
    </summary>

### ci(tools)

- fix pushing of git notes and tags

</details>

<details>
    <summary>
<h2> [2023-12-31] Release v0.7.3: release trigger fixes, use git notes for release message
</h2>
    </summary>

### ci(dagger-go)

- fix git tag trigger pattern
- use git notes for release message body

### ci(tools)

- add a git note when tagging

</details>

<details>
    <summary>
<h2> [2023-12-31] Release v0.7.2: add tag push and manual dagger-go workflow triggers
</h2>
    </summary>

### ci(dagger-go)

- add manual trigger
- add tag push trigger

</details>

<details>
    <summary>
<h2> [2023-12-31] Release v0.7.15: gitlab release fixes
</h2>
    </summary>

### ci(gitlab)

- dotenv doesn't support multiline variables

</details>

<details>
    <summary>
<h2> [2023-11-16] Release v0.7.0: bump golang version to 1.21
</h2>
    </summary>

### build(go)

- bump golang version to 1.21
- bump golang version to 1.21.0

### ci(tools)

- fix next version detection bug in scripts/go-version-bump
- fix next version detection bug in scripts/go-version-bump - continued

### revert

- bump golang version to 1.21.0"

</details>

<details>
    <summary>
<h2> [2023-06-05] Release v0.6.7: dependabot testify update
</h2>
    </summary>

### build(deps)

- bump github.com/stretchr/testify from 1.8.3 to 1.8.4

</details>

<details>
    <summary>
<h2> [2023-05-23] Release v0.6.6: fix lint and formatting
</h2>
    </summary>

### chore

- clean up format errors
- clean up lint errors

### fix

- fixing tests that broke after formatting

</details>

<details>
    <summary>
<h2> [2023-05-23] Release v0.6.5: fix git workflow
</h2>
    </summary>

### ci(git)

- use different action and update config file

</details>

<details>
    <summary>
<h2> [2023-05-22] Release v0.6.4: bump github.com/stretchr/testify from 1.8.2 to 1.8.3
</h2>
    </summary>

### build(deps)

- bump github.com/stretchr/testify from 1.8.2 to 1.8.3

### ci

- add commitlint configuration file

### ci(golang-bump)

- fix checks for pr number

### ci(go)

- silence invalid shellcheck warnings

</details>

<details>
    <summary>
<h2> [2023-03-03] Release v0.6.3: bump golang from 1.19 to 1.20
</h2>
    </summary>

### build(go)

- bump golang version to 1.20

### ci(git)

- change job name from block-fixup to git-commit-message-check

### ci(go)

- add go-sumtype check
- fix go-consistent error after updating to go 1.20.1

</details>

<details>
    <summary>
<h2> [2023-03-01] Release v0.6.2: bump github.com/stretchr/testify from 1.8.1 to 1.8.2
</h2>
    </summary>

### build(deps)

- bump github.com/stretchr/testify from 1.8.1 to 1.8.2

</details>

<details>
    <summary>
<h2> [2022-11-11] Release v0.6.1: ci chore work
</h2>
    </summary>

### build(tools)

- generate coverage.xml

### chore

- add node package json files to git ignore
- add proselint config

### ci(codacy-go)

- add initial codacy code coverage check
- remove workflow

### ci(draft-check)

- move git fixup check to new workflow

### ci(git)

- new workflow that checks for convensional commit messages
- run on push to main & develop branches

### doc

- set git workflow badge to track main branch

</details>

<details>
    <summary>
<h2> [2022-11-10] Release v0.6.0: bump golang version to 1.19
</h2>
    </summary>

### build(go)

- bump golang version to 1.19

### ci(golang-bump)

- fix merge instructions comment
- fix pr comment generation
- put tag-release command comment in a code block

</details>

<details>
    <summary>
<h2> [2022-11-10] Release v0.5.4: give up on the auto-tag and auto-merge dream for now
</h2>
    </summary>

### ci(golang-bump)

- disable auto-merge command

</details>

<details>
    <summary>
<h2> [2022-11-10] Release v0.5.3: give up on the auto-tag and auto-merge dream for now
</h2>
    </summary>

### build(go)

- bump golang version to 1.19
- revert golang version to 1.18

### ci(golang-bump)

- don't auto-approve or auto-merge and add comments and summary with merge and tag commands

### ci(tools)

- check to see if the tag already exits before tagging a release

</details>

<details>
    <summary>
<h2> [2022-11-10] Release v0.5.2: sort changelog entries
</h2>
    </summary>

### ci(tools)

- sort section entries

</details>

<details>
    <summary>
<h2> [2022-11-10] Release v0.5.1: more gh auto-bump golang version in go.mod experiments
</h2>
    </summary>

### build(go)

- bump golang version to 1.19
- revert golang version to 1.18

### ci(golang-bump)

- change name of pr create steps and remove auto-approved footnote
- use gh to enable pr auto-merge

</details>

<details>
    <summary>
<h2> [2022-11-10] Release v0.5.0: bump golang version to 1.19
</h2>
    </summary>

### build(go)

- bump golang version to 1.19
- downgrade golang version to 1.18

### chore

- update todo list

### ci(golang-bump)

- switch to using gh to create pr

</details>

<details>
    <summary>
<h2> [2022-11-09] Release v0.4.1: add govulncheck check
</h2>
    </summary>

### ci(go)

- add govulncheck check

</details>

<details>
    <summary>
<h2> [2022-11-09] Release v0.4.0: bump golang version to 1.19
</h2>
    </summary>

### build(go)

- bump golang version to 1.19

### ci(golang-bump)

- add steps to run tag-release

### ci(tools)

- don't edit markdown files when bumping the golang version number

</details>

<details>
    <summary>
<h2> [2022-11-08] Release v0.3.9: ci fixes, add codebeat ci job, and add job that auto updates golang version
</h2>
    </summary>

### build(go)

- bump golang version to 1.19
- knock down the Golang version from 1.19 to 1.18 to test the bump up script
- set golang version to 1.18

### ci(codebeat)

- add codebeat go code coverage workflow

### ci(codeql)

- use golang version from go.mod

### ci(golang-bump)

- add scheduled workflow that bump golang version
- auto-approve after go checks pass
- don't run on push
- fix errors from runs after version bump

### ci(go)

- fix shellcheck issue with which command
- use go.mod for go version

### ci(json)

- add step to check golang, npm, reviewdog, jsonlint versions
- use golang version from go.mod

### ci(markdown)

- use golang version from go.mod

### ci(tools)

- add script to bump the golang version to latest version
- add yamllint check to gha-checks

### ci(yaml)

- add step to check yamllint version

### doc

- add information about automatic golang version bumps
- fix go version badge

</details>

<details>
    <summary>
<h2> [2022-11-05] Release v0.3.8: ci updates and new ci checks
</h2>
    </summary>

### chore

- add new tasks to todo list and sort it
- update todo list

### chore(markdownlint)

- fix change log markdown lint issues

### chore(yamllint)

- clean up comments lint warnings
- clean up document-start lint errors
- clean up truthy lint warnings
- clean up 'brackets' lint errors
- clean up 'indentation' lint errors

### ci

- clean up comments/headers for all workflows
- set workflows to run on push and prs

### ci(bash)

- add develop branch to push list

### ci(codeql)

- fix stage1 dependency name
- fix workflow file name

### ci(fossa)

- fix stage1 dependency name

### ci(gha)

- add -oneline argument to actionlint
- fix comment typo

### ci(git)

- revert to only running on pull-requests

### ci(go)

- change reviewdog reporter from github-pr-check to github-check
- exclude changelog.md and pages-gh from misspell results
- fix CodacyCoverageReporter error
- fix checkout-pr-branch so it can run on main or develop
- fix duplicate step id
- fix typos
- generate coverage.xml file
- only generate coverage.xml for linux build
- output data getting sent to reviewdog
- remove codacy coverage repoerter and badge
- remove misspell check
- upgrade bash on macos
- use correct input file for gocover-cobertura

### ci(json)

- add initial json check workflow

### ci(markdown)

- add initial markdownlint checks workflow

### ci(spelling)

- add spellcheck workflow

### ci(tools)

- fix release change log regeneration
- fix typo in tag-release
- redirect misspell stdout to stderr when generating change log
- run misspell on change log files after generating them

### ci(yaml)

- add yaml checks workflow
- set badge to use results from main branch

### doc

- add codacy badge to readme
- add dependabot notes to readme
- change heading structure of readme
- only report the git workflow badge status for pull-requests
- update badges to reflect state of main branch

### docs

- add commit message format information to readme
- add extra blank line before headings
- add release information to readme
- add version information to readme
- update build releases section
- update change log section in readme

### fix(codacy)

- add missing go package comment

</details>

<details>
    <summary>
<h2> [2022-10-24] Release v0.3.7: add repo name and section folds to change log
</h2>
    </summary>

### chore

- update todo list completed items and add new ones

### ci(tools)

- add repo name to change log
- add section folds to the change log
- the latest fold in the change log defaults to open

</details>

<details>
    <summary>
<h2> [2022-10-24] Release v0.3.6: release fixes and update change log when releasing
</h2>
    </summary>

### build(deps)

- bump github.com/stretchr/testify from 1.8.0 to 1.8.1

### chore

- update todo list completed items and add new ones

### ci(tools)

- add change log update to tag-release
- add rename release mode to release-short and add mode release-full to generate a complete change log at relase time
- fix logic error in tag-release
- fix typo in tag-release

### doc

- manually add v0.3.5 change log

</details>

<details>
    <summary>
<h2> [2022-10-23] Release v0.3.5: ci and doc updates, code clean up
</h2>
    </summary>

### build(tools)

- format bash scripts with shfmt

### ci(bash)

- change reviewdog reporter to github-pr-check
- set the default run shell to bash

### ci(codeql)

- fix comment
- fix set-output deprecation notice and convert from pwsh to bash
- fix typo in comment
- set the default run shell to bash

### ci(fossa)

- add fossa test step
- add --help and list-targets output
- rename action yaml file, fix set-output deprecation notice and convert from pwsh to bash
- set the default run shell to bash

### ci(gha)

- add git hub workflow linter action
- fix set-output deprecation notice and convert from pwsh to bash
- fix typo in comment
- set the default run shell to bash

### ci(git)

- set the default run shell to bash

### ci(go)

- add coverage info to job summary
- also run all the checks when the action changes
- enable tests on windows and macos
- fix comment
- fix set-output deprecation notice and convert from pwsh to bash
- resolve issues found by actionlint
- set the default run shell to bash
- temporary fix for misspell
- upload coverage reports

### ci(hugo)

- comment clean up

### ci(links)

- change version from v1.5.0 to v1
- set the default run shell to bash

### ci(tools)

- add current release notes script
- add github action check script
- add tag-release script for automating releases
- update git-release-notes to generate a full change log
- update git-release-notes to generate a release commit message

### ci(woke)

- change reviewdog reporter to github-pr-check
- set the default run shell to bash

### doc

- add change log for existing releases
- add collapsable sections for each chapter
- add collapsable section in ch01 for the examples
- update readme to show v0.3.4 version usage

### fix

- simplify setup() with bit flags

</details>

<details>
    <summary>
<h2> [2022-10-05] Release v0.3.4: fix --version and --version + --verbose usage
</h2>
    </summary>

### build(tools)

- update build and build-all steps and remove version step

### chore

- set nvm nodejs version to lts
- update todo list

### ci(fossa)

- add fossa workflow and badge
- run when the license file, go.sum or go.mod change

### ci(hugo)

- add hugo github pages workflow
- change upload source path to ./pages-gh/public/
- install docsy dependencies
- only run on pages-pr branch
- only run on pull-requests
- only run pr workflow when on branch pages-pr
- rename workflow file from hugo to github-pages-hugo
- run hugo command(s) from the pages-gh directory

### ci(links)

- disable badge checks on svgshare.com
- ignore css files
- ignore github.io links

### ci(woke)

- add configuration file and ignore ./site/

### doc

- add Go version badge to readme
- add codefactor badge to readme
- add linux, macos, windows badges to readme
- add made with go badge to readme

### fix

- only show extra version info when the --verbose flag is used

### site

- add about page
- add chapter 01 wc page
- add hugo theme docsy
- add initial index page
- add initial skeleton with chapter 01
- reset, move from site to pages-gh
- update chapter 01

### test

- add basic TestMain(m) functions to test dirs
- add test cases for cli arg flags
- rename testCase to dataTestCase

</details>

<details>
    <summary>
<h2> [2022-09-03] Release v0.3.3: ci, doc and chore updates
</h2>
    </summary>

### chore

- add wiki submodule
- update issue templates
- update todo list

### ci(codeclimate)

- add codeclimate code coverage reporting and badges

### ci(codeql)

- add initial CodeQL workflow
- add name to stage1-setup job

### ci(go)

- add names to jobs
- fix cc-test-reporter error
- only run when Go files change

### ci(security)

- add initial dependabot.yml configuration

### doc

- add codeql badge and create a new badge row for the main checks
- add code of conduct
- add goreleaser installation and artifact build instructions
- add latest version and release badges to readme
- add security policy
- move Go Report badge to 1st badge row
- reorder badges to row1:health, row2:version, row3:check_status

### test

- add unicode tests, test both short and long arg tests

</details>

<details>
    <summary>
<h2> [2022-08-30] Release v0.3.2: testing goreleaser
</h2>
    </summary>

### ci(goreleaser)

- add initial config file

</details>

<details>
    <summary>
<h2> [2022-08-22] Release v0.3.1: simplify version data
</h2>
    </summary>

### chore

- update todo list
- update todo list
- update todo list
- update todo list

### fix

- remove gitVersion and use version instead

</details>

<details>
    <summary>
<h2> [2022-08-14] Release v0.3.0: add --max-line-length option
</h2>
    </summary>

### feat

- add option to show the maximum line length

</details>

<details>
    <summary>
<h2> [2022-08-14] Release v0.2.1: fix count print order
</h2>
    </summary>

### chore

- comment clean up

### doc

- update readme with new cli-wc output

### fix

- fix count print order (line, word, char, byte)
- make byte, char, line and word flags plural
- simplify setup() and remove flag override/exclusions

### test

- make tests easier to ready by using long cli options

</details>

<details>
    <summary>
<h2> [2022-08-13] Release v0.2.0: newline, word, byte and char fixes
</h2>
    </summary>

### build(tools)

- add release summary to release commit

### ci(go)

- add go-consistent check
- add staticcheck check
- move most go install commands to their respective step
- run gocritic check after the other linters have run
- run golangci-lint check after the other linters have run
- run gosec check after the other linters have run
- run staticcheck check after the other linters have run
- run test and coverage checks after the other linters have run

### ci(make)

- fix make clean

### feat

- add long command-line options
- add usage help message
- add -r rune mode
- let -l, -w, -c, -b be used like they are in the coreutils wc cli

### fix

- add more test cases, fix bugs with getCounts()
- add -w argument and let -b and -l override -w
- change non-verbose output to match wc from coreutils
- change the help string for -V to match wc from coreutils
- change the help string for -l and -w to match wc from coreutils
- change "rune" to "char" to match coreutils wc
- change -b to -c to match wc from coreutils
- change -r to -m to match wc from coreutils
- include new lines in byte count
- make sure Usage() shows supported flags
- properly define flags.Parse() side-effect behavior at run-time and during tests
- set -b flat to byteMode default
- simplify config object flag count mode settings
- simplify config object flag count mode settings

### test

- make it easier to share common test setup and teardown code in test files

</details>

<details>
    <summary>
<h2> [2022-08-08] Release v0.1.1: fix bugs
</h2>
    </summary>

### fix

- add a wordMode config variable to simplify things
- get rid of the side-effect friendly global logVerbose var
- include new line characters in byte count

</details>

<details>
    <summary>
<h2> [2022-08-08] Release v0.1.0: version bump
</h2>
    </summary>

### feat

- add verboseMode
- add versionMode functionality to -V flag
- add word counter
- create application skeleton
- implement byteMode counter
- implement lineMode counter

### fix

- don't allow -b and -l at the same time

### test

- fix missing test coverage - at 100%

</details>

<details>
    <summary>
<h2> [2022-08-07] Release v0.0.0: initial release
</h2>
    </summary>

### build(go)

- add initial go.mod and go.sum files

### build(tools)

- add initial Makefile and scripts

### chore

- add inititial go directory skeleton
- add todo list
- add .editorconfig file
- set coverage to 0%

### ci

- add codeowners file

### ci(bash)

- add initial bash workflow

### ci(git)

- add initial git workflow

### ci(go)

- add initial Go workflow

### ci(links)

- add initial link check workflow

### ci(woke)

- add initial woke workflow

### doc

- update readme with project intro

</details>
