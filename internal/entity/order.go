package entity

type Order struct {
	ID       string `db:"id"`
	UserId   string `db:"user_id"`
	CourseId string `db:"course_id"`
}
