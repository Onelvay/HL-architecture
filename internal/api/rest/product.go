package rest

import (
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func productRoutes(router *gin.Engine) {
	h := ProductHandler{}
	group := router.Group("/products")
	{
		group.GET("", h.getAll)

		group.GET("/id", h.getById)
	}

}

func (h *ProductHandler) getAll(ctx *gin.Context) {

}
func (h *ProductHandler) getById(ctx *gin.Context) {

}
