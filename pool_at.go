package buffer

import (
	"encoding/gob"
	"sync"
)

// PoolAt provides a way to Allocate and Release BufferAt objects
// PoolAt's mut be concurrent-safe for calls to Get() and Put().
type PoolAt interface {
	Get() (BufferAt, error) // Allocate a BufferAt
	Put(buf BufferAt) error // Release or Reuse a BufferAt
}

type poolAt struct {
	poolAt sync.Pool
}

// NewPoolAt returns a PoolAt(), it's backed by a sync.Pool so its safe for concurrent use.
// Get() and Put() errors will always be nil.
// It will not work with gob.
func NewPoolAt(New func() BufferAt) PoolAt { _ = "STUB: not implemented"; return *new(PoolAt) }

func (p *poolAt) Get() (BufferAt, error) { _ = "STUB: not implemented"; return *new(BufferAt), nil }

func (p *poolAt) Put(buf BufferAt) error { _ = "STUB: not implemented"; return nil }

type memPoolAt struct {
	N int64
	PoolAt
}

// NewMemPoolAt returns a PoolAt, Get() returns an in memory buffer of max size N.
// Put() returns the buffer to the pool after resetting it.
// Get() and Put() errors will always be nil.
func NewMemPoolAt(N int64) PoolAt { _ = "STUB: not implemented"; return *new(PoolAt) }

func (m *memPoolAt) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *memPoolAt) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

type filePoolAt struct {
	N         int64
	Directory string
}

// NewFilePoolAt returns a PoolAt, Get() returns a file-based buffer of max size N.
// Put() closes and deletes the underlying file for the buffer.
// Get() may return an error if it fails to create a file for the buffer.
// Put() may return an error if it fails to delete the file.
func NewFilePoolAt(N int64, dir string) PoolAt { _ = "STUB: not implemented"; return *new(PoolAt) }

func (p *filePoolAt) Get() (BufferAt, error) { _ = "STUB: not implemented"; return *new(BufferAt), nil }

func (p *filePoolAt) Put(buf BufferAt) (err error) { _ = "STUB: not implemented"; return nil }

func init() {
	gob.Register(&memPoolAt{})
	gob.Register(&filePoolAt{})
}
