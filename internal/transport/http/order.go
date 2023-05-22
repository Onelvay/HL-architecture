package http

import "github.com/gin-gonic/gin"

type OrderHandler struct {
}

func newOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func orderRoutes(router *gin.Engine, order *OrderHandler) {
	group := router.Group("/order")
	group.Use(Middleware)
	{
		group.POST("", order.POST)
		group.GET("", order.GET)
	}
}
func (o *OrderHandler) POST(ctx *gin.Context) {

}
func (o *OrderHandler) GET(ctx *gin.Context) {

}
