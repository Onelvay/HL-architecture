package rest

import (
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
)

type Authorization struct {
	authService service.AuthService
}

func NewAuthorizationHandler() *Authorization {
	return &Authorization{}
}

func authRoutes(router *gin.Engine, a *Authorization) {
	group := router.Group("/auth")
	{
		group.GET("/sign-in", a.SignIn)
		group.POST("/sign-up", a.SignUp)
	}

}

func (a *Authorization) SignUp(c *gin.Context) {
	var req dto.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := a.authService.SignUp(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, res)
}

func (a *Authorization) SignIn(ctx *gin.Context) {

}
