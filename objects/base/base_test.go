package base

import (
	"fmt"
	"os"
	"testing"
)

func getAPITokenEnv() (string, error) {
	token := os.Getenv("VK_API_TOKEN")
	if len(token) <= 0 {
		return "", fmt.Errorf("env variable VK_API_TOKEN not set")
	}

	return token, nil
}

func getAPIUserIDEnv() (string, error) {
	userId := os.Getenv("VK_API_USER_ID")
	if len(userId) <= 0 {
		return "", fmt.Errorf("env variable VK_API_USER_ID not set")
	}

	return userId, nil
}

func errorCheck(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestBaseBoolInt_ToJson(t *testing.T) {
	//var bbiTrue, bbiFalse BoolInt = true, false
	//
	//bbiTrueExpect, bbiFalseExpect := "1", "0"
	//
	//json.Unmarshal(&bbiTrue, bbiJsonTrueResult)
	//bbiJsonFalseResult, _ := bbiFalse.ToJson(&bbiFalse)
	//
	//if string(bbiJsonTrueResult) != bbiTrueExpect {
	//	t.Fail()
	//}
	//
	//if string(bbiJsonFalseResult) != bbiFalseExpect {
	//	t.Fail()
	//}
	//
	//fmt.Println(string(bbiJsonTrueResult), string(bbiJsonFalseResult))

	//vkk := NewApiWithToken("123")
	//
	//acc := &Account{vkk}
	//
	//acc.GetInfo()

	return
}

//func TestAccount_GetAccountInfo(t *testing.T) {
//	token, err := getAPITokenEnv()
//	//userId, err2 := getAPIUsersIDEnv()
//
//	if err != nil {
//		t.Error(err)
//		t.Fail()
//		return
//	}
//	//errorCheck(t, err)
//	//errorCheck(t, err2)
//
//	vkapi := go_vkapi.NewApiWithToken(token)
//
//	acc := &go_vkapi.Account{vkapi}
//
//	accInfo, err := acc.GetInfo("")
//
//	if err != nil {
//		t.Error(err)
//		t.Fail()
//		return
//	}
//
//	//jsn, _ := accInfo.ToJson(accInfo)
//	jsn2, _ := json.Marshal(accInfo)
//
//	fmt.Println(string(jsn2))
//
//}
