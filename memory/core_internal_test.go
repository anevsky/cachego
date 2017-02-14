package memory

import (
  "testing"
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
