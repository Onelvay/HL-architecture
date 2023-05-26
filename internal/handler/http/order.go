package http

import (
	"fmt"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/Onelvay/HL-architecture/pkg/logger"
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

	userID, err := ctx.Get("x-userId")
	if !err {
		ctx.JSON(401, gin.H{"error": "no user_id"})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.UserId = fmt.Sprintf("%s", userID)
	fmt.Println(req)
	res := o.orderService.Create(ctx, req)
	if res.Error != nil {
		logger.Error(res.Error)
	}
	ctx.JSON(int(res.Status), res)
}

func (o *OrderHandler) GET(ctx *gin.Context) {
	userID, ok := ctx.Get("x-userId")
	if !ok {
		ctx.JSON(400, gin.H{"error": "no user_id"})
		return
	}

	orders, err := o.orderService.GetMany(ctx, fmt.Sprintf("%s", userID))
	if err != nil {
		logger.Error(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)
}
