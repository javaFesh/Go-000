package main

import (
	"Week02/service"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func main() {
	//直接模拟调用
	userlist, err := service.UserService.GetUserByAge(123)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userlist)
	user, err := service.UserService.GetUserInfo(34234)
	if err != nil {
		log.Fatal(errors.Cause(err))
	} else {
		fmt.Println(user)
	}

}
