package model

type Product struct {
	ID    int     `gorm:"primaryKey;autoIncrement" json:"id_product"`
	Name  string  `gorm:"column:productname;type:varchar(255)" json:"name"`
	Price float64 `gorm:"type:decimal(10,2)" json:"price"`
}
