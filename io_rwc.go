// Copyright © 2021 Nadeen Udantha udanthan@gmail.com

package boiio

import "io"

type io_ReadWriteCloser struct {
	io.Reader
	io.Writer
	io.Closer
}

func NewIoReadWriter(r io.Reader, w io.Writer) io.ReadWriter {
	return &io_ReadWriteCloser{
		Reader: r,
		Writer: w,
	}
}

func NewIoReadWriteCloser(r io.Reader, w io.Writer, c io.Closer) io.ReadWriteCloser {
	return &io_ReadWriteCloser{
		Reader: r,
		Writer: w,
		Closer: c,
	}
}
