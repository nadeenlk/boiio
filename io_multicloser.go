// Copyright © 2021 Nadeen Udantha udanthan@gmail.com

package boiio

import (
	"fmt"
	"io"
)

type io_MultiCloser struct {
	cs []io.Closer
}

func NewIoMultiCloser(cs ...io.Closer) io.Closer {
	return &io_MultiCloser{
		cs: cs,
	}
}

func (mc *io_MultiCloser) Close() error {
	errs := make([]error, 0)
	for _, c := range mc.cs {
		err := c.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return fmt.Errorf("%v", errs)
	}
}
