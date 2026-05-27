package buffer

import (
	"encoding/gob"
)

type swap struct {
	A BufferAt
	B BufferAt
}

// NewSwap creates a Buffer which writes to a until you write past a.Cap()
// then it io.Copy's from a to b and writes to b.
// Once the Buffer is empty again, it starts over writing to a.
// Note that if b.Cap() <= a.Cap() it will cause a panic, b is expected
// to be larger in order to accommodate writes past a.Cap().
func NewSwap(a, b Buffer) Buffer { _ = "STUB: not implemented"; return *new(Buffer) }

// NewSwapAt creates a BufferAt which writes to a until you write past a.Cap()
// then it io.Copy's from a to b and writes to b.
// Once the Buffer is empty again, it starts over writing to a.
// Note that if b.Cap() <= a.Cap() it will cause a panic, b is expected
// to be larger in order to accommodate writes past a.Cap().
func NewSwapAt(a, b BufferAt) BufferAt { _ = "STUB: not implemented"; return *new(BufferAt) }

func (buf *swap) Len() int64 { _ = "STUB: not implemented"; return 0 }

func (buf *swap) Cap() int64 { _ = "STUB: not implemented"; return 0 }

func (buf *swap) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *swap) ReadAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (buf *swap) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *swap) WriteAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (buf *swap) Reset() { _ = "STUB: not implemented"; return }

func init() {
	gob.Register(&swap{})
}
