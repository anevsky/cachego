## cachego - Redis-like in-memory cache

[![GoDoc](https://godoc.org/github.com/anevsky/cachego?status.svg)](https://godoc.org/github.com/anevsky/cachego)
[![Report Card](https://goreportcard.com/badge/github.com/anevsky/cachego)](https://goreportcard.com/report/github.com/anevsky/cachego)

## Features:
- Key-value storage with string, lists, dict support
- Per-key TTL
- Operations:
  - Get
  - Set
  - Update
  - Remove
  - Keys
- Custom operations (Get i element on list, get value by key from dict, etc)
- Golang API client
- Telnet-like/HTTP-like API protocol
- Embed or client-server architecture

## Optional features
- auth (Done)
- persistence to disk/db (TODO)
- scaling (on server-side or on client-side) (TODO)
- perfomance tests (TODO)

## Run server

    go run $GOPATH/src/src/github.com/anevsky/cachego/main.go

## Run client

    # copy source code from client/example.go
    go run $GOPATH/src/github.com/anevsky/clientrun/main.go

## Use as embed cache storage

```Go
cache := memory.Alloc()

cache.SetString("stringTest", "hi alex")
v, err := cache.Get("stringTest")
if err != nil {
  fmt.Printf("Error: %v\n", err)
}
fmt.Println(v)
```

## Use as server cache storage

```Go
server := server.Create()
server.StartUp()
```

## cURL examples to server

* Get total number of objects 
* `curl -i -w "\n" --user alex:secret localhost:8027/v1/len`
* Get list of keys 
* `curl -i -w "\n" --user alex:secret localhost:8027/v1/keys`
* Get cache stats 
* `curl -i -w "\n" --user alex:secret localhost:8027/v1/stats`
* Get value from cache by key 
* `curl -i -w "\n" --user alex:secret localhost:8027/v1/get/vvv`
* Get element from list by index 
* `curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":1}' localhost:8027/v1/list/element/lll`
* Get element from dict by key 
* `curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":"k1"}' localhost:8027/v1/dict/element/ddd`
* Check if object exists in cache by key 
* `curl -i -w "\n" --user alex:secret localhost:8027/v1/key/lll`
* Set string 
* `curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":"s1"}' localhost:8027/v1/string/sss`
* Set int 
* `curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":121}' localhost:8027/v1/int/iii`
* Set list 
* `curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":["aa", "bb"]}' localhost:8027/v1/list/lll`
* Set dict 
* `curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":{"k1": "v1", "k2": "v2"}}' localhost:8027/v1/dict/ddd`
* Update string by key 
* `curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":"s2"}' localhost:8027/v1/string/sss`
* Update int by key 
* `curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":123}' localhost:8027/v1/int/iii`
* Update list by key 
* `curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":["aa2", "bb2"]}' localhost:8027/v1/list/lll`
* Update dict by key 
* `curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":{"k12": "v12", "k22": "v22"}}' localhost:8027/v1/dict/ddd`
* Append to list a string element 
* `curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json' -d '{"value":"aa3"}' localhost:8027/v1/list/element/lll`
* Increment an integer value by key 
* `curl -i -w "\n" -X PUT --user alex:secret -H 'Content-Type: application/json'  localhost:8027/v1/int/increment/iii`
* Remove object from cache by key 
* `curl -i -w "\n" -X DELETE --user alex:secret -H 'Content-Type: application/json'  localhost:8027/v1/remove/iii`
* Remove object from list by value 
* `curl -i -w "\n" -X DELETE --user alex:secret -H 'Content-Type: application/json' -d '{"value":"aa3"}' localhost:8027/v1/list/element/lll`
* Remove object from dict by key 
* `curl -i -w "\n" -X DELETE --user alex:secret -H 'Content-Type: application/json' -d '{"value":"k12"}' localhost:8027/v1/dict/element/ddd`
* Set TTL (time-to-live) in nanoseconds for object by key 
* `curl -i -w "\n" -X POST --user alex:secret -H 'Content-Type: application/json' -d '{"value":5211}' localhost:8027/v1/ttl/iii`

## Client example

```Go
package main

import (
	"fmt"

	"github.com/anevsky/cachego/client"
	"github.com/anevsky/cachego/util"
)

func main() {
	fmt.Printf("Hello, Client!\n")

	cli := client.Create()
	cli.Url = "http://localhost:8027"
	cli.APIUrl = "/v1"

	//

	v1, errs := cli.Len()
	printInfo(v1, errs)

	v2, errs := cli.Keys()
	printInfo(v2, errs)

	v3, errs := cli.Stats()
	printInfo(v3, errs)

	v4, errs := cli.GetString("test")
	printInfo(v4, errs)

	errs = cli.SetInt("vvv", 21)
	printInfo(nil, errs)

	v5, errs := cli.GetInt("vvv")
	printInfo(v5, errs)

	errs = cli.SetInt("vvv", 29)
	printInfo(nil, errs)

	v6, errs := cli.GetInt("vvv")
	printInfo(v6, errs)

	//

	errs = cli.SetString("sss", "s1")
	printInfo(nil, errs)

	errs = cli.SetInt("iii", 21)
	printInfo(nil, errs)

	errs = cli.SetList("lll", util.List{"aa", "bb"})
	printInfo(nil, errs)

	errs = cli.SetDict("ddd", util.Dict{"k1": "v1", "k2": "v2"})
	printInfo(nil, errs)

	//

	v7, errs := cli.GetString("sss")
	printInfo(v7, errs)

	v8, errs := cli.GetInt("iii")
	printInfo(v8, errs)

	v9, errs := cli.GetListElement("lll", 1)
	printInfo(v9, errs)

	v10, errs := cli.GetDictElement("ddd", "k2")
	printInfo(v10, errs)

	v11, errs := cli.HasKey("iii")
	printInfo(v11, errs)

	//

	v12, errs := cli.UpdateString("sss", "su2")
	printInfo(v12, errs)

	v13, errs := cli.UpdateInt("iii", 222)
	printInfo(v13, errs)

	v14, errs := cli.UpdateList("lll", util.List{"cc", "dd", "gg"})
	printInfo(v14, errs)

	v15, errs := cli.UpdateDict("ddd", util.Dict{"k12": "v12", "k22": "v22"})
	printInfo(v15, errs)

	errs = cli.AppendToList("lll", "ww")
	printInfo(nil, errs)

	v16, errs := cli.Increment("iii")
	printInfo(v16, errs)

	v17, errs := cli.Remove("iii")
	printInfo(v17, errs)

	v18, errs := cli.RemoveFromList("lll", "gg")
	printInfo(v18, errs)

	v19, errs := cli.RemoveFromDict("ddd", "k22")
	printInfo(v19, errs)

	v20, errs := cli.SetTTL("lll", 7500)
	printInfo(v20, errs)
}

func printInfo(value interface{}, errs []error) {
	if errs != nil {
		fmt.Printf("#Errors: %v \n", errs)
	} else {
		switch v := value.(type) {
		case int:
			fmt.Printf("#Result: %d \n", v)
		case string:
			fmt.Printf("#Result: %s \n", v)
		case util.List:
			fmt.Printf("#Result: %v \n", v)
		case util.Dict:
			fmt.Printf("#Result: %v \n", v)
		default:
			fmt.Printf("#Result: %v \n", v)
		}
	}
}
```
