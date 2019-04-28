package go_vkapi

import (
	"encoding/json"
	"gitlab.com/Burmuley/go-vkapi/objects/account"
)

type Account struct {
	*VKApi
}

func (a *Account) GetAccountInfo(fields string) (*account.Info, error) {
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

func (a *Account) GetAccountProfile() (*account.UserSettings, error) {
	return &account.UserSettings{}, nil
}
