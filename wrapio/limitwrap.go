package wrapio

import (
	"encoding/gob"
	"io"
)

// ReadWriterAt implements io.ReaderAt and io.WriterAt
type ReadWriterAt interface {
	io.ReaderAt
	io.WriterAt
}

// Wrapper implements a io.ReadWriter and ReadWriterAt such that
// when reading/writing goes past N bytes, it "wraps" back to the beginning.
type Wrapper struct {
	// N is the offset at which to "wrap" back to the start
	N int64
	// L is the length of the data written
	L int64
	// O is our offset in the data
	O   int64
	rwa ReadWriterAt
}

// NewWrapper creates a Wrapper based on ReadWriterAt rwa.
// L is the current length, O is the current offset, and N is offset at which we "wrap".
func NewWrapper(rwa ReadWriterAt, L, O, N int64) *Wrapper { _ = "STUB: not implemented"; return nil }

// Len returns the # of bytes in the Wrapper
func (wpr *Wrapper) Len() int64 {
	_ = "STUB: not implemented"

	// Cap returns the "wrap" offset (max # of bytes)
	return 0
}

func (wpr *Wrapper) Cap() int64 {
	_ = "STUB: not implemented"

	// Reset seeks to the start (0 offset), and sets the length to 0.
	return 0
}

func (wpr *Wrapper) Reset() { _ = "STUB: not implemented"; return }

// SetReadWriterAt lets you switch the underlying Read/WriterAt
func (wpr *Wrapper) SetReadWriterAt(rwa ReadWriterAt) {
	_ = "STUB: not implemented"

	// Read reads from the current offset into p, wrapping at Cap()
	return
}

func (wpr *Wrapper) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// ReadAt reads from the current offset+off into p, wrapping at Cap()
func (wpr *Wrapper) ReadAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Write writes p to the end of the Wrapper (at Len()), wrapping at Cap()
func (wpr *Wrapper) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteAt writes p at the current offset+off, wrapping at Cap()
func (wpr *Wrapper) WriteAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func init() {
	gob.Register(&Wrapper{})
}
