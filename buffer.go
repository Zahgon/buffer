// Package buffer implements a series of Buffers which can be composed to implement complicated buffering strategies
package buffer

import (
	"io"
)

// Buffer is used to Write() data which will be Read() later.
type Buffer interface {
	Len() int64 // How much data is Buffered in bytes
	Cap() int64 // How much data can be Buffered at once in bytes.
	io.Reader   // Read() will read from the top of the buffer [io.EOF if empty]
	io.Writer   // Write() will write to the end of the buffer [io.ErrShortWrite if not enough space]
	Reset()     // Truncates the buffer, Len() == 0.
}

// BufferAt is a buffer which supports io.ReaderAt and io.WriterAt
type BufferAt interface {
	Buffer
	io.ReaderAt
	io.WriterAt
}

func len64(p []byte) int64 { _ = "STUB: not implemented"; return 0 }

// Gap returns buf.Cap() - buf.Len()
func Gap(buf Buffer) int64 { _ = "STUB: not implemented"; return 0 }

// Full returns true iff buf.Len() == buf.Cap()
func Full(buf Buffer) bool { _ = "STUB: not implemented"; return false }

// Empty returns false iff buf.Len() == 0
func Empty(buf Buffer) bool { _ = "STUB: not implemented"; return false }

// NewUnboundedBuffer returns a Buffer which buffers "mem" bytes to memory
// and then creates file's of size "file" to buffer above "mem" bytes.
func NewUnboundedBuffer(mem, file int64) Buffer { _ = "STUB: not implemented"; return *new(Buffer) }
