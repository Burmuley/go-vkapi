package go_vkapi

import (
	"fmt"
	"os"
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
