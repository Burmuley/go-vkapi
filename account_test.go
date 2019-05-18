package go_vkapi

import (
	"fmt"
	"testing"
)

//func TestAccount_GetProfileInfo(t *testing.T) {
//	token, err := getAPITokenEnv()
//
//	if err != nil {
//		t.Errorf("%s", err)
//		t.Fail()
//	}
//
//	Api := NewApiWithToken(token)
//
//	VKAcc := &Account{Api}
//
//	info, err := VKAcc.GetProfileInfo()
//
//	if err != nil {
//		t.Errorf("%s", err)
//		t.Fail()
//	}
//
//	fmt.Println(info)
//}

func TestAccount_GetCounters(t *testing.T) {
	token, err := getAPITokenEnv()

	if err != nil {
		t.Errorf("%s", err)
		t.Fail()
		return
	}

	Api := NewApiWithToken(token)

	VKAcc := &Account{Api}

	info, err := VKAcc.GetCounters()

	if err != nil {
		t.Errorf("%s", err)
		t.Fail()
	}

	fmt.Println(*info)
}

//func TestAccount_GetInfo(t *testing.T) {
//	token, err := getAPITokenEnv()
//
//	if err != nil {
//		t.Errorf("%s", err)
//		t.Fail()
//		return
//	}
//
//	Api := NewApiWithToken(token)
//
//	VKAcc := &Account{Api}
//
//	info, err := VKAcc.GetInfo()
//
//	if err != nil {
//		t.Errorf("%s", err)
//		t.Fail()
//	}
//
//	fmt.Println(info)
//
//}
