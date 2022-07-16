package model

type Delivery struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Phone   string `db:"phone"`
	Zip     string `db:"zip"`
	City    string `db:"city"`
	Address string `db:"address"`
	Region  string `db:"region"`
	Email   string `db:"email"`
}
