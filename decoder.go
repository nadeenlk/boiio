// Copyright © 2021 Nadeen Udantha udanthan@gmail.com

package boiio

import "io"

type Decoder struct {
	r *Reader
}

type Decodable interface {
	Decode(r *Reader)
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: NewReader(r),
	}
}

func (d *Decoder) Decode(bd Decodable) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	bd.Decode(d.r)
	return err
}
