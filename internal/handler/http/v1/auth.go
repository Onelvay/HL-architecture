package v1

import (
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/Onelvay/HL-architecture/pkg/server/status"
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
		group.POST("/sign-in", a.signIn)
		group.POST("/sign-up", a.signUp)
	}

}

// @Summary 	sign up user
// @Tags		authorization
// @Accept		json
// @Produce		json
//@Param request body 	dto.SignUpRequest true "body param"
// @Success	200	{object}	dto.SignUpResponse
// @Failure	500	{object}	status.Response
// @Failure	400	{object}	status.Response
// @Router		/auth/sign-up [post]

func (a *authorization) signUp(c *gin.Context) {
	var req dto.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		status.NewResponse(c, 400, err.Error())
		return
	}

	res, err := a.authService.SignUp(c, req)
	if err != nil {
		status.NewResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, res)
}

// @Summary 	sign in user
// @Tags		authorization
// @Accept		json
// @Produce		json
//@Param request body 	dto.SignInRequest true "body param"
// @Success	200	{object}	dto.SignIpResponse
// @Failure	500	{object}	status.Response
// @Failure	400	{object}	status.Response
// @Router		/auth/sign-in [post]

func (a *authorization) signIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		status.NewResponse(c, 400, err.Error())
		return
	}

	res, err := a.authService.SignIn(c, req)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			status.NewResponse(c, 400, err.Error())
			return
		}

		status.NewResponse(c, 500, err.Error())
		return
	}
	c.JSON(200, res)
}
