package wrapio

import "io"

// DoerAt is a common interface for wrappers WriteAt or ReadAt functions
type DoerAt interface {
	DoAt([]byte, int64) (int, error)
}

// DoAtFunc is implemented by ReadAt/WriteAt
type DoAtFunc func([]byte, int64) (int, error)

type wrapper struct {
	off    int64
	wrapAt int64
	doat   DoAtFunc
}

func (w *wrapper) Offset() int64 { _ = "STUB: not implemented"; return 0 }

func (w *wrapper) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *wrapper) DoAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0,

		// WrapWriter wraps writes around a section of data.
		nil
}

type WrapWriter struct {
	*wrapper
}

// NewWrapWriter creates a WrapWriter starting at offset off, and wrapping at offset wrapAt.
func NewWrapWriter(w io.WriterAt, off int64, wrapAt int64) *WrapWriter {
	_ = "STUB: not implemented"
	return nil
}

// Write writes p starting at the current offset, wrapping when it reaches the end.
// The current offset is shifted forward by the amount written.
func (w *WrapWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteAt writes p starting at offset off, wrapping when it reaches the end.
func (w *WrapWriter) WriteAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WrapReader wraps reads around a section of data.
type WrapReader struct {
	*wrapper
}

// NewWrapReader creates a WrapReader starting at offset off, and wrapping at offset wrapAt.
func NewWrapReader(r io.ReaderAt, off int64, wrapAt int64) *WrapReader {
	_ = "STUB: not implemented"
	return nil
}

// Read reads into p starting at the current offset, wrapping if it reaches the end.
// The current offset is shifted forward by the amount read.
func (r *WrapReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// ReadAt reads into p starting at the current offset, wrapping when it reaches the end.
func (r *WrapReader) ReadAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// maxConsecutiveEmptyActions determines how many consecutive empty reads/writes can occur before giving up
const maxConsecutiveEmptyActions = 100

// Wrap causes an action on an array of bytes (like read/write) to be done from an offset off,
// wrapping at offset wrapAt.
func Wrap(w DoerAt, p []byte, off int64, wrapAt int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
