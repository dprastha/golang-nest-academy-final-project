package main

import (
	"final-project/db"
	"final-project/server"

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

	app.Start(":8080")
}
