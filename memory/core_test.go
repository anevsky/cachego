package memory

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/anevsky/cachego/util"
)

func TestAlloc(t *testing.T) {
	t.Log("Testing Alloc method...")

	cache := Alloc()
	if cache.data == nil {
		t.Errorf("Allocation failed.")
	}
}

func TestLen(t *testing.T) {
	t.Log("Testing Len method...")

	cache := Alloc()
	cache.SetString("stringTest", "hi alex")
	cache.SetInt("intTest", 123)
	if cache.Len() != 2 {
		t.Errorf("Expected 2, but it was %d instead.", cache.Len())
	}
}

func TestKeys(t *testing.T) {
	t.Log("Testing Keys method...")

	cache := Alloc()
	cache.SetString("stringTest", "hi alex")
	cache.SetInt("intTest", 123)
	k := cache.Keys()

	r := false
	for _, kk := range k {
		if kk == "stringTest" {
			r = true
		}
	}
	if r != true {
		t.Errorf("Desired key 'stringTest' not found in %s.", cache.Keys())
	}

	r = false
	for _, kk := range k {
		if kk == "intTest" {
			r = true
		}
	}
	if r != true {
		t.Errorf("Desired key 'intTest' not found in %s.", cache.Keys())
	}
}

func TestMutex(t *testing.T) {
	cache := Alloc()

	cache.SetInt("intTest", 1)

	for i := 0; i < 30; i++ {
		go cache.Increment("intTest")
	}

	_, err := cache.Get("intTest")
	if err != nil {
		t.Error(err)
	}
}

func TestStats(t *testing.T) {
	cache := Alloc()

	cache.SetString("stringTest", "hi alex")
	cache.SetInt("intTest", 123)

	for i := 0; i < 30; i++ {
		cache.Increment("intTest")
	}

	stats, err := json.Marshal(cache.Stats())
	if err != nil {
		t.Error(err)
	}

	m := make(map[string]int64)
	err = json.Unmarshal(stats, &m)
	if err != nil {
		t.Error(err)
	}

	if m["num_gc"] != 0 {
		t.Errorf("Expected %s, but it was %d instead.", "0", m["num_gc"])
	}
}

func TestSetTTL(t *testing.T) {
	t.Log("Testing TTL method...")

	cache := Alloc()

	cache.SetInt("intTest", 123)
	cache.SetTTL("intTest", 500)
	time.Sleep(time.Millisecond*500 + time.Millisecond*20)
	_, err := cache.Get("intTest")
	if err != util.ErrorKeyNotFound {
		t.Error("Expected ErrorKeyNotFound, but key found.")
	}

	cache.SetInt("intTest2", 123)
	cache.SetTTL("intTest2", 100)
	time.Sleep(time.Millisecond*100 + time.Millisecond*20)
	_, err = cache.Get("intTest2")
	if err != util.ErrorKeyNotFound {
		t.Error("Expected ErrorKeyNotFound, but key found.")
	}

	cache.SetInt("intTest3", 123)
	cache.SetTTL("intTest3", 100)
	time.Sleep(time.Millisecond * 50)
	_, err = cache.Get("intTest3")
	if err != nil {
		t.Errorf("Expected nil error, but it was %v instead.", err)
	}

	time.Sleep(time.Millisecond*50 + time.Millisecond*20)
	_, err = cache.Get("intTest3")
	if err != util.ErrorKeyNotFound {
		t.Error("Expected ErrorKeyNotFound, but key found.")
	}
}
