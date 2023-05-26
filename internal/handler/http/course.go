package http

import (
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/Onelvay/HL-architecture/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseHandler struct {
	courseService service.CourseService
}

func newCourseHandler(productService service.CourseService) *CourseHandler {
	return &CourseHandler{
		courseService: productService,
	}
}

func courseRoutes(router *gin.Engine, h *CourseHandler) {
	group := router.Group("/courses")
	//group.Use(Middleware)

	{
		group.GET("", h.getAll)

		group.GET("/:id", h.getById)
	}
}

func (h *CourseHandler) getAll(ctx *gin.Context) {
	res, err := h.courseService.GetRowById(ctx, "test")

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ctx.AbortWithError(http.StatusNotFound, err)
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	var test []dto.CourseResponse
	test = append(test, res)
	test = append(test, res)
	test = append(test, res)
	test = append(test, res)
	ctx.JSON(http.StatusOK, test)
}

func (h *CourseHandler) getById(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.courseService.GetRowById(ctx, id)

	if err != nil {
		logger.Error(err)
		if err.Error() == "sql: no rows in result set" {
			ctx.AbortWithError(http.StatusNotFound, err)
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	ctx.JSON(http.StatusOK, res)
}
