package http

import (
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/gin-gonic/gin"
)

type authorization struct {
	authService service.AuthService
}

func newAuthorizationHandler(authService service.AuthService) *authorization {
	return &authorization{authService: authService}
}

func authRoutes(router *gin.Engine, a *authorization) {
	group := router.Group("/auth")
	{
		group.GET("/sign-in", a.signIn)
		group.POST("/sign-up", a.signUp)
	}

}

func (a *authorization) signUp(c *gin.Context) {
	var req dto.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := a.authService.SignUp(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

func (a *authorization) signIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := a.authService.SignIn(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
