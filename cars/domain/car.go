package domain

type Car struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"unique"`
	Company string `json:"company"`
	Price   int    `json:"price"`
}
