package go_vkapi

import (
	"encoding/json"
	"gitlab.com/Burmuley/go-vkapi/objects"
)

type Account struct {
	*VKApi
}

func (a *Account) GetAccountInfo(fields string) (*objects.AccountInfo, error) {
	if len(fields) == 0 {
		fields = "country,lang"
	}

	params := map[string]string{"fields": fields}
	info, err := a.SendRequest("account.getInfo", params)

	if err != nil {
		return nil, err
	}

	var accInfo objects.AccountInfo

	if err := json.Unmarshal(info, &accInfo); err != nil {
		return nil, err
	}

	return &accInfo, nil
}

func (a *Account) GetAccountProfile() (*objects.AccountUserSettings, error) {
	return &objects.AccountUserSettings{}, nil
}
