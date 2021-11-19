package handler

import (
	entities "products-api/entities/product"
	"products-api/helper"
	"products-api/layer/products"

	"github.com/gin-gonic/gin"
)

type produkHandler struct {
	service products.Service
}

func NewProductsHandler(service products.Service) *produkHandler {
	return &produkHandler{service}
}

func (h *produkHandler) GetAllProducts(c *gin.Context) {
	sortBy, _ := c.GetQuery("sortBy")
	sortType, _ := c.GetQuery("sortType")
	if sortType != "asc" {
		sortType = "desc"
	}
	if sortBy == "" {
		sortBy = "created_at"
	}

	query := sortBy + " " + sortType
	products, err := h.service.GetAllProducts(query)

	if err != nil {
		res := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, res)
		return
	}

	res := helper.APIResponse(200, "success", products)
	c.JSON(200, res)
}

func (h *produkHandler) CreateProductHandler(c *gin.Context) {

	var inputProduk entities.ProductInput

	if err := c.ShouldBindJSON(&inputProduk); err != nil {
		responseErr := helper.APIFailure(400, "input data required", gin.H{"errors": err})
		c.JSON(400, responseErr)
		return
	}

	newProduct, err := h.service.CreateProduct(inputProduk)

	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}
	response := helper.APIResponse(201, "success create new products", newProduct)
	c.JSON(201, response)
}
