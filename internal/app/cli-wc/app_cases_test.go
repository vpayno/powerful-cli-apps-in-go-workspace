package appwc

type testCase struct {
	name     string
	input    string
	wantByte int
	wantChar int
	wantWord int
	wantLine int
}

type testCases []testCase

var testData = testCases{
	{
		name:     "7 spaces",
		input:    "       ",
		wantByte: 7,
		wantChar: 7,
		wantWord: 0,
		wantLine: 0,
	},
	{
		name:     "3 spaces, nl, 3 spaces",
		input:    "   \n   ",
		wantByte: 7,
		wantChar: 7,
		wantWord: 0,
		wantLine: 1,
	},
	{
		name:     "3 spaces, nl",
		input:    "   \n",
		wantByte: 4,
		wantChar: 4,
		wantWord: 0,
		wantLine: 1,
	},
	{
		name:     "nl, 3 spaces",
		input:    "\n   ",
		wantByte: 4,
		wantChar: 4,
		wantWord: 0,
		wantLine: 1,
	},
	{
		name:     "1 word, no nl",
		input:    "one",
		wantByte: 3,
		wantChar: 3,
		wantWord: 1,
		wantLine: 0,
	},
	{
		name:     "1 word and nl",
		input:    "one\n",
		wantByte: 4,
		wantChar: 4,
		wantWord: 1,
		wantLine: 1,
	},
	{
		name:     "2 words, no nl",
		input:    "one two",
		wantByte: 7,
		wantChar: 7,
		wantWord: 2,
		wantLine: 0,
	},
	{
		name:     "word, nl, word",
		input:    "one\ntwo",
		wantByte: 7,
		wantChar: 7,
		wantWord: 2,
		wantLine: 1,
	},
	{
		name:     "2 words and nl",
		input:    "one two\n",
		wantByte: 8,
		wantChar: 8,
		wantWord: 2,
		wantLine: 1,
	},
	{
		name:     "nl and 2 words",
		input:    "\none two",
		wantByte: 8,
		wantChar: 8,
		wantWord: 2,
		wantLine: 1,
	},
	{
		name:     "word, nl, word, nl, word",
		input:    "one\ntwo\nthree",
		wantByte: 13,
		wantChar: 13,
		wantWord: 3,
		wantLine: 2,
	},
	{
		name:     "word, nl, word, nl, 2 words, nl",
		input:    "one\ntwo\nthree four\n",
		wantByte: 19,
		wantChar: 19,
		wantWord: 4,
		wantLine: 3,
	},
	{
		name:     "word, nl, word, 2 nl, 2 words, nl",
		input:    "one\ntwo\n\nthree four\n",
		wantByte: 20,
		wantChar: 20,
		wantWord: 4,
		wantLine: 4,
	},
	{
		name:     "3 spaces, word, 3 spaces, word, 3 spaces, word, 3 spaces",
		input:    "   one   two   three   ",
		wantByte: 23,
		wantChar: 23,
		wantWord: 3,
		wantLine: 0,
	},
	{
		name:     "3 words with spaces between them with newlines",
		input:    " \n one \n two \n three \n ",
		wantByte: 23,
		wantChar: 23,
		wantWord: 3,
		wantLine: 4,
	},
}
