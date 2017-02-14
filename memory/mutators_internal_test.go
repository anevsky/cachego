package memory

import (
  "testing"
  "reflect"
  "github.com/anevsky/cachego/util"
)

func TestSetString(t *testing.T) {
  t.Log("Testing Set method...")

  cache := Alloc()

  cache.SetString("stringTest", "hi alex")
  v, err := cache.Get("stringTest")
  if err != nil {
    t.Error(err)
  }
  if v != "hi alex" {
    t.Errorf("Expected 'hi alex', but it was '%s' instead.", v)
  }
}

func TestSetInt(t *testing.T) {
  t.Log("Testing Set method...")

  cache := Alloc()

  cache.SetInt("intTest", 123)
  v, err := cache.Get("intTest")
  if err != nil {
    t.Error(err)
  }
  if v != 123 {
    t.Errorf("Expected 123, but it was %d instead.", v)
  }
}

func TestSetList(t *testing.T) {
  t.Log("Testing Set method...")

  cache := Alloc()

  b := util.List{"one", "two"}
  cache.SetList("listTest", b)
  v, err := cache.Get("listTest")
  if err != nil {
    t.Error(err)
  }
  if !reflect.DeepEqual(v, b) {
    t.Errorf("Expected %s, but it was %s instead.", b, v)
  }
}

func TestSetDict(t *testing.T) {
  t.Log("Testing Set method...")

  cache := Alloc()

  c := make(util.Dict)
  c["k1"] = "v1"
  c["k2"] = "v2"
  cache.SetDict("dictTest", c)
  v, err := cache.Get("dictTest")
  if err != nil {
    t.Error(err)
  }
  if !reflect.DeepEqual(v, c) {
    t.Errorf("Expected %s, but it was %s instead.", c, v)
  }
}

func TestUpdateString(t *testing.T) {
  t.Log("Testing Set method...")
}

func TestUpdateInt(t *testing.T) {
  t.Log("Testing Update method...")
}

func TestUpdateList(t *testing.T) {
  t.Log("Testing Update method...")
}

func TestUpdateDict(t *testing.T) {
  t.Log("Testing Update method...")
}

func TestRemove(t *testing.T) {
  t.Log("Testing Remove method...")
}

func TestRemoveFromList(t *testing.T) {
  t.Log("Testing Remove method...")
}

func TestRemoveFromDict(t *testing.T) {
  t.Log("Testing Remove method...")
}
