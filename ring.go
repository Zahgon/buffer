package buffer

import (
	"github.com/djherbis/buffer/wrapio"
)

type ring struct {
	BufferAt
	L int64
	*wrapio.WrapReader
	*wrapio.WrapWriter
}

// NewRing returns a Ring Buffer from a BufferAt.
// It overwrites old data in the Buffer when needed (when its full).
func NewRing(buffer BufferAt) Buffer { _ = "STUB: not implemented"; return *new(Buffer) }

func (buf *ring) Len() int64 { _ = "STUB: not implemented"; return 0 }

func (buf *ring) Cap() int64 { _ = "STUB: not implemented"; return 0 }

func (buf *ring) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *ring) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf *ring) Reset() { _ = "STUB: not implemented"; return }
