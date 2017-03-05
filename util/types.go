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

type BasicDTO struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type LenDTO struct {
	BasicDTO
	Length int `json:"length"`
}

type KeysDTO struct {
	BasicDTO
	Keys []string `json:"keys"`
}

type StatsDTO struct {
	BasicDTO
	Stats
}

type StringDTO struct {
	BasicDTO
	Value string `json:"value"`
}

type IntDTO struct {
	BasicDTO
	Value int `json:"value"`
}

type ListDTO struct {
	BasicDTO
	Value List `json:"value"`
}

type DictDTO struct {
	BasicDTO
	Value Dict `json:"value"`
}

type BoolDTO struct {
	BasicDTO
	Value bool `json:"value"`
}
