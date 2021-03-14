package main

import (
	"github.com/Juice/Project01/common-identity-srv/handler"
	common_identity "github.com/Juice/Project01/proto/common-identity"
	"github.com/micro/micro/v3/service"
)

func main(){
	srv := service.New(
		service.Name("common_identity"))

	srv.Init()

	_ = common_identity.RegisterCommonIdentityHandler(srv.Server(),&handler.Imp{})

	err := srv.Run()
	if err != nil{
		panic(err)
	}


}
