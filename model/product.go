package model

type Product struct {
	ID          int     `gorm:"primaryKey;autoIncrement" json:"id_product"`
	Name        string  `gorm:"column:name;type:varchar(255)" json:"name"`
	Price       float64 `gorm:"type:decimal(10,2)" json:"price"`
	Description string  `gorm:"column:description;type:varchar(255)" json:"description"`
}
