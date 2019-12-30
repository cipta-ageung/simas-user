package main

import (
	"log"
	"time"

	auth "github.com/cipta-ageung/simas-user/proto/auth"
	services "github.com/cipta-ageung/simas-user/services"
	"github.com/micro/go-micro"
)

func main() {

	log.Println("starting services user . . .")

	service := micro.NewService(
		micro.Name("go.micro.srv.simasuser"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "simas_user",
		}),
	)

	service.Init()
	auth.RegisterAuthServiceHandler(service.Server(), new(services.AuthService))

}
