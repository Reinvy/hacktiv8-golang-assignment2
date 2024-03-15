package models

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"id"`
	OrderID     uint   `json:"-"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
