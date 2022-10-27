# "Powerful Cli Applications in Go" Workspace Release Notes

<details open>
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
- add collapsable section in ch01 for the examples
- add collapsable sections for each chapter
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

- add codefactor badge to readme
- add Go version badge to readme
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

- add code of conduct
- add codeql badge and create a new badge row for the main checks
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
- add -r rune mode
- add usage help message
- let -l, -w, -c, -b be used like they are in the coreutils wc cli

### fix

- add more test cases, fix bugs with getCounts()
- add -w argument and let -b and -l override -w
- change -b to -c to match wc from coreutils
- change non-verbose output to match wc from coreutils
- change -r to -m to match wc from coreutils
- change "rune" to "char" to match coreutils wc
- change the help string for -l and -w to match wc from coreutils
- change the help string for -V to match wc from coreutils
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
- add versionMode functionallity to -V flag
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

- add .editorconfig file
- add inititial go directory skeleton
- add todo list
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
