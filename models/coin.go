package models

type Coin struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Coin    string `json:"coin"`
	Balance int    `json:"balance"`
}
