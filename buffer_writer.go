// Copyright © 2021 Nadeen Udantha udanthan@gmail.com

package boiio

import (
	"bufio"
	"io"
)

type BufferedWriter struct {
	*Writer
	bw *bufio.Writer
}

func (bw *BufferedWriter) Flush() {
	err := bw.bw.Flush()
	if err != nil {
		throw(err)
	}
}

func NewBufferedWriter(w io.Writer, sz int) *BufferedWriter {
	bw := bufio.NewWriterSize(w, sz)
	return &BufferedWriter{
		Writer: NewWriter(bw),
		bw:     bw,
	}
}
