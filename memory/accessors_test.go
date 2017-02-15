package memory

import (
  "testing"
  "github.com/anevsky/cachego/util"
)

func TestGet(t *testing.T) {
  t.Log("Testing Get method...")

  cache := Alloc()

  cache.SetString("stringTest", "hi alex")
  _, err := cache.Get("stringTest")
  if err != nil {
    t.Error(err)
  }

  cache.SetInt("intTest", 123)
  _, err = cache.Get("intTest")
  if err != nil {
    t.Error(err)
  }

  b := util.List{"one", "two"}
  cache.SetList("listTest", b)
  _, err = cache.Get("listTest")
  if err != nil {
    t.Error(err)
  }

  c := make(util.Dict)
  c["k1"] = "v1"
  c["k2"] = "v2"
  cache.SetDict("dictTest", c)
  _, err = cache.Get("dictTest")
  if err != nil {
    t.Error(err)
  }
}

func TestGetListElement(t *testing.T) {
  t.Log("Testing Get method...")

  t.FailNow()
}

func TestGetDictElement(t *testing.T) {
  t.Log("Testing Get method...")

  t.FailNow()
}

func TestHasKey(t *testing.T) {
  t.Log("Testing HasKey method...")

  cache := Alloc()

  cache.SetString("stringTest", "hi alex")
  v, err := cache.HasKey("stringTest")
  if err != nil {
    t.Error(err)
  }
  if v != true {
    t.Errorf("Expected %t, but it was %t instead.", true, v)
  }
}
