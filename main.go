package main

import (
	//"ecommerce/cmd"
	"ecommerce/util"
	"fmt"
)

func main(){
	//cmd.Serve()
	jwt, err := util.CreateJwt("my-secret", util.Payload{
		Sub:         "55",
		FirstName:   "Mojammel Haque",
		LastName:    "Bhuiyan",
		Email:       "mhbhuiyan10023@gmail.com",
		IsShopOwner: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(jwt)
}
