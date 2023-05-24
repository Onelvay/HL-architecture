package dto

type (
	OrderRequest struct {
		CourseId string `json:"course_id"`
		UserId   string
	}

	OrderResponse struct {
		Status int16 `json:"status"`
		Error  error `json:"error"`
	}
)
