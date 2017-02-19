package util

type List []string
type Dict map[string]string
type Stats struct {
	MemoryAlloc       uint64 `json:"memory_alloc"`
	MemoryTotalAlloc  uint64 `json:"memory_total_alloc"`
	MemoryHeapAlloc   uint64 `json:"memory_heap_alloc"`
	MemoryHeapSys     uint64 `json:"memory_heap_sys"`
	MemoryHeapObjects uint64 `json:"memory_heap_objects"`
	MemoryMallocs     uint64 `json:"memory_mallocs"`
	MemoryFrees       uint64 `json:"memory_frees"`
	GCPauseTotalNs    uint64 `json:"gc_pause_total_ns"`
	NumGC             uint32 `json:"num_gc"`
}

type ResponseBasic struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type ResponseLen struct {
	ResponseBasic
	Length int `json:"length"`
}

type ResponseKeys struct {
	ResponseBasic
	Keys []string `json:"keys"`
}

type ResponseStats struct {
	ResponseBasic
	Stats
}

type ResponseString struct {
	ResponseBasic
	Value string `json:"value"`
}

type ResponseInt struct {
	ResponseBasic
	Value int `json:"value"`
}
