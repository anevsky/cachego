package memory

import (
	"runtime"
	"sync"
	"time"

	"github.com/anevsky/cachego/util"
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

func (cache *CACHE) Stats() util.Stats {
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)

	stats := util.Stats{
		MemoryAlloc:       memStats.Alloc,
		MemoryTotalAlloc:  memStats.TotalAlloc,
		MemoryHeapAlloc:   memStats.HeapAlloc,
		MemoryHeapSys:     memStats.HeapSys,
		MemoryHeapObjects: memStats.HeapObjects,
		MemoryMallocs:     memStats.Mallocs,
		MemoryFrees:       memStats.Frees,
		GCPauseTotalNs:    memStats.PauseTotalNs,
		NumGC:             memStats.NumGC,
	}

	return stats
}

func (cache *CACHE) SetTTL(key string, ttl int) error {
	if ttl < 0 {
		return util.ErrorInvalidTTLValue
	}

	if ttl == 0 {
		return nil
	}

	time.AfterFunc(time.Millisecond*time.Duration(ttl), func() {
		cache.Lock()
		delete(cache.data, key)
		cache.Unlock()
	})

	return nil
}
