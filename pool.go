package buffer

import (
	"encoding/gob"
	"sync"
)

// Pool provides a way to Allocate and Release Buffer objects
// Pools mut be concurrent-safe for calls to Get() and Put().
type Pool interface {
	Get() (Buffer, error) // Allocate a Buffer
	Put(buf Buffer) error // Release or Reuse a Buffer
}

type pool struct {
	pool sync.Pool
}

// NewPool returns a Pool(), it's backed by a sync.Pool so its safe for concurrent use.
// Get() and Put() errors will always be nil.
// It will not work with gob.
func NewPool(New func() Buffer) Pool { _ = "STUB: not implemented"; return *new(Pool) }

func (p *pool) Get() (Buffer, error) { _ = "STUB: not implemented"; return *new(Buffer), nil }

func (p *pool) Put(buf Buffer) error { _ = "STUB: not implemented"; return nil }

type memPool struct {
	N int64
	Pool
}

// NewMemPool returns a Pool, Get() returns an in memory buffer of max size N.
// Put() returns the buffer to the pool after resetting it.
// Get() and Put() errors will always be nil.
func NewMemPool(N int64) Pool { _ = "STUB: not implemented"; return *new(Pool) }

func (m *memPool) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *memPool) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

type filePool struct {
	N         int64
	Directory string
}

// NewFilePool returns a Pool, Get() returns a file-based buffer of max size N.
// Put() closes and deletes the underlying file for the buffer.
// Get() may return an error if it fails to create a file for the buffer.
// Put() may return an error if it fails to delete the file.
func NewFilePool(N int64, dir string) Pool { _ = "STUB: not implemented"; return *new(Pool) }

func (p *filePool) Get() (Buffer, error) { _ = "STUB: not implemented"; return *new(Buffer), nil }

func (p *filePool) Put(buf Buffer) (err error) { _ = "STUB: not implemented"; return nil }

func init() {
	gob.Register(&memPool{})
	gob.Register(&filePool{})
}
