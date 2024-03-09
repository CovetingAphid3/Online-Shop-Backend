package main

import (
	"gen-shop/api"
	"gen-shop/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	api.Routes(r)
	r.Run()
}
