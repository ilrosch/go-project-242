package test

import (
	pathsize "code"

	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name	 string
	expected int64
	path	 string
}

func TestGetSize(t *testing.T) {
	cases := []Case{
		{name: "Default dir", expected: 1332, path: "./testdata/test_dir_1"},
		{name: "Default file", expected: 331, path: "./testdata/file.json"},
		{name: "Not found file", expected: 0, path: "./testdata/test_dir_1/unknown"},
	}

	for _, test := range cases {
		res, _ := pathsize.GetSize(test.path)
		assert.Equal(t, test.expected, res, test.name)
	}
}