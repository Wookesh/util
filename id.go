package util

import "sync"

// ID is thread safe per run unique ID generator
// can serve up to 2^31 unique IDs
type ID struct {
	last int32
	m    sync.Mutex
}

// Next gives next available ID
func (id *ID) Next() int32 {
	id.m.Lock()
	defer id.m.Unlock()
	id.last++
	return id.last
}
