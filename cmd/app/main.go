package main

import (
	"freepass-bcc/internal/handler/rest"
	"freepass-bcc/internal/repository"
	"freepass-bcc/internal/service"
	"freepass-bcc/pkg/bcrypt"
	"freepass-bcc/pkg/config"
	"freepass-bcc/pkg/database/mariadb"
	"freepass-bcc/pkg/jwt"
	"freepass-bcc/pkg/middleware"
	"log"
)

func main() {
	config.LoadEnvironment()

	db, err := mariadb.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = mariadb.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()
	svc := service.NewService(repo, bcrypt, jwt)
	middleware := middleware.Init(svc, jwt)

	r := rest.NewRest(svc, middleware)
	r.MountEndpoint()
	r.Run()
}
