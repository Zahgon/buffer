package buffer

// List is a slice of Buffers, it's the backing for NewPartition
type List []Buffer

// Len is the sum of the Len()'s of the Buffers in the List.
func (l *List) Len() (n int64) { _ = "STUB: not implemented"; return 0 }

// Cap is the sum of the Cap()'s of the Buffers in the List.
func (l *List) Cap() (n int64) { _ = "STUB: not implemented"; return 0 }

// Reset calls Reset() on each of the Buffers in the list.
func (l *List) Reset() { _ = "STUB: not implemented"; return }

// Push adds a Buffer to the end of the List
func (l *List) Push(b Buffer) {
	_ = "STUB: not implemented"

	// Pop removes and returns a Buffer from the front of the List
	return
}

func (l *List) Pop() (b Buffer) { _ = "STUB: not implemented"; return *new(Buffer) }
