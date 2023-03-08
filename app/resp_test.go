package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	strs := []struct {
		input string
		want  string
	}{
		{"OK", "$2\r\nOK\r\n"},
		{"PONG", "$4\r\nPONG\r\n"},
		{"", "$0\r\n\r\n"},
	}

	for _, tt := range strs {
		testName := fmt.Sprintf("Testing input %s", tt.input)
		t.Run(testName, func(t *testing.T) {
			var buffer bytes.Buffer
			rw := NewRespWriter(&buffer)
			err := rw.Write(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(buffer.Bytes(), []byte(tt.want)) {
				t.Errorf(`Wrong output want %v got %v`, []byte(tt.want), buffer.Bytes())
			}
		})
	}
}

func TestWriteError(t *testing.T) {
	errs := []struct {
		input string
		want  string
	}{
		{"ERR", "-ERR\r\n"},
		{"ERR unknown command 'foobar'", "-ERR unknown command 'foobar'\r\n"},
	}

	for _, tt := range errs {
		testName := fmt.Sprintf("Testing input : %s", tt.input)

		t.Run(testName, func(t *testing.T) {
			var buffer bytes.Buffer
			rw := NewRespWriter(&buffer)
			err := rw.WriteError(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(buffer.Bytes(), []byte(tt.want)) {
				t.Errorf(`Wrong output want %s got %s`, []byte(tt.want), buffer.Bytes())
			}
		})
	}
}
