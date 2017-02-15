package memory

import (
  "github.com/anevsky/cachego/util"
)

func (cache *CACHE) SetString(key, value string) error {
  cache.data[key] = value

  return nil
}

func (cache *CACHE) SetInt(key string, value int) error {
  cache.data[key] = value

  return nil
}

func (cache *CACHE) SetList(key string, value util.List) error {
  cache.data[key] = value

  return nil
}

func (cache *CACHE) SetDict(key string, value util.Dict) error {
  cache.data[key] = value

  return nil
}

func (cache *CACHE) UpdateString(key, value string) error {
  if _, success := cache.data[key]; !success {
    return util.ErrorKeyNotFound
  }

  cache.data[key] = value

  return nil
}

func (cache *CACHE) UpdateInt(key string, value int) error {
  if _, success := cache.data[key]; !success {
    return util.ErrorKeyNotFound
  }

  cache.data[key] = value

  return nil
}

func (cache *CACHE) UpdateList(key string, value util.List) error {
  if _, success := cache.data[key]; !success {
    return util.ErrorKeyNotFound
  }

  cache.data[key] = value

  return nil
}

func (cache *CACHE) UpdateDict(key string, value util.Dict) error {
  if _, success := cache.data[key]; !success {
    return util.ErrorKeyNotFound
  }

  cache.data[key] = value

  return nil
}

func (cache *CACHE) Remove(key string) error {
  delete(cache.data, key)

  return nil
}

func (cache *CACHE) RemoveFromList(key string, value string) error {
  list, success := cache.data[key]
  if !success {
    return util.ErrorKeyNotFound
  }

  l, success := list.(util.List)
  if !success {
    return util.ErrorWrongType
  }

  index := util.SentinelLinearSearch(l, value)
  if index != -1 {
    l = append(l[:index], l[index+1:]...)
  }

  return nil
}

func (cache *CACHE) RemoveFromDict(key string, value string) error {
  dict, success := cache.data[key]
  if !success {
    return util.ErrorKeyNotFound
  }

  d, success := dict.(util.Dict)
  if !success {
    return util.ErrorWrongType
  }

  delete(d, value)

  return nil
}

func (cache *CACHE) AppendToList(key, value string) error {
  list, success := cache.data[key]
  if !success {
    return util.ErrorKeyNotFound
  }

  l, success := list.(util.List)
  if !success {
    return util.ErrorWrongType
  }

  newList := append(l, value)

  cache.data[key] = newList

  return nil
}

func (cache *CACHE) Increment(key string) error {
  value, success := cache.data[key]
  if !success {
    return util.ErrorKeyNotFound
  }

  v, success := value.(int)
  if !success {
    return util.ErrorWrongType
  }

  cache.data[key] = v + 1

  return nil
}
