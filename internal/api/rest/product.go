package rest

import (
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func productRoutes(router *gin.Engine, h *ProductHandler) {
	group := router.Group("/products")
	{
		group.GET("", h.getAll)

		group.GET("/:id", h.getById)
	}

}

func (h *ProductHandler) getAll(ctx *gin.Context) {

}
func (h *ProductHandler) getById(ctx *gin.Context) {

}
