package memory

import (
	"github.com/anevsky/cachego/util"
)

func (cache *CACHE) SetString(key, value string) error {
	cache.Lock()
	defer cache.Unlock()

	cache.data[key] = value

	return nil
}

func (cache *CACHE) SetInt(key string, value int) error {
	cache.Lock()
	defer cache.Unlock()

	cache.data[key] = value

	return nil
}

func (cache *CACHE) SetList(key string, value util.List) error {
	cache.Lock()
	defer cache.Unlock()

	cache.data[key] = value

	return nil
}

func (cache *CACHE) SetDict(key string, value util.Dict) error {
	cache.Lock()
	defer cache.Unlock()

	cache.data[key] = value

	return nil
}

func (cache *CACHE) UpdateString(key, value string) (string, error) {
	cache.Lock()
	defer cache.Unlock()

	oldValue, success := cache.data[key]

	if !success {
		return "", util.ErrorKeyNotFound
	}

	cache.data[key] = value

	return oldValue.(string), nil
}

func (cache *CACHE) UpdateInt(key string, value int) (int, error) {
	cache.Lock()
	defer cache.Unlock()

	oldValue, success := cache.data[key]

	if !success {
		return -1, util.ErrorKeyNotFound
	}

	cache.data[key] = value

	return oldValue.(int), nil
}

func (cache *CACHE) UpdateList(key string, value util.List) (util.List, error) {
	cache.Lock()
	defer cache.Unlock()

	oldValue, success := cache.data[key]

	if !success {
		return nil, util.ErrorKeyNotFound
	}

	cache.data[key] = value

	return oldValue.(util.List), nil
}

func (cache *CACHE) UpdateDict(key string, value util.Dict) (util.Dict, error) {
	cache.Lock()
	defer cache.Unlock()

	oldValue, success := cache.data[key]

	if !success {
		return nil, util.ErrorKeyNotFound
	}

	cache.data[key] = value

	return oldValue.(util.Dict), nil
}

func (cache *CACHE) Remove(key string) error {
	cache.Lock()
	defer cache.Unlock()

	delete(cache.data, key)

	return nil
}

func (cache *CACHE) RemoveFromList(key string, value string) (int, error) {
	cache.Lock()
	defer cache.Unlock()

	list, success := cache.data[key]
	if !success {
		return 0, util.ErrorKeyNotFound
	}

	l, success := list.(util.List)
	if !success {
		return 0, util.ErrorWrongType
	}

	index := util.SentinelLinearSearch(l, value)
	if index != -1 {
		l = append(l[:index], l[index+1:]...)
	}

	cache.data[key] = l

	return index, nil
}

func (cache *CACHE) RemoveFromDict(key string, value string) error {
	cache.Lock()
	defer cache.Unlock()

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
	cache.Lock()
	defer cache.Unlock()

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

func (cache *CACHE) Increment(key string) (int, error) {
	cache.Lock()
	defer cache.Unlock()

	value, success := cache.data[key]
	if !success {
		return 0, util.ErrorKeyNotFound
	}

	v, success := value.(int)
	if !success {
		return 0, util.ErrorWrongType
	}

	cache.data[key] = v + 1

	return v + 1, nil
}
