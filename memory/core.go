package memory

import (
)

type CACHE struct {
  data map[string]interface{}
}

func Alloc() CACHE {
  cache := CACHE{
    data: map[string]interface{}{},
  }

  return cache
}

func (cache *CACHE) Len() int {
  return len(cache.data)
}

func (cache *CACHE) Keys() []string {
  result := make([]string, 0, len(cache.data))
  for key := range cache.data {
    result = append(result, key)
  }

  return result
}
