package utils

import (
	"io"
)

var GCloser *closer = newСloser()

type closer struct {
	closer []io.Closer
}

func newСloser() *closer {
	return new(closer)
}

func (c *closer) Append(cl io.Closer) {
	c.closer = append(c.closer, cl)
}

func (c *closer) Close() {
	for i := len(c.closer) - 1; i >= 0; i-- {
		c.closer[i].Close()
	}
}
