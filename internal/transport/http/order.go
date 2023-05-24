package http

import (
	"fmt"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
}

func newOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
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
	var req dto.OrderRequest
	fmt.Println(ctx.Get("x-testing-abay"))
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res := o.orderService.Create(ctx, req)
	ctx.JSON(int(res.Status), res)
}
func (o *OrderHandler) GET(ctx *gin.Context) {

}
