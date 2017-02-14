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
  cache.data[key] = value
  return nil
}

func (cache *CACHE) UpdateInt(key string, value int) error {
  cache.data[key] = value
  return nil
}

func (cache *CACHE) UpdateList(key string, value util.List) error {
  cache.data[key] = value
  return nil
}

func (cache *CACHE) UpdateDict(key string, value util.Dict) error {
  cache.data[key] = value
  return nil
}

func (cache *CACHE) Remove(key string) error {
  delete(cache.data, key)
  return nil
}

func (cache *CACHE) RemoveFromList(key string, value string) error {
  return nil
}

func (cache *CACHE) RemoveFromDict(key string, value string) error {
  return nil
}
