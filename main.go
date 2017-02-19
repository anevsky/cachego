package main

import (
	"fmt"

	"github.com/anevsky/cachego/server"
)

func main() {
	fmt.Printf("Hello, Go!\n")

	server := server.Create()
	server.StartUp()

	/*
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

		stats, err := json.Marshal(cache.Stats())
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Println(string(stats))

		fmt.Println("TTL")
		cache.SetTTL("listTest", 100)
		time.Sleep(time.Millisecond*50 + time.Millisecond*20)
		_, err = cache.Get("listTest")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		time.Sleep(time.Millisecond*50 + time.Millisecond*20)
		_, err = cache.Get("listTest")
		if err != util.ErrorKeyNotFound {
			fmt.Printf("Expected ErrorKeyNotFound, but key found.")
		}
	*/
}
