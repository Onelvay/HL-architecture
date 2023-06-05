package course

type CourseRequest struct {
	ID          string  `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description float32 `json:"description" validate:"required"`
	ImgURL      string  `json:"imgURL"`
}

type (
	CourseResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImgURL      string `json:"imgURL"`
	}

	UserCourseResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImgURL      string `json:"imgURL"`
		OrderId     string `json:"order_id"`
	}
)

func ParseFromCourse(data Course) (res CourseResponse) {
	res = CourseResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		ImgURL:      data.ImgURL,
	}
	return
}
