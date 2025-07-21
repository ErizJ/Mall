package model

type Product struct {
	ProductId   int64   `gorm:"column:product_id;primaryKey" json:"product_id"`
	Name        string  `gorm:"column:name" json:"name"`
	Description string  `gorm:"column:description" json:"description"`
	Price       float64 `gorm:"column:price" json:"price"`
	Stock       int64   `gorm:"column:stock" json:"stock"`
	ImageUrl    string  `gorm:"column:image_url" json:"image_url"`
	CategoryId  int64   `gorm:"column:category_id" json:"category_id"`
	IsAvailable bool    `gorm:"column:is_available" json:"is_available"`
}
