package entity

type Product struct {
	ID    string  `db:"id"`
	Name  string  `db:"name"`
	Price float32 `db:"price"`
}
