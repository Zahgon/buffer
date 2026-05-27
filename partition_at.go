package buffer

import (
	"encoding/gob"
)

type partitionAt struct {
	ListAt
	PoolAt
}

// NewPartitionAt returns a BufferAt which uses a PoolAt to extend or shrink its size as needed.
// It automatically allocates new buffers with pool.Get() to extend is length, and
// pool.Put() to release unused buffers as it shrinks.
func NewPartitionAt(pool PoolAt, buffers ...BufferAt) BufferAt {
	_ = "STUB: not implemented"
	return *new(BufferAt)
}

func (buf *partitionAt) Cap() int64 { _ = "STUB: not implemented"; return 0 }

func (buf *partitionAt) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *partitionAt) ReadAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Find the buffer where this offset is found.

// We need to read more, starting from 0 in the next buffer.

func (buf *partitionAt) grow() error { _ = "STUB: not implemented"; return nil }

func (buf *partitionAt) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (buf *partitionAt) WriteAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// writing at the end special case

// Find the buffer where this offset is found.

// Everything should fit.

// Assume it won't all fit, only write what should fit.

// All writes are at offset 0 of following buffers now.

func (buf *partitionAt) Reset() { _ = "STUB: not implemented"; return }

func init() {
	gob.Register(&partitionAt{})
}
