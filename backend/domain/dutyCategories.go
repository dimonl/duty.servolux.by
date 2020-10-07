package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type DutyCategoriesData struct {
	Category string `json:"category"`
}

type IdDutyCategoriesData struct {
	ID string `json:"id"`
}

type PatchDutyCategoriesData struct {
	ID   string `json:"id"`
	Category string `json:"category"`
}

type DutyCategories struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Category string
}

func NewDutyCategories() *DutyCategories {
	return &DutyCategories{}
}

func ConvertDutyCategories(data DutyCategoriesData) *DutyCategories {
	newCategory := NewDutyCategories()
	newCategory.Category = data.Category
	return newCategory
}
