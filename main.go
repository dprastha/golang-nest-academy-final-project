package main

import (
	"final-project/config"
	"final-project/db"
	"final-project/server"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO assign db value to repository
	_, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	app := server.NewRouter(router)
	port := fmt.Sprintf(":%s", config.GetEnvVariable("APP_PORT"))

	app.Start(port)
}
