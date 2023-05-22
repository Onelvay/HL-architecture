package http

import (
	"github.com/Onelvay/HL-architecture/internal/service"
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
	group := router.Group("/products")
	group.Use(Middleware)

	{
		group.GET("", h.getAll)

		group.GET("/:id", h.getById)
	}
}

func (h *CourseHandler) getAll(ctx *gin.Context) {

}

func (h *CourseHandler) getById(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.courseService.GetRowById(ctx, id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			ctx.AbortWithError(http.StatusNotFound, err)
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	ctx.JSON(http.StatusOK, res)
}
