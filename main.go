package main

import (
	"products-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.ProductsRoute(r)

	r.Run()
}
