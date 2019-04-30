package go_vkapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	apiVersion = "5.95"
	apiUrl     = "https://api.vk.com/method/"
)

type VKApi struct {
	userToken  string
	apiVersion string
	apiUrl     string
}

func (vk *VKApi) SendRequest(method string, parameters map[string]string) ([]byte, error) {
	//Format API endpoint
	u, err := url.Parse(apiUrl + method)

	if err != nil {
		return nil, err
	}

	//Fill mandatory parameters
	parameters["access_token"] = vk.userToken
	parameters["v"] = apiVersion

	// Format URL-encoded key-value parameters
	request := url.Values{}
	for k, v := range parameters {
		request.Add(k, v)
	}

	// Send request and read response
	resp, err := http.PostForm(u.String(), request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	//Read response body and check for errors
	rBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var apiResp ApiResponse

	if err := json.Unmarshal(rBody, &apiResp); err != nil {
		return nil, err
	}

	if apiResp.Error != nil {
		return nil, apiResp.Error
	}

	return apiResp.Response, nil
}

func NewApiWithToken(token string) *VKApi {
	envApiUrl := os.Getenv("VK_API_URL")
	locApiUrl := apiUrl

	if len(envApiUrl) > 0 {
		locApiUrl = envApiUrl
	}

	return &VKApi{userToken: token,
		apiVersion: apiVersion,
		apiUrl:     locApiUrl,
	}
}
