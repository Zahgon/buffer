package buffer

import (
	"bytes"
	"encoding/gob"
	"io"
)

type memory struct {
	N int64
	*bytes.Buffer
}

// New returns a new in memory BufferAt with max size N.
// It's backed by a bytes.Buffer.
func New(n int64) BufferAt { _ = "STUB: not implemented"; return *new(BufferAt) }

func (buf *memory) Cap() int64 { _ = "STUB: not implemented"; return 0 }

func (buf *memory) Len() int64 { _ = "STUB: not implemented"; return 0 }

func (buf *memory) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *memory) WriteAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (buf *memory) ReadAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (buf *memory) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *memory) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func init() {
	gob.Register(&memory{})
}

func (buf *memory) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (buf *memory) UnmarshalBinary(bindata []byte) error { _ = "STUB: not implemented"; return nil }
