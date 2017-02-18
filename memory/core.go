package memory

import (
	"sync"
)

// CACHE In-memory cache with synchronization
type CACHE struct {
	data map[string]interface{}
	*sync.RWMutex
	// @see http://stackoverflow.com/a/19168242/721525
	// @see https://medium.com/@deckarep/dancing-with-go-s-mutexes-92407ae927bf
}

func Alloc() CACHE {
	cache := CACHE{
		data:    map[string]interface{}{},
		RWMutex: new(sync.RWMutex),
	}

	return cache
}

func (cache *CACHE) Len() int {
	cache.RLock()
	defer cache.RUnlock()

	return len(cache.data)
}

func (cache *CACHE) Keys() []string {
	cache.RLock()
	defer cache.RUnlock()

	result := make([]string, 0, len(cache.data))
	for key := range cache.data {
		result = append(result, key)
	}

	return result
}
