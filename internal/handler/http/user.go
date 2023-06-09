package http

import (
	"errors"
	"fmt"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/Onelvay/HL-architecture/pkg/loggerlocal"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func newUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func userRoutes(router *gin.Engine, user *UserHandler) {
	group := router.Group("/user")
	group.Use(Middleware)

	{
		group.GET("/course", user.getCourses)

	}
}

// @Summary	List of courses from the store
// @Tags		courses
// @Accept		json
// @Produce	json
// @Success	200	{array}		course.Response
// @Failure	500	{object}	status.Response
// @Router		/course  [get]
func (u *UserHandler) getCourses(ctx *gin.Context) {
	userID, ok := ctx.Get("x-userId")

	if !ok {
		loggerlocal.Error(errors.New("no user id"))
		ctx.JSON(400, gin.H{"error": "no user_id"})
		return
	}

	courses, err := u.userService.GetCourses(ctx, fmt.Sprint(userID))
	if err != nil {
		loggerlocal.Error(err)
		ctx.Status(500)
		return
	}

	ctx.JSON(200, courses)
}
