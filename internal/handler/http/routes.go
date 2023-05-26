package http

import (
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, s service.Service) {
	c := newCourseHandler(s.Course)
	courseRoutes(router, c)

	a := newAuthorizationHandler(s.Auth)
	authRoutes(router, a)

	o := newOrderHandler(s.Order)
	orderRoutes(router, o)

	u := newUserHandler(s.User)
	userRoutes(router, u)
}
