package go_vkapi

import (
	"gitlab.com/Burmuley/go-vkapi/objects/account"
	"strings"
)

type Account struct {
	*VKApi
}

func (a *Account) GetInfo(fields ...string) (*account.Info, error) {
	params := map[string]string{"fields": strings.Join(fields, ",")}
	info := &account.Info{}

	if err := a.SendObjRequest("account.getInfo", params, info); err != nil {
		return &account.Info{}, err
	}

	return info, nil
}

func (a *Account) GetProfileInfo() (*account.UserSettings, error) {
	profile := &account.UserSettings{}

	if err := a.SendObjRequest("account.getProfileInfo", make(map[string]string, 0), profile); err != nil {
		return &account.UserSettings{}, err
	}

	return profile, nil
}

func (a *Account) GetPushSettings(deviceId string) (*account.PushSettings, error) {
	params := map[string]string{"device_id": deviceId}
	pushSettings := &account.PushSettings{}

	if err := a.SendObjRequest("account.getPushSettings", params, pushSettings); err != nil {
		return &account.PushSettings{}, err
	}

	return pushSettings, nil
}

func (a *Account) GetCounters() (*account.Counters, error) {
	counters := &account.Counters{}

	if err := a.SendObjRequest("account.getCounters", map[string]string{}, counters); err != nil {
		return &account.Counters{}, err
	}

	return counters, nil

}
