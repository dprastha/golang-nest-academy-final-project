package main

import (
	"final-project/config"
	"final-project/db"
	"final-project/server"
	"final-project/server/controller"
	"final-project/server/repository"
	"final-project/server/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO assign db value to repository
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)
	userService := service.NewService(userRepo)
	userHandler := controller.NewUserHandler(userService)

	router := gin.Default()
	app := server.NewRouter(router, userHandler)
	port := fmt.Sprintf(":%s", config.GetEnvVariable("APP_PORT"))

	app.Start(port)
}
