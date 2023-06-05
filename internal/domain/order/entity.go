package order

type Order struct {
	ID       string `db:"order_id"`
	UserId   string `db:"user_id"`
	CourseId string `db:"course_id"`
}

type OrderReview struct {
	OrderID string `db:"order_id"`
	Comment string `db:"comment"`
	Rating  int8   `db:"rating"`
}
