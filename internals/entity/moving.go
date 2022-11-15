package entity

type Movie struct {
	ID        int    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	MovieName string `gorm:"column:movieName"`
	Date      string `gorm:"column:date"`
	Time      string `gorm:"column:time"`
	CinemaNo  string `gorm:"column:cinemaNo"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
	CreatedBy string `gorm:"column:created_by"`
	UpdatedBy string `gorm:"column:updated_by"`
}

func (Movie) TableName() string {
	return "movies"
}
