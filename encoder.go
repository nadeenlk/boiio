// Copyright © 2021 Nadeen Udantha udanthan@gmail.com

package boiio

import "io"

type Encoder struct {
	w *Writer
}

type Encodable interface {
	Encode(w *Writer)
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: NewWriter(w),
	}
}

func (e *Encoder) Encode(be Encodable) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	be.Encode(e.w)
	return err
}
