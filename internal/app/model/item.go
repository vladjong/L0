package model

type Item struct {
	ChrtId      int    `db:"chrt_id"`
	TrackNumber string `db:"track_number"`
	Price       int    `db:"price"`
	Rid         string `db:"rid"`
	Name        string `db:"name"`
	Sale        int    `db:"sale"`
	Size        string `db:"size"`
	TotalPrice  int    `db:"total_price"`
	NmId        int    `db:"nm_id"`
	Brand       string `db:"brand"`
	Status      int    `db:"status"`
}
