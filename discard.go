package buffer

import (
	"encoding/gob"
)

type discard struct{}

// Discard is a Buffer which writes to ioutil.Discard and read's return 0, io.EOF.
// All of its methods are concurrent safe.
var Discard Buffer = discard{}

func (buf discard) Len() int64 { _ = "STUB: not implemented"; return 0 }

func (buf discard) Cap() int64 { _ = "STUB: not implemented"; return 0 }

func (buf discard) Reset() { _ = "STUB: not implemented"; return }

func (buf discard) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (buf discard) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func init() {
	gob.Register(&discard{})
}
