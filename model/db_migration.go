package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserId    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	Name      string
	Email     string `gorm:"unique;not null"`
	Password  string
	IsAdmin   bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:current_timestamp;"`
	UpdateAt  time.Time `gorm:"type:timestamptz;"`
}

type ProductCategory struct {
	ProductCategoryId uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	Name              string
	CreatedAt         time.Time `gorm:"type:timestamptz;default:current_timestamp;"`
	UpdateAt          time.Time `gorm:"type:timestamptz;"`
}

type Product struct {
	ProductId         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	Name              string
	ImageUrl          string `gorm:"type:text" json:"imageUrl"`
	Price             float64
	ProductCategoryId uuid.UUID `gorm:"type:uuid"`
	CreatedAt         time.Time `gorm:"type:timestamptz;default:current_timestamp;"`
	UpdateAt          time.Time `gorm:"type:timestamptz;"`
}

type ProductDetail struct {
	ProductDetailId uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	ProductId       uuid.UUID `gorm:"type:uuid"`
	Description     string    `json:"description" gorm:"type:text"`
	CreatedAt       time.Time `gorm:"type:timestamptz;default:current_timestamp;"`
	UpdateAt        time.Time `gorm:"type:timestamptz;"`
}

type Cart struct {
	CartId    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	UserId    uuid.UUID `gorm:"type:uuid"`
	ProductId uuid.UUID `gorm:"type:uuid"`
	Quantity  int
	CreatedAt time.Time `gorm:"type:timestamptz;default:current_timestamp;"`
	UpdateAt  time.Time `gorm:"type:timestamptz;"`
}
