package main

import (
	"log"
	"time"

	"context"

	"github.com/jinzhu/gorm"

	// driver postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	pgdb "github.com/cipta-ageung/simas-db/proto"
	auth "github.com/cipta-ageung/simas-user/proto/auth"
	services "github.com/cipta-ageung/simas-user/services"
	"github.com/micro/go-micro"

	microclient "github.com/micro/go-micro/client"

	query1 "github.com/infobloxopen/atlas-app-toolkit/query"
)

var (
	db *gorm.DB
)

func main() {
	//db.DB().Ping()

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
	auth.RegisterAuthServiceHandler(service.Server(), &services.AuthService{})

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

func init() {
	// use the generated client stub
	serviceDb := pgdb.NewServiceConnectionService("go.micro.srv.simasdb", microclient.DefaultClient)

	svcDb, err := serviceDb.SetupDb(context.TODO(), &pgdb.ServiceApp{Svc: "go.micro.srv.simasuser"})
	if err != nil || svcDb == nil {
		log.Fatalf("cannot setup database")
	}

	db, err := gorm.Open("postgres", svcDb.ConnectionDb)
	if err != nil {
		log.Printf("error")
	}

	//	auth.DefaultCreateUser(context.Background(), &auth.User{Name: "cipta", SekolahId: "1", Email: "ci@gma.com", Password: "agi12345",
	//		UserRole: auth.UserRole_ROLE_ADMIN, Token: "asd"}, db)

	flds := &query1.FieldSelection{Fields: query1.FieldSelectionMap{"Name": &query1.Field{Name: "Name"}}}

	user := &auth.User{}
	user, err = auth.DefaultReadUser(context.Background(), &auth.User{Name: "cipta"}, db, flds)
	log.Println(user.GetName())
	defer db.Close()
}
