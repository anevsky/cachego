package client

import (
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
		"secret",
	}

	cli := CLIENT{
		agent: request,
	}

	return cli
}

func (cli *CLIENT) Len() (result int, errs []error) {
	var dto util.LenDTO
	resp, body, errs := cli.agent.
		Get(cli.Url + cli.APIUrl + "/len").
		EndStruct(&dto)

	if errs != nil {
		return 0, errs
	}

	if resp == nil || body == nil {
		return 0, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Length, errs
}

func (cli *CLIENT) Keys() (result []string, errs []error) {
	var dto util.KeysDTO
	resp, body, errs := cli.agent.
		Get(cli.Url + cli.APIUrl + "/keys").
		EndStruct(&dto)

	if errs != nil {
		return nil, errs
	}

	if resp == nil || body == nil {
		return nil, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Keys, errs
}

func (cli *CLIENT) Stats() (result util.Stats, errs []error) {
	var dto util.StatsDTO
	resp, body, errs := cli.agent.
		Get(cli.Url + cli.APIUrl + "/stats").
		EndStruct(&dto)

	if errs != nil {
		return util.Stats{}, errs
	}

	if resp == nil || body == nil {
		return util.Stats{}, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Stats, errs
}

func (cli *CLIENT) GetString(key string) (result string, errs []error) {
	var dto util.StringDTO
	resp, body, errs := cli.agent.
		Get(cli.Url + cli.APIUrl + "/get/" + key).
		EndStruct(&dto)

	if errs != nil {
		return "", errs
	}

	if resp == nil || body == nil {
		return "", []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) GetInt(key string) (result int, errs []error) {
	var dto util.IntDTO
	resp, body, errs := cli.agent.
		Get(cli.Url + cli.APIUrl + "/get/" + key).
		EndStruct(&dto)

	if errs != nil {
		return -1, errs
	}

	if resp == nil || body == nil {
		return -1, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) GetListElement(key string) (result string, errs []error) {
	var dto util.StringDTO
	resp, body, errs := cli.agent.
		Get(cli.Url + cli.APIUrl + "/list/element/" + key).
		EndStruct(&dto)

	if errs != nil {
		return "", errs
	}

	if resp == nil || body == nil {
		return "", []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) GetDictElement(key string) (result string, errs []error) {
	var dto util.StringDTO
	resp, body, errs := cli.agent.
		Get(cli.Url + cli.APIUrl + "/dict/element/" + key).
		EndStruct(&dto)

	if errs != nil {
		return "", errs
	}

	if resp == nil || body == nil {
		return "", []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) HasKey(key string) (result bool, errs []error) {
	var dto util.BoolDTO
	resp, body, errs := cli.agent.
		Get(cli.Url + cli.APIUrl + "/key/" + key).
		EndStruct(&dto)

	if errs != nil {
		return false, errs
	}

	if resp == nil || body == nil {
		return false, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) SetString(key, v string) (errs []error) {
	var dto util.BasicDTO
	resp, body, errs := cli.agent.
		Post(cli.Url + cli.APIUrl + "/string/" + key).
		Send(util.StringDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return errs
	}

	if resp == nil || body == nil {
		return []error{util.ErrorResponseOrBodyNil}
	}

	return errs
}

func (cli *CLIENT) SetInt(key string, v int) (errs []error) {
	var dto util.BasicDTO
	resp, body, errs := cli.agent.
		Post(cli.Url+cli.APIUrl+"/int/"+key).
		Set("Notes", "gorequst is coming!"). // Header
		//Send(`{"value":"` + strconv.Itoa(v) + `"}`). // JSON
		Send(util.IntDTO{Value: v}). // JSON
		EndStruct(&dto)

	if errs != nil {
		return errs
	}

	if resp == nil || body == nil {
		return []error{util.ErrorResponseOrBodyNil}
	}

	return errs
}

func (cli *CLIENT) SetList(key string, v util.List) (errs []error) {
	var dto util.BasicDTO
	resp, body, errs := cli.agent.
		Post(cli.Url + cli.APIUrl + "/list/" + key).
		Send(util.ListDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return errs
	}

	if resp == nil || body == nil {
		return []error{util.ErrorResponseOrBodyNil}
	}

	return errs
}

func (cli *CLIENT) SetDict(key string, v util.Dict) (errs []error) {
	var dto util.BasicDTO
	resp, body, errs := cli.agent.
		Post(cli.Url + cli.APIUrl + "/dict/" + key).
		Send(util.DictDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return errs
	}

	if resp == nil || body == nil {
		return []error{util.ErrorResponseOrBodyNil}
	}

	return errs
}

func (cli *CLIENT) UpdateString(key, v string) (result string, errs []error) {
	var dto util.StringDTO
	resp, body, errs := cli.agent.
		Put(cli.Url + cli.APIUrl + "/string/" + key).
		Send(util.StringDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return "", errs
	}

	if resp == nil || body == nil {
		return "", []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) UpdateInt(key string, v int) (result int, errs []error) {
	var dto util.IntDTO
	resp, body, errs := cli.agent.
		Put(cli.Url + cli.APIUrl + "/int/" + key).
		Send(util.IntDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return -1, errs
	}

	if resp == nil || body == nil {
		return -1, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) UpdateList(key string, v util.List) (result util.List, errs []error) {
	var dto util.ListDTO
	resp, body, errs := cli.agent.
		Put(cli.Url + cli.APIUrl + "/list/" + key).
		Send(util.ListDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return nil, errs
	}

	if resp == nil || body == nil {
		return nil, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) UpdateDict(key string, v util.Dict) (result util.Dict, errs []error) {
	var dto util.DictDTO
	resp, body, errs := cli.agent.
		Put(cli.Url + cli.APIUrl + "/dict/" + key).
		Send(util.DictDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return nil, errs
	}

	if resp == nil || body == nil {
		return nil, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) AppendToList(key, v string) (errs []error) {
	var dto util.BasicDTO
	resp, body, errs := cli.agent.
		Put(cli.Url + cli.APIUrl + "/list/element/" + key).
		Send(util.StringDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return errs
	}

	if resp == nil || body == nil {
		return []error{util.ErrorResponseOrBodyNil}
	}

	return errs
}

func (cli *CLIENT) Increment(key string) (result int, errs []error) {
	var dto util.IntDTO
	resp, body, errs := cli.agent.
		Put(cli.Url + cli.APIUrl + "/int/increment/" + key).
		EndStruct(&dto)

	if errs != nil {
		return -1, errs
	}

	if resp == nil || body == nil {
		return -1, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) Remove(key string) (result int, errs []error) {
	var dto util.BasicDTO
	resp, body, errs := cli.agent.
		Delete(cli.Url + cli.APIUrl + "/remove/" + key).
		EndStruct(&dto)

	if errs != nil {
		return -1, errs
	}

	if resp == nil || body == nil {
		return -1, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.ErrorCode, errs
}

func (cli *CLIENT) RemoveFromList(key, v string) (result int, errs []error) {
	var dto util.IntDTO
	resp, body, errs := cli.agent.
		Delete(cli.Url + cli.APIUrl + "/list/element/" + key).
		Send(util.StringDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return -1, errs
	}

	if resp == nil || body == nil {
		return -1, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.Value, errs
}

func (cli *CLIENT) RemoveFromDict(key, v string) (result int, errs []error) {
	var dto util.BasicDTO
	resp, body, errs := cli.agent.
		Delete(cli.Url + cli.APIUrl + "/dict/element/" + key).
		Send(util.StringDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return -1, errs
	}

	if resp == nil || body == nil {
		return -1, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.ErrorCode, errs
}

func (cli *CLIENT) SetTTL(key string, v int) (result int, errs []error) {
	var dto util.BasicDTO
	resp, body, errs := cli.agent.
		Post(cli.Url + cli.APIUrl + "/ttl/" + key).
		Send(util.IntDTO{Value: v}).
		EndStruct(&dto)

	if errs != nil {
		return -1, errs
	}

	if resp == nil || body == nil {
		return -1, []error{util.ErrorResponseOrBodyNil}
	}

	return dto.ErrorCode, errs
}
