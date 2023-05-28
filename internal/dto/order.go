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

	OrderReviewRequest struct {
		OrderId string `json:"order_id"`
		Rating  int8   `json:"rating"`
		Comment string `json:"comment"`
	}
)
