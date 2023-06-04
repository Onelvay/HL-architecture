package v1

import (
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

	{
		group.GET("", h.getAll)

		group.GET("/:id", h.getById)
		//group.GET("/hello", Helloworld)
	}
}

type badResponse struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

// CreateTags godoc
// @Summary Get all courses
// @Description get all courses
// @Produce application/json
// @Tags courses
// @Success 200 {array} dto.CourseResponse
// @Router /courses [get]

func (h *CourseHandler) getAll(ctx *gin.Context) {
	res, err := h.courseService.GetMany(ctx)
	if err != nil {
		logger.Error(err)
		ctx.JSON(500, badResponse{500, err.Error()})
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
// @Failure	500	{object}	error
// @Failure	400	{object}	error
// @Router		/courses/{id} [get]

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
