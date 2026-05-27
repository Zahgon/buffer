package buffer

// ListAt is a slice of BufferAt's, it's the backing for NewPartitionAt
type ListAt []BufferAt

// Len is the sum of the Len()'s of the BufferAt's in the list.
func (l *ListAt) Len() (n int64) { _ = "STUB: not implemented"; return 0 }

// Cap is the sum of the Cap()'s of the BufferAt's in the list.
func (l *ListAt) Cap() (n int64) { _ = "STUB: not implemented"; return 0 }

// Reset calls Reset() on each of the BufferAt's in the list.
func (l *ListAt) Reset() { _ = "STUB: not implemented"; return }

// Push adds a BufferAt to the end of the list
func (l *ListAt) Push(b BufferAt) {
	_ = "STUB: not implemented"

	// Pop removes and returns a BufferAt from the front of the list
	return
}

func (l *ListAt) Pop() (b BufferAt) { _ = "STUB: not implemented"; return *new(BufferAt) }
