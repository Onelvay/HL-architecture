package http

import (
	"github.com/Onelvay/HL-architecture/internal/service"
	"github.com/Onelvay/HL-architecture/pkg/server/status"
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

	{
		group.GET("", h.getAll)

		group.GET("/:id", h.getById)
	}
}

// CreateTags godoc
// @Summary Get all courses
// @Description get all courses
// @Produce application/json
// @Tags courses
// @Success 200 {array} dto.CourseResponse
// @Failure 500 {object} status.Response
// @Router /courses [get]
func (h *CourseHandler) getAll(ctx *gin.Context) {
	res, err := h.courseService.GetMany(ctx)
	if err != nil {
		status.NewResponse(ctx, 500, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary get
// @Tags		courses
// @Accept		json
// @Produce		json
// @Param		id	path		string	true	"path param"
// @Success	200	{object}	dto.CourseResponse
// @Failure	500	{object}	status.Response
// @Failure	400	{object}	status.Response
// @Router		/courses/{id} [get]
func (h *CourseHandler) getById(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.courseService.GetRowById(ctx, id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			status.NewResponse(ctx, 404, err.Error())
		} else {
			status.NewResponse(ctx, 500, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, res)
}
