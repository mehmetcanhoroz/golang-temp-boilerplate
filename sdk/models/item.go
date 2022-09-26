package models

const ItemModelTableName = "items"

type Item struct {
	ID         uint64    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug" gorm:"unique;not null"`
	Price      uint32    `json:"price,omitempty"`
	CategoryID *uint64   `json:"-"`
	Category   *Category `json:"category,omitempty"`
}

func (Item) TableName() string {
	return ItemModelTableName
}
