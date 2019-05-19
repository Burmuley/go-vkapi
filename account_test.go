package go_vkapi

import (
	"fmt"
	"testing"
)

func NewAccount() (*Account, error) {
	token, err := getAPITokenEnv()

	if err != nil {
		return &Account{}, err
	}

	Api := NewApiWithToken(token)

	return &Account{Api}, nil
}
func TestAccount_GetProfileInfo(t *testing.T) {
	VKAcc, err := NewAccount()

	if err != nil {
		t.Errorf("%s", err)
		t.Fail()
	}

	info, err := VKAcc.GetProfileInfo()

	if err != nil {
		t.Errorf("%s", err)
		t.Fail()
	}

	fmt.Println(info)
}

func TestAccount_GetCounters(t *testing.T) {
	VKAcc, err := NewAccount()

	if err != nil {
		t.Errorf("%s", err)
		t.Fail()
	}

	info, err := VKAcc.GetCounters()

	if err != nil {
		t.Errorf("%s", err)
		t.Fail()
	}

	fmt.Println(*info)
}

func TestAccount_GetInfo(t *testing.T) {
	VKAcc, err := NewAccount()

	if err != nil {
		t.Errorf("%s", err)
		t.Fail()
	}

	info, err := VKAcc.GetInfo()

	if err != nil {
		t.Errorf("%s", err)
		t.Fail()
	}

	fmt.Println(info)

}
