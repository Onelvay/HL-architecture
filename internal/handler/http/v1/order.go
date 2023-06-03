package v1

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
		group.POST("", order.create)
		group.GET("", order.get)
		group.POST("/review", order.addReview)
		group.GET("/review", order.getAllReviews)
		group.DELETE("/:id", order.deleteOrder)
	}
}
func (o *OrderHandler) create(ctx *gin.Context) {
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
	res := o.orderService.CreateRow(ctx, req)
	if res.Error != nil {
		logger.Error(res.Error)
	}
	ctx.JSON(int(res.Status), res)
}

func (o *OrderHandler) get(ctx *gin.Context) {
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

func (o *OrderHandler) addReview(ctx *gin.Context) {
	var req dto.OrderReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := o.orderService.AddReview(ctx, req)
	if err != nil {
		logger.Error(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}
func (o *OrderHandler) getAllReviews(ctx *gin.Context) {
	orders, err := o.orderService.GetAllReviews(ctx)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, orders)
}

func (o *OrderHandler) deleteOrder(ctx *gin.Context) {
	var req dto.OrderDeleteRequest
	req.OrderId = ctx.Param("id")

	if req.OrderId == "" {
		ctx.JSON(400, gin.H{"error": "no order_id"})
		return
	}

	err := o.orderService.DeleteByOrderId(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}