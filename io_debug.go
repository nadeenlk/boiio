// Copyright © 2021 Nadeen Udantha udanthan@gmail.com

package boiio

import (
	"fmt"
	"io"
)

type io_debugRWC struct {
	n string
	r io.Reader
	w io.Writer
	c io.Closer
}

func DebugIoReader(name string, r io.Reader) io.Reader {
	return &io_debugRWC{
		n: name,
		r: r,
	}
}

func DebugIoWriter(name string, w io.Writer) io.Writer {
	return &io_debugRWC{
		n: name,
		w: w,
	}
}

func DebugIoReadWriter(name string, rw io.ReadWriter) io.ReadWriter {
	return &io_debugRWC{
		n: name,
		r: rw,
		w: rw,
	}
}

func DebugIoReadWriteCloser(name string, rwc io.ReadWriteCloser) io.ReadWriteCloser {
	return &io_debugRWC{
		n: name,
		r: rwc,
		w: rwc,
		c: rwc,
	}
}

func (drw *io_debugRWC) log(format string, a ...interface{}) {
	fmt.Printf("debugRWC %s %s\n", drw.n, fmt.Sprintf(format, a...))
}

func (drw *io_debugRWC) Read(p []byte) (int, error) {
	drw.log("Read0")
	n, err := drw.r.Read(p)
	drw.log("Read1 %d %v %v", n, err, p[:n])
	return n, err
}

func (drw *io_debugRWC) Write(p []byte) (int, error) {
	drw.log("Write0 %v", p)
	n, err := drw.w.Write(p)
	drw.log("Write1 %d %v %v", n, err, p[:n])
	return n, err
}

func (drw *io_debugRWC) Close() error {
	drw.log("Close0")
	err := drw.c.Close()
	drw.log("Close1 %v", err)
	return err
}
