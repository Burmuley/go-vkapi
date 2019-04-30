package go_vkapi

import (
	"encoding/json"
	"gitlab.com/Burmuley/go-vkapi/objects/account"
)

type Account struct {
	*VKApi
}

func (a *Account) GetInfo(fields string) (*account.Info, error) {
	if len(fields) == 0 {
		fields = "country,lang"
	}

	params := map[string]string{"fields": fields}
	info, err := a.SendRequest("account.getInfo", params)

	if err != nil {
		return nil, err
	}

	var accInfo account.Info

	if err := json.Unmarshal(info, &accInfo); err != nil {
		return nil, err
	}

	return &accInfo, nil
}

func (a *Account) GetProfileInfo() (*account.UserSettings, error) {
	profile, err := a.SendRequest("account.getProfileInfo", make(map[string]string, 1))

	if err != nil {
		return &account.UserSettings{}, err
	}

	var profileInfo account.UserSettings

	if err := json.Unmarshal(profile, &profileInfo); err != nil {
		return &account.UserSettings{}, err
	}

	return &profileInfo, nil
}
