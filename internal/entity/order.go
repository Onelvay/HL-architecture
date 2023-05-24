package entity

type Order struct {
	ID       string `db:"order_id"`
	UserId   string `db:"user_id"`
	CourseId string `db:"course_id"`
}
