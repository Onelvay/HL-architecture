package dto

import (
	"github.com/Onelvay/HL-architecture/internal/entity"
)

type ProductRequest struct {
	ID    string  `json:"id"`
	Name  string  `json:"fullName" validate:"required"`
	Price float32 `json:"pseudonym" validate:"required"`
}

type ProductResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"fullName"`
	Price float32 `json:"pseudonym"`
}

func ParseFromProduct(data entity.Product) (res ProductResponse) {
	res = ProductResponse{
		ID:    data.ID,
		Name:  data.Name,
		Price: data.Price,
	}
	return
}
