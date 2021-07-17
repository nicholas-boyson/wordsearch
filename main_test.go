package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		test  string
		bytes []byte
	}{
		{
			test:  "Quit",
			bytes: []byte("quit"),
		},
		{
			test:  "ListOptionsThenQuit",
			bytes: []byte("2\nquit\n"),
		},
	}

	for _, tt := range tests {
		var stdin bytes.Buffer
		stdin.Write(tt.bytes)

		var scanner = bufio.NewScanner(&stdin)

		err := process(scanner)
		assert.Nil(t, err)
	}
}
