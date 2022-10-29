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

	rajaongkirAdaptor := adaptor.NewRajaOngkirAdaptor(config.GetEnvVariable("RAJA_ONGKIR_HOST"), config.GetEnvVariable("RAJA_ONGKIR_API_KEY"))

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
	middleware := server.NewMiddleware(userService)

	app := server.NewRouter(router, userHandler, productHandler, transactionHandler, middleware)
	port := fmt.Sprintf(":%s", config.GetEnvVariable("APP_PORT"))

	app.Start(port)
}
