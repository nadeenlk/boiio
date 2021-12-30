// Copyright © 2021 Nadeen Udantha udanthan@gmail.com

package boiio

import (
	"encoding/binary"
	"io"
	"sync"
)

type Writer struct {
	w   io.Writer
	l   sync.Mutex
	tmp []byte
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w:   w,
		tmp: make([]byte, 8),
	}
}

func (w *Writer) Lock() {
	w.l.Lock()
}

func (w *Writer) Unlock() {
	w.l.Unlock()
}

func (w *Writer) wb(x []byte) {
	_, err := w.w.Write(x)
	if err != nil {
		throw(err)
	}
}

func (w *Writer) wn(i int) {
	w.wb(w.tmp[:i])
}

func (w *Writer) U8(x uint8) {
	w.tmp[0] = x
	w.wn(1)
}

func (w *Writer) U16(x uint16) {
	binary.BigEndian.PutUint16(w.tmp, x)
	w.wn(2)
}

func (w *Writer) U32(x uint32) {
	binary.BigEndian.PutUint32(w.tmp, x)
	w.wn(4)
}

func (w *Writer) U64(x uint64) {
	binary.BigEndian.PutUint64(w.tmp, x)
	w.wn(8)
}

func (w *Writer) S8(x int8) {
	w.tmp[0] = byte(x)
	w.wn(1)
}

func (w *Writer) S16(x int16) {
	binary.BigEndian.PutUint16(w.tmp, uint16(x))
	w.wn(2)
}

func (w *Writer) S32(x int32) {
	binary.BigEndian.PutUint32(w.tmp, uint32(x))
	w.wn(4)
}

func (w *Writer) S64(x int64) {
	binary.BigEndian.PutUint64(w.tmp, uint64(x))
	w.wn(8)
}

func (w *Writer) BA(x []byte) {
	w.wb(x)
}
