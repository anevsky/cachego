package memory

import (
  "github.com/anevsky/cachego/util"
)

func (cache *CACHE) Get(key string) (interface{}, error) {
  value, success := cache.data[key]
  if !success {
    return "", util.ErrorKeyNotFound
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
    return "", util.ErrorWrongType
  }
}

func (cache *CACHE) GetListElement(key string, index int) (interface{}, error) {
  if index < 0 {
    return "", util.ErrorIndexOutOfBounds
  }

  value, success := cache.data[key]
  if !success {
    return "", util.ErrorKeyNotFound
  }

  v, success := value.(util.List)
  if !success {
    return "", util.ErrorWrongType
  }

  if index >= len(v) {
    return v[index], nil
  } else {
    return "", util.ErrorIndexOutOfBounds
  }
}

func (cache *CACHE) GetDictElement(key string, dKey string) (interface{}, error) {
  value, success := cache.data[key]
  if !success {
    return "", util.ErrorKeyNotFound
  }

  v, success := value.(util.Dict)
  if !success {
    return "", util.ErrorWrongType
  }

  e, success := v[dKey]
  if !success {
    return "", util.ErrorDictKeyNotFound
  }

  return e, nil
}

func (cache *CACHE) HasKey(key string) (bool, error) {
  if _, ok := cache.data[key]; !ok {
    return false, util.ErrorKeyNotFound
  }

  return true, nil
}
