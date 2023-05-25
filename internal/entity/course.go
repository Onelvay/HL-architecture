package entity

type Course struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ImgURL      string `db:"imgurl"`
}
