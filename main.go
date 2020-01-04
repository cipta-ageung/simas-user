package main

import (
	"log"
	"time"

	"context"

	_ "github.com/jinzhu/gorm"

	// driver postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	cfg "github.com/cipta-ageung/simas-db/config"
	pgdb "github.com/cipta-ageung/simas-db/proto"
	auth "github.com/cipta-ageung/simas-user/proto/auth"
	services "github.com/cipta-ageung/simas-user/services"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

func main() {

	log.Println("starting services user . . .")

	serviceClient := grpc.NewService()
	serviceClient.Init()

	// use the generated client stub
	serviceDb := pgdb.NewServiceConnectionService("go.micro.srv.simasdb", serviceClient.Client())
	if serviceDb != nil {
		log.Println("fuck")
	}
	svcDb, err := serviceDb.SetupDb(context.TODO(), &pgdb.ServiceApp{Svc: "go.micro.srv.simasuser"})
	if err != nil || svcDb == nil {
		log.Fatalf("cannot setup database")
	}

	db, err := cfg.Connect(svcDb.ConnectionDb)
	if err != nil {
		log.Fatal("cannot connet database")
	}

	ctx := context.WithValue(context.Background(), "db", db)

	log.Println(ctx)

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

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
