package appwc

type dataTestCase struct {
	name       string
	input      string
	wantLength int
	wantByte   int
	wantChar   int
	wantLine   int
	wantWord   int
}

type dataTestCases []dataTestCase

var testData = dataTestCases{
	{
		name:       "1 unicode char",
		input:      "😂",
		wantByte:   4,
		wantChar:   1,
		wantLength: 1,
		wantLine:   0,
		wantWord:   1,
	},
	{
		name:       "7 unicode chars",
		input:      "😂😂😂😂😂😂😂",
		wantByte:   28,
		wantChar:   7,
		wantLength: 7,
		wantLine:   0,
		wantWord:   1,
	},
	{
		name:       "7 unicode chars with 3 newlines",
		input:      "😂😂\n😂\n😂😂😂\n😂",
		wantByte:   31,
		wantChar:   10,
		wantLength: 7,
		wantLine:   3,
		wantWord:   4,
	},
	{
		name:       "7 unicode chars with 5 newlines",
		input:      "\n😂😂\n😂\n😂😂😂\n😂\n",
		wantByte:   33,
		wantChar:   12,
		wantLength: 7,
		wantLine:   5,
		wantWord:   4,
	},
	{
		name:       "7 spaces",
		input:      "		",
		wantByte:   7,
		wantChar:   7,
		wantLength: 7,
		wantLine:   0,
		wantWord:   0,
	},
	{
		name:       "3 spaces, nl, 3 spaces",
		input:      "	\n	 ",
		wantByte:   7,
		wantChar:   7,
		wantLength: 7,
		wantLine:   1,
		wantWord:   0,
	},
	{
		name:       "3 spaces, nl",
		input:      "	\n",
		wantByte:   4,
		wantChar:   4,
		wantLength: 4,
		wantLine:   1,
		wantWord:   0,
	},
	{
		name:       "nl, 3 spaces",
		input:      "\n   ",
		wantByte:   4,
		wantChar:   4,
		wantLength: 4,
		wantLine:   1,
		wantWord:   0,
	},
	{
		name:       "1 word, no nl",
		input:      "one",
		wantByte:   3,
		wantChar:   3,
		wantLength: 3,
		wantLine:   0,
		wantWord:   1,
	},
	{
		name:       "1 word and nl",
		input:      "one\n",
		wantByte:   4,
		wantChar:   4,
		wantLength: 4,
		wantLine:   1,
		wantWord:   1,
	},
	{
		name:       "2 words, no nl",
		input:      "one two",
		wantByte:   7,
		wantChar:   7,
		wantLength: 7,
		wantLine:   0,
		wantWord:   2,
	},
	{
		name:       "word, nl, word",
		input:      "one\ntwo",
		wantByte:   7,
		wantChar:   7,
		wantLength: 7,
		wantLine:   1,
		wantWord:   2,
	},
	{
		name:       "2 words and nl",
		input:      "one two\n",
		wantByte:   8,
		wantChar:   8,
		wantLength: 8,
		wantLine:   1,
		wantWord:   2,
	},
	{
		name:       "nl and 2 words",
		input:      "\none two",
		wantByte:   8,
		wantChar:   8,
		wantLength: 8,
		wantLine:   1,
		wantWord:   2,
	},
	{
		name:       "word, nl, word, nl, word",
		input:      "one\ntwo\nthree",
		wantByte:   13,
		wantChar:   13,
		wantLength: 13,
		wantLine:   2,
		wantWord:   3,
	},
	{
		name:       "word, nl, word, nl, 2 words, nl",
		input:      "one\ntwo\nthree four\n",
		wantByte:   19,
		wantChar:   19,
		wantLength: 19,
		wantLine:   3,
		wantWord:   4,
	},
	{
		name:       "word, nl, word, 2 nl, 2 words, nl",
		input:      "one\ntwo\n\nthree four\n",
		wantByte:   20,
		wantChar:   20,
		wantLength: 20,
		wantLine:   4,
		wantWord:   4,
	},
	{
		name:       "3 spaces, word, 3 spaces, word, 3 spaces, word, 3 spaces",
		input:      "	one   two	three	",
		wantByte:   23,
		wantChar:   23,
		wantLength: 23,
		wantLine:   0,
		wantWord:   3,
	},
	{
		name:       "3 words with spaces between them with newlines",
		input:      " \n one \n two \n three \n ",
		wantByte:   23,
		wantChar:   23,
		wantLength: 23,
		wantLine:   4,
		wantWord:   3,
	},
}

type flagTestCase struct {
	name  string
	flags []string
	want  string
}

type flagTestCases []flagTestCase

var testFlags = flagTestCases{
	{
		name:  "invalid flag",
		flags: []string{"-x"},
		want:  "flag provided but not defined: -x",
	},
	{
		name:  "show usage",
		flags: []string{"-h"},
		want:  "Usage: cli [OPTION]...",
	},
}
