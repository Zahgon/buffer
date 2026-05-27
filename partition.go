package buffer

import (
	"encoding/gob"
)

type partition struct {
	List
	Pool
}

// NewPartition returns a Buffer which uses a Pool to extend or shrink its size as needed.
// It automatically allocates new buffers with pool.Get() to extend is length, and
// pool.Put() to release unused buffers as it shrinks.
func NewPartition(pool Pool, buffers ...Buffer) Buffer {
	_ = "STUB: not implemented"
	return *new(Buffer)
}

func (buf *partition) Cap() int64 { _ = "STUB: not implemented"; return 0 }

func (buf *partition) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *partition) grow() error { _ = "STUB: not implemented"; return nil }

func (buf *partition) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *partition) Reset() { _ = "STUB: not implemented"; return }

func init() {
	gob.Register(&partition{})
}
