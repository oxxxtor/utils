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
	for i := 0; i < len(c.closer); i++ {
		c.closer[i].Close()
	}
}
