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

type RespWriter struct {
	*bufio.Writer
}

func NewRespWriter(w io.Writer) *RespWriter {
	return &RespWriter{
		Writer: bufio.NewWriter(w),
	}
}

func (rw *RespWriter) Write(strs ...string) error {

	for _, str := range strs {
		rw.Writer.Write(bulkStringByte)
		rw.Writer.WriteString(strconv.Itoa(len(str)))
		rw.Writer.Write(linEndByte)
		rw.Writer.WriteString(str)
		rw.Writer.Write(linEndByte)
	}

	return rw.Writer.Flush()
}

func (rw *RespWriter) WriteError(errStr string) error {

	rw.Writer.Write(errorByte)

	rw.Writer.Write([]byte(errStr))

	rw.Writer.Write(linEndByte)

	return rw.Writer.Flush()
}
