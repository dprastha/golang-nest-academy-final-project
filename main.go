package main

import (
	"final-project/adaptor"
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

	rajaongkirAdaptor := adaptor.NewRajaOngkirAdaptor("https://api.rajaongkir.com/starter", "6d4d26125ea1c2991b801880cf3842f7")

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserServices(userRepo, *rajaongkirAdaptor)
	userHandler := controller.NewUserHandler(userService)

	productRepo := repository.NewProductRepo(db)
	productService := service.NewProductServices(productRepo)
	productHandler := controller.NewProductHandler(productService)

	transactionRepo := repository.NewTransactionRepo(db)
	transactionService := service.NewTransactionServices(transactionRepo, rajaongkirAdaptor)
	transactionHandler := controller.NewTranscationHandler(transactionService)

	router := gin.Default()
	app := server.NewRouter(router, userHandler, productHandler, transactionHandler)
	port := fmt.Sprintf(":%s", config.GetEnvVariable("APP_PORT"))

	app.Start(port)
}
