package buffer

import (
	"encoding/gob"
)

type chain struct {
	Buf  BufferAt
	Next BufferAt
}

type nopBufferAt struct {
	Buffer
}

func (buf *nopBufferAt) ReadAt(p []byte, off int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (buf *nopBufferAt) WriteAt(p []byte, off int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// toBufferAt converts a Buffer to a BufferAt with nop ReadAt and WriteAt funcs
func toBufferAt(buf Buffer) BufferAt { _ = "STUB: not implemented"; return *new(BufferAt) }

// NewMultiAt returns a BufferAt which is the logical concatenation of the passed BufferAts.
// The data in the buffers is shifted such that there is no non-empty buffer following
// a non-full buffer, this process is also run after every Read.
// If no buffers are passed, the returned Buffer is nil.
func NewMultiAt(buffers ...BufferAt) BufferAt { _ = "STUB: not implemented"; return *new(BufferAt) }

// NewMulti returns a Buffer which is the logical concatenation of the passed buffers.
// The data in the buffers is shifted such that there is no non-empty buffer following
// a non-full buffer, this process is also run after every Read.
// If no buffers are passed, the returned Buffer is nil.
func NewMulti(buffers ...Buffer) Buffer { _ = "STUB: not implemented"; return *new(Buffer) }

func (buf *chain) Reset() { _ = "STUB: not implemented"; return }

func (buf *chain) Cap() (n int64) { _ = "STUB: not implemented"; return 0 }

func (buf *chain) Len() (n int64) { _ = "STUB: not implemented"; return 0 }

func (buf *chain) Defrag() { _ = "STUB: not implemented"; return }

func (buf *chain) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *chain) ReadAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (buf *chain) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *chain) WriteAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// past the end

// fits in

// partial fit

func init() {
	gob.Register(&chain{})
	gob.Register(&nopBufferAt{})
}

func (buf *chain) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (buf *chain) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }
