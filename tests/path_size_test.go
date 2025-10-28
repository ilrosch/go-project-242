package test

import (
	pathsize "code"

	"testing"

	"github.com/stretchr/testify/assert"
)

type GetSizeCase struct {
	name	 	string
	path		string
	expected 	int64
	recursive	bool
	all			bool
}

func TestGetSize(t *testing.T) {
	cases := []GetSizeCase{
		{name: "Default: directory", path: "./testdata/", expected: 662, recursive: false, all: false},
		{name: "Default: file", path: "./testdata/file.json", expected: 331, recursive: false, all: false},
		{name: "Enable recursive", path: "./testdata/", expected: 1554, recursive: true, all: false},
		{name: "Enable recursive + all", path: "./testdata/", expected: 1621, recursive: true, all: true},
		{name: "Empty directory", path: "./testdata/1/2/3", expected: 0, recursive: false, all: false},
	}

	for _, test := range cases {
		res, _ := pathsize.GetSize(test.path, test.recursive, test.all)
		assert.Equal(t, test.expected, res, test.name)
	}
}


type FormatSizeCase struct {
	name	 string
	size	 int64
	expected string
	human	 bool
}

func TestFormatSize(t *testing.T) {
	cases := []FormatSizeCase{
		{name: "Default", size: 1234567, expected: "1234567B", human: false},
		{name: "Enable human", size: 1234567, expected: "1.2MB", human: true},
		{name: "Enable human with small size", size: 66, expected: "66B", human: true},
	}

	for _, test := range cases {
		res := pathsize.FormatSize(test.size, test.human)
		assert.Equal(t, test.expected, res, test.name)
	}
}