package main

import (
	"bufio"
	"io"
	"strconv"
)

var (
	bulkStringByte = []byte{'$'}

	errorByte = []byte{'-'}

	linEndByte = []byte{'\r', '\n'}
)

type RespWriter struct{}

func NewRespWriter() *RespWriter {
	return &RespWriter{}
}

func (rw *RespWriter) Write(w io.Writer, strs ...string) error {
	bw := bufio.NewWriter(w)

	for _, str := range strs {
		bw.Write(bulkStringByte)
		bw.WriteString(strconv.Itoa(len(str)))
		bw.Write(linEndByte)
		bw.WriteString(str)
		bw.Write(linEndByte)
	}

	return bw.Flush()
}

func (rw *RespWriter) WriteError(w io.Writer, errStr string) error {

	bw := bufio.NewWriter(w)

	bw.Write(errorByte)

	bw.WriteString(errStr)

	bw.Write(linEndByte)

	return bw.Flush()
}
