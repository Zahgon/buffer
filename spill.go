package buffer

import (
	"encoding/gob"
	"io"
)

type spill struct {
	Buffer
	Spiller io.Writer
}

// NewSpill returns a Buffer which writes data to w when there's an error
// writing to buf. Such as when buf is full, or the disk is full, etc.
func NewSpill(buf Buffer, w io.Writer) Buffer { _ = "STUB: not implemented"; return *new(Buffer) }

func (buf *spill) Cap() int64 { _ = "STUB: not implemented"; return 0 }

func (buf *spill) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func init() {
	gob.Register(&spill{})
}
