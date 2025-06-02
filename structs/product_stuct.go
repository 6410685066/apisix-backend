package structs

import "time"

type ProductsEntity struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Category  string    `gorm:"not null"`
	Price     float64   `gorm:"not null default:0;check:price >= 0"`
	Stock     uint      `gorm:"not null default:0;check:price >= 0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CreateBy  string    `gorm:"not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UpdateBy  string    `gorm:"not null"`
}

func (ProductsEntity) TableName() string {
	return "products"
}

type ProductDto struct {
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     float64   `json:"price"`
	Stock     uint      `json:"stock"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductGetRequest struct {
	Name     string  `query:"name"`
	Category string  `query:"category"`
	Price    float64 `query:"price"`
	Stock    uint    `query:"stock"`
}

type ProductPostRequest struct {
	Name     string  `json:"name" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Stock    uint    `json:"stock" binding:"required,gt=0"`
	CreateBy string  `json:"create_by" binding:"required"`
	UpdateBy string  `json:"update_by" binding:"required"`
}

type ProductPutRequest struct {
	ID       uint    `json:"id" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Stock    uint    `json:"stock" binding:"required,gt=0"`
	CreateBy string  `json:"create_by" binding:"required"`
	UpdateBy string  `json:"update_by" binding:"required"`
}

type ProductPatchRequest struct {
	ID       uint     `json:"id" binding:"required"`
	Name     *string  `json:"name"`
	Category *string  `json:"category"`
	Price    *float64 `json:"price"`
	Stock    *uint    `json:"stock"`
	CreateBy string   `json:"create_by" binding:"required"`
	UpdateBy string   `json:"update_by" binding:"required"`
}
