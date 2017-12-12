# vk-api-go
GO library for vk api

##Example
 
````
package main

import (
	"fmt"
	apigo "github.com/drossha/vk-api-go"
	"github.com/drossha/vk-api-go/users"
)

func main() {
	client := apigo.NewClient()

	service := users.New(&client)

	data, err := service.Get(users.GetRequest{
		UserIds: "100",
	})
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("resp: ", string(data))
}
````