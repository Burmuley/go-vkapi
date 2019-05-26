package go_vkapi

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/responses"
)

// Auth represents collection of methods related to VK Auth
type Auth struct {
	*VKApi
}

// CheckPhone Checks a user's phone number for correctness.
//
// Parameters:
//  * phone - Phone number.
//  * clientId - User ID.
//  * clientSecret -
//  * authByPhone -

func (a *Auth) CheckPhone(phone, clientSecret string, clientId int, authByPhone objects.BaseBoolInt) (responses.OkResponse, error) {
	params := map[string]string{"phone": phone,
		"client_secret": clientSecret,
		"client_id":     string(clientId),
		"auth_by_phone": fmt.Sprint(authByPhone),
	}

	var resp responses.OkResponse

	if err := a.SendObjRequest("auth.checkPhone", params, &resp); err != nil {
		return responses.OkResponse(0), err
	}

	return resp, nil
}

// Restore Allows to restore account access using a code received via SMS. " This method is only available for apps with [vk.com/dev/auth_direct|Direct authorization] access. "
//
// Parameters:
//  * phone - User phone number.
//  * last_name - User last name.
func (a *Auth) Restore(phone, lastName string) (responses.AuthRestore, error) {
	params := map[string]string{"phone": phone, "last_name": lastName}

	var resp responses.AuthRestore

	if err := a.SendObjRequest("auth.restore", params, &resp); err != nil {
		return responses.AuthRestore{}, err
	}

	return resp, nil
}
