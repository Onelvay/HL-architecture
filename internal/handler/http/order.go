package http

import (
	"fmt"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/Onelvay/HL-architecture/pkg/logger"
	"github.com/Onelvay/HL-architecture/pkg/server/status"
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

// @Summary create order
// @Description create order by course id and user id
// @Security ApiKeyAuth
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		request	body dto.OrderRequest true	"order req"
// @Success	200	{object}	dto.OrderResponse
// @Failure	401	{object}	status.Response
// @Failure	400	{object}	status.Response
// @Router		/order [post]
func (o *OrderHandler) create(ctx *gin.Context) {
	var req dto.OrderRequest

	userID, err := ctx.Get("x-userId")
	if !err {
		status.NewResponse(ctx, 401, "not registered")
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		status.NewResponse(ctx, 400, err.Error())
		return
	}

	req.UserId = fmt.Sprintf("%s", userID)

	res := o.orderService.CreateRow(ctx, req)
	if res.Error != nil {
		logger.Error(res.Error)
	}
	ctx.JSON(int(res.Status), res)
}

// @Summary get user's orders
// @Security ApiKeyAuth
// @Tags		order
// @Accept		json
// @Produce		json
// @Success	200	{array}	entity.Order
// @Failure	500	{object}	status.Response
// @Failure	401	{object}	status.Response
// @Router		/order [get]
func (o *OrderHandler) get(ctx *gin.Context) {
	userID, ok := ctx.Get("x-userId")
	if !ok {
		status.NewResponse(ctx, 401, "not registered")
		return
	}

	orders, err := o.orderService.GetMany(ctx, fmt.Sprintf("%s", userID))
	if err != nil {
		status.NewResponse(ctx, 500, err.Error())
		return
	}

	ctx.JSON(200, orders)
}

// @Summary add review to order
// @Security ApiKeyAuth
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		request	body dto.OrderReviewRequest true	"review req"
// @Success	200
// @Failure	500	{object}	status.Response
// @Failure	401	{object}	status.Response
// @Router		/order/review [post]
func (o *OrderHandler) addReview(ctx *gin.Context) {
	var req dto.OrderReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		status.NewResponse(ctx, 400, err.Error())
		return
	}

	err := o.orderService.AddReview(ctx, req)
	if err != nil {
		status.NewResponse(ctx, 500, err.Error())
		return
	}

	ctx.Status(200)
}

// @Summary get all reviews
// @Security ApiKeyAuth
// @Tags		order
// @Accept		json
// @Produce		json
// @Success	200 {array} dto.ReviewResponse
// @Failure	500	{object}	status.Response
// @Failure	401	{object}	status.Response
// @Router		/order/review [get]
func (o *OrderHandler) getAllReviews(ctx *gin.Context) {
	orders, err := o.orderService.GetAllReviews(ctx)

	if err != nil {
		status.NewResponse(ctx, 500, err.Error())
		return
	}
	ctx.JSON(200, orders)
}

// @Summary delete review by id
// @Security ApiKeyAuth
// @Tags		order
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"path param"
// @Success	200
// @Failure	500	{object}	status.Response
// @Failure	401	{object}	status.Response
// @Router		/order/review [delete]
func (o *OrderHandler) deleteOrder(ctx *gin.Context) {
	var req dto.OrderDeleteRequest
	req.OrderId = ctx.Param("id")

	if req.OrderId == "" {
		status.NewResponse(ctx, 400, "no order_id")
		return
	}

	err := o.orderService.DeleteByOrderId(ctx, req)
	if err != nil {
		status.NewResponse(ctx, 500, err.Error())
		return
	}

	ctx.Status(200)
}
