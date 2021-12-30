// Copyright © 2021 Nadeen Udantha udanthan@gmail.com

package boiio

type BufferWriter struct {
	*Writer
	Buf []byte
	Off int
}

func NewBufferWriter(buf []byte) *BufferWriter {
	bw := &BufferWriter{
		Buf: buf,
	}
	bw.Writer = NewWriter(bw)
	return bw
}

func (bw *BufferWriter) Write(p []byte) (int, error) {
	n := copy(bw.Buf[bw.Off:], p)
	bw.Off += n
	return n, nil
}
