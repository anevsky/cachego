package memory

import (
  "github.com/anevsky/cachego/util"
)

func (cache *CACHE) Get(key string) (interface{}, error) {
  value, ok := cache.data[key]
  if !ok {
    return "", util.ErrorBadRequest
  }

  switch v := value.(type) {
  case int:
    return v, nil
  case string:
    return v, nil
  case util.List:
    return v, nil
  case util.Dict:
    return v, nil
  default:
    return "", util.ErrorBadRequest
  }
}

func (cache *CACHE) HasKey(key string) (bool, error) {
  if _, ok := cache.data[key]; !ok {
    return false, util.ErrorKeyNotFound
  }

  return true, nil
}
