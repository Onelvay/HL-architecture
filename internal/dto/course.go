package dto

import (
	"github.com/Onelvay/HL-architecture/internal/entity"
)

type CourseRequest struct {
	ID    string  `json:"id"`
	Name  string  `json:"fullName" validate:"required"`
	Price float32 `json:"pseudonym" validate:"required"`
}

type CourseResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"fullName"`
	Price float32 `json:"pseudonym"`
}

func ParseFromCourse(data entity.Course) (res CourseResponse) {
	res = CourseResponse{
		ID:   data.ID,
		Name: data.Name,
	}
	return
}
