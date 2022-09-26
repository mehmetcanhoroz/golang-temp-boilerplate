package models

const CategoryModelTableName = "categories"
const CategoryModelItemsForeignObjectName = "Items"
const CategoryModelSubCategoriesForeignObjectName = "SubCategories"

type Category struct {
	ID               uint64      `json:"id" gorm:"primaryKey"`
	Name             string      `json:"name"`
	Slug             string      `json:"slug" gorm:"unique;not null"`
	Items            *[]Item     `json:"items,omitempty"`
	ParentCategoryID *uint64     `json:"-"`
	SubCategories    *[]Category `json:"sub_categories,omitempty" gorm:"foreignKey:ParentCategoryID"`
}

func (Category) TableName() string {
	return CategoryModelTableName
}
