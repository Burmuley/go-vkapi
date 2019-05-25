# VK (Vkontakte) API library for Golang

This repo contains library with implementation of the [VK](https://vk.com/) public API social network for [Golang](https://www.golang.org). 

All platform objects and methods were implemented according to the public [JSON schema](https://github.com/VKCOM/vk-api-schema/) in official VK repository.

## Repo structure
 * dir `errors` - package contains VK errors representation
 * dir `objects` - packages contains Go structures representing VK API objects, ready for marshaling/unmarshaling from/to JSON
 * dir `responses` - package contains Go structures representing VK API responses
 * `api.go` - contains  implementation of the `VK` interface for basic functionality
 * `<method name>.go` - file contains implementation of all methods related to appropriate API method

## Examples 

### Account operations
```gotemplate
package main

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi"
)

package main

func main() {
  token := "VK API token"
  Api := NewApiWithToken(token)
  VKAccount := Account{*Api}

  AccountInfo, err := VKAccount.GetProfileInfo()

  fmt.Println(AccountInfo)	
}
```
