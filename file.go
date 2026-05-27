package buffer

import (
	"encoding/gob"
	"io"
	"os"

	"github.com/djherbis/buffer/wrapio"
)

// File is used as the backing resource for a the NewFile BufferAt.
type File interface {
	Name() string
	Stat() (fi os.FileInfo, err error)
	io.ReaderAt
	io.WriterAt
	Close() error
}

type fileBuffer struct {
	file File
	*wrapio.Wrapper
}

// NewFile returns a new BufferAt backed by "file" with max-size N.
func NewFile(N int64, file File) BufferAt { _ = "STUB: not implemented"; return *new(BufferAt) }

func init() {
	gob.Register(&fileBuffer{})
}

func (buf *fileBuffer) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (buf *fileBuffer) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }
