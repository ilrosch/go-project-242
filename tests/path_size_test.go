package test

import (
	pathsize "code"

	"testing"

	"github.com/stretchr/testify/assert"
)

type CaseGetSize struct {
	name	 	string
	expected 	int64
	path	 	string
	hiddenFiles bool
}

func TestGetSize(t *testing.T) {
	cases := []CaseGetSize{
		{name: "Default dir", expected: 1332, path: "./testdata/test_dir_1", hiddenFiles: false},
		{name: "Default file", expected: 331, path: "./testdata/file.json",  hiddenFiles: false},
		{name: "Include hidden files", expected: 2664, path: "./testdata/test_dir_1",  hiddenFiles: true},
		{name: "Hidden file", expected: 1332, path: "./testdata/test_dir_1/.file.txt",  hiddenFiles: true},
		{name: "Hidden file", expected: 0, path: "./testdata/test_dir_1/.file.txt",  hiddenFiles: false},
		{name: "Not found file", expected: 0, path: "./testdata/test_dir_1/unknown",  hiddenFiles: false},
	}

	for _, test := range cases {
		res, _ := pathsize.GetSize(test.path, test.hiddenFiles)
		assert.Equal(t, test.expected, res, test.name)
	}
}


type CaseFormatSize struct {
	name	 string
	expected string
	size	 int64
	flag	 bool
}

func TestFormatSize(t *testing.T) {
	cases := []CaseFormatSize{
		{name: "Default", expected: "1234567B", size: 1234567, flag: false},
		{name: "Human readable normal", expected: "1.2MB", size: 1234567, flag: true},
		{name: "Human readable small", expected: "12B", size: 12, flag: true},
	}

	for _, test := range cases {
		res := pathsize.FormatSize(test.size, test.flag)
		assert.Equal(t, test.expected, res, test.name)
	}
}