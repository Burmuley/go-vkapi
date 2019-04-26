package go_vkapi

type VK interface {
	SendRequest(method string, parameters map[string]string) ([]byte, error)
}
