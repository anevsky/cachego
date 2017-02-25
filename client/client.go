package client

import (
	"fmt"

	"github.com/anevsky/cachego/util"
	"github.com/parnurzeal/gorequest"
)

type CLIENT struct {
	Url    string
	APIUrl string
	agent  *gorequest.SuperAgent
}

type Credentials struct {
	Username, Password string
}

func Create() CLIENT {
	request := gorequest.New()
	request.BasicAuth = Credentials{
		"alex",
		"juno",
	}

	cli := CLIENT{
		agent: request,
	}

	return cli
}

func (cli *CLIENT) Len() (result int, errs []error) {
	var value util.ResponseLen
	resp, body, errs := cli.agent.Get(cli.Url + cli.APIUrl + "/len").EndStruct(&value)

	if errs != nil {
		return 0, errs
	}

	fmt.Printf("\nResponse:\n%v \n", resp)
	fmt.Printf("\nBody:\n%s \n", body)
	fmt.Printf("\nErrors:\n%v \n", errs)
	fmt.Printf("\nValue:\n%d \n", value.Length)

	return value.Length, errs
}

func (cli *CLIENT) Keys() (result []string, errs []error) {
	var value util.ResponseKeys
	resp, body, errs := cli.agent.Get(cli.Url + cli.APIUrl + "/keys").EndStruct(&value)

	if errs != nil {
		return nil, errs
	}

	fmt.Printf("\nResponse:\n%v \n", resp)
	fmt.Printf("\nBody:\n%s \n", body)
	fmt.Printf("\nErrors:\n%v \n", errs)
	fmt.Printf("\nValue:\n%v \n", value.Keys)

	return value.Keys, errs
}

func (cli *CLIENT) Stats() (result util.Stats, errs []error) {
	var value util.ResponseStats
	resp, body, errs := cli.agent.Get(cli.Url + cli.APIUrl + "/stats").EndStruct(&value)

	if errs != nil {
		return util.Stats{}, errs
	}

	fmt.Printf("\nResponse:\n%v \n", resp)
	fmt.Printf("\nBody:\n%s \n", body)
	fmt.Printf("\nErrors:\n%v \n", errs)
	fmt.Printf("\nValue:\n%v \n", value.Stats)

	return value.Stats, errs
}

func (cli *CLIENT) Get(key string) (result interface{}, errs []error) {
	var value util.ResponseString
	resp, body, errs := cli.agent.Get(cli.Url + cli.APIUrl + "/get" + "/:" + key).EndStruct(&value)

	if errs != nil {
		return nil, errs
	}

	fmt.Printf("\nResponse:\n%v \n", resp)
	fmt.Printf("\nBody:\n%s \n", body)
	fmt.Printf("\nErrors:\n%v \n", errs)
	fmt.Printf("\nValue:\n%v \n", value.Value)

	return value.Value, errs
}
