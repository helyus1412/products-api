package routes

import (
	"products-api/config"
	"products-api/handler"
	"products-api/layer/products"

	"github.com/gin-gonic/gin"
)

var (
	DB              = config.Connection()
	productsRepo    = products.NewRepository(DB)
	productsService = products.NewService(productsRepo)
	productsHandler = handler.NewProductsHandler(productsService)
)

func ProductsRoute(r *gin.Engine) {
	r.GET("/product", productsHandler.GetAllProducts)
	r.POST("/product", productsHandler.CreateProductHandler)
}
