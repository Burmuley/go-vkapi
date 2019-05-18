package go_vkapi

import (
	//"gitlab.com/Burmuley/go-vkapi/objects/account"
	//"gitlab.com/Burmuley/go-vkapi/objects/account"
	"gitlab.com/Burmuley/go-vkapi/responses"
	//	"gitlab.com/Burmuley/go-vkapi/responses/account"
	//"strings"
)

type Account struct {
	*VKApi
}

//func (a *Account) GetInfo(fields ...string) (*account.Info, error) {
//	params := map[string]string{"fields": strings.Join(fields, ",")}
//	info := &account.Info{}
//
//	if err := a.SendObjRequest("account.getInfo", params, info); err != nil {
//		return &account.Info{}, err
//	}
//
//	return info, nil
//}
//
//func (a *Account) GetProfileInfo() (*account.UserSettings, error) {
//	profile := &account.UserSettings{}
//
//	if err := a.SendObjRequest("account.getProfileInfo", make(map[string]string, 0), profile); err != nil {
//		return &account.UserSettings{}, err
//	}
//
//	return profile, nil
//}
//
//func (a *Account) GetPushSettings(deviceId string) (*account.PushSettings, error) {
//	params := map[string]string{"device_id": deviceId}
//	pushSettings := &account.PushSettings{}
//
//	if err := a.SendObjRequest("account.getPushSettings", params, pushSettings); err != nil {
//		return &account.PushSettings{}, err
//	}
//
//	return pushSettings, nil
//}

func (a *Account) GetCounters() (*responses.AccountGetCounters, error) {
	//counters := &account.Counters{}
	resp := &responses.AccountGetCounters{}

	if err := a.SendObjRequest("account.getCounters", map[string]string{}, resp); err != nil {
		return &responses.AccountGetCounters{}, err
	}

	return resp, nil

}
