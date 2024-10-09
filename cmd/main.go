package main

import (
	"log"

	"github.com/golang_auth/db"
	"github.com/golang_auth/internal/user"
	"github.com/golang_auth/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Database Connection Failed...", err)
	}
	userRepo := user.NewUserRepository(dbConn.GetDB())
	userSvc := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
