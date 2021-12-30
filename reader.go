// Copyright © 2021 Nadeen Udantha udanthan@gmail.com

package boiio

import (
	"encoding/binary"
	"io"
)

type Reader struct {
	r   io.Reader
	tmp []byte
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r:   r,
		tmp: make([]byte, 8),
	}
}

func (r *Reader) rb(x []byte) []byte {
	_, err := io.ReadFull(r.r, x)
	if err != nil {
		throw(err)
	}
	return x
}

func (r *Reader) rn(i int) []byte {
	return r.rb(r.tmp[:i])
}

func (r *Reader) U8() uint8   { return r.rn(1)[0] }
func (r *Reader) U16() uint16 { return binary.BigEndian.Uint16(r.rn(2)) }
func (r *Reader) U32() uint32 { return binary.BigEndian.Uint32(r.rn(4)) }
func (r *Reader) U64() uint64 { return binary.BigEndian.Uint64(r.rn(8)) }

func (r *Reader) S8() int8   { return int8(r.rn(1)[0]) }
func (r *Reader) S16() int16 { return int16(binary.BigEndian.Uint16(r.rn(2))) }
func (r *Reader) S32() int32 { return int32(binary.BigEndian.Uint32(r.rn(4))) }
func (r *Reader) S64() int64 { return int64(binary.BigEndian.Uint64(r.rn(8))) }

func (r *Reader) BA(x []byte) []byte {
	return r.rb(x)
}
