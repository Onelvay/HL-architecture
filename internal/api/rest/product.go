package rest

import (
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	id := ctx.Param("id")
	res, err := h.productService.GetRowById(ctx, id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ctx.AbortWithError(http.StatusNotFound, err)
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	ctx.JSON(http.StatusOK, res)
}
