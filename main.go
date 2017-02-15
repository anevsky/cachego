package main

import (
  "fmt"
  "time"
  "github.com/anevsky/cachego/memory"
  "github.com/anevsky/cachego/util"
)

func main() {
  fmt.Printf("Hello, Go!\n")

  a := []int{9, 4, 2, 8, 3, 5, 7, 10, 1, 6}
  fmt.Println(a)

  cache := memory.Alloc()

  cache.SetString("stringTest", "hi alex")
  v, err := cache.Get("stringTest")
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Println(v)

  for i := 0; i < 2; i++ {
    e2 := util.LogError{}
    e2.When = time.Now()
    fmt.Println(e2)
	}

  cache.SetInt("intTest", 123)
  v, err = cache.Get("intTest")
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Println(v)

  b := util.List{"one", "two"}
  cache.SetList("listTest", b)
  v, err = cache.Get("listTest")
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Println(v)

  c := make(util.Dict)
  c["k1"] = "v1"
  c["k2"] = "v2"
  cache.SetDict("dictTest", c)
  v, err = cache.Get("dictTest")
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Println(v)

  fmt.Println(cache.Keys())

  cache.Remove("dictTest")
  fmt.Println("Remove")

  fmt.Println(cache.Keys())
}
