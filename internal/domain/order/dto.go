package order

import "time"

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

	ReviewResponse struct {
		Username   string    `json:"user_name" db:"user_name"`
		CourseName string    `json:"course_name" db:"course_name"`
		Rating     int       `json:"rating" db:"rating"`
		Comment    string    `json:"comment" db:"comment"`
		CreatedAt  time.Time `json:"created_at" db:"created_at"`
	}

	OrderDeleteRequest struct {
		UserId  string
		OrderId string `json:"order_id"`
	}

	OrderDeleteRequestByUser struct {
		UserId   string
		CourseId string
	}
)
