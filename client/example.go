package client

import (
	"fmt"

	"github.com/anevsky/cachego/util"
)

func testClient() {
	fmt.Printf("Hello, Client!\n")

	cli := Create()
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
