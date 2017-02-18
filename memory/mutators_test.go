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

  cache := Alloc()

  _, err := cache.UpdateString("stringTest", "hi alex")
  if err != util.ErrorKeyNotFound {
    t.Errorf("Expected ErrorKeyNotFound, but it was %s instead.", err)
  }

  cache.SetString("stringTest", "hi alex")

  _, err = cache.UpdateString("stringTest", "hi 2")
  if err != nil {
    t.Error(err)
  }
}

func TestUpdateInt(t *testing.T) {
  t.Log("Testing Update method...")

  cache := Alloc()

  _, err := cache.UpdateInt("intTest", 123)
  if err != util.ErrorKeyNotFound {
    t.Errorf("Expected ErrorKeyNotFound, but it was %s instead.", err)
  }

  cache.SetInt("intTest", 123)

  _, err = cache.UpdateInt("intTest", 123)
  if err != nil {
    t.Error(err)
  }
}

func TestUpdateList(t *testing.T) {
  t.Log("Testing Update method...")

  cache := Alloc()

  b := util.List{"one", "two"}

  _, err := cache.UpdateList("listTest", b)
  if err != util.ErrorKeyNotFound {
    t.Errorf("Expected ErrorKeyNotFound, but it was %s instead.", err)
  }

  cache.SetList("listTest", b)

  _, err = cache.UpdateList("listTest", b)
  if err != nil {
    t.Error(err)
  }
}

func TestUpdateDict(t *testing.T) {
  t.Log("Testing Update method...")

  cache := Alloc()

  c := make(util.Dict)
  c["k1"] = "v1"
  c["k2"] = "v2"

  _, err := cache.UpdateDict("dictTest", c)
  if err != util.ErrorKeyNotFound {
    t.Errorf("Expected ErrorKeyNotFound, but it was %s instead.", err)
  }

  cache.SetDict("dictTest", c)

  _, err = cache.UpdateDict("dictTest", c)
  if err != nil {
    t.Error(err)
  }
}

func TestRemove(t *testing.T) {
  t.Log("Testing Remove method...")

  cache := Alloc()

  cache.SetString("stringTest", "hi alex")

  c := make(util.Dict)
  c["k1"] = "v1"
  c["k2"] = "v2"
  cache.SetDict("dictTest", c)

  cache.Remove("stringTest")
  cache.Remove("dictTest")

  _, err := cache.Get("stringTest")
  if err != util.ErrorKeyNotFound {
    t.Errorf("Expected ErrorKeyNotFound, but it was %s instead.", err)
  }

  _, err = cache.Get("dictTest")
  if err != util.ErrorKeyNotFound {
    t.Errorf("Expected ErrorKeyNotFound, but it was %s instead.", err)
  }
}

func TestRemoveFromList(t *testing.T) {
  t.Log("Testing Remove method...")

  cache := Alloc()

  b := util.List{"one", "two"}
  cache.SetList("listTest", b)

  index, err := cache.RemoveFromList("listTest", "two")
  if err != nil {
    t.Error(err)
  }
  if index != 1 {
    t.Errorf("Expected %d, but it was %d instead.", 1, index)
  }

  v, err := cache.Get("listTest")
  if err != nil {
    t.Error(err)
  }

  c := util.List{"one"}
  if !reflect.DeepEqual(v, c) {
    t.Errorf("Expected %s, but it was %s instead.", c, v)
  }
}

func TestRemoveFromDict(t *testing.T) {
  t.Log("Testing Remove method...")

  cache := Alloc()

  c := make(util.Dict)
  c["k1"] = "v1"
  c["k2"] = "v2"
  cache.SetDict("dictTest", c)

  cache.RemoveFromDict("dictTest", "k2")

  v, err := cache.Get("dictTest")
  if err != nil {
    t.Error(err)
  }

  d := make(util.Dict)
  d["k1"] = "v1"

  if !reflect.DeepEqual(v, d) {
    t.Errorf("Expected %s, but it was %s instead.", d, v)
  }
}

func TestAppendToList(t *testing.T) {
  t.Log("Testing Append method...")

  cache := Alloc()

  b := util.List{"one", "two"}
  cache.SetList("listTest", b)

  cache.AppendToList("listTest", "three")

  v, err := cache.Get("listTest")
  if err != nil {
    t.Error(err)
  }

  c := util.List{"one", "two", "three"}
  if !reflect.DeepEqual(v, c) {
    t.Errorf("Expected %s, but it was %s instead.", c, v)
  }
}

func TestIncrement(t *testing.T) {
  t.Log("Testing Increment method...")

  cache := Alloc()

  cache.SetInt("intTest", 123)

  cache.Increment("intTest")

  v, err := cache.Get("intTest")
  if err != nil {
    t.Error(err)
  }
  if v != 124 {
    t.Errorf("Expected 124, but it was %d instead.", v)
  }
}
