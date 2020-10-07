package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type IdDutyWorkersData struct {
	ID string `json:"id"`
}

type DutyWorkersData struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	IDCategory string `json:"idcategory"`
}


type PatchDutyWorkersData struct {
	ID string `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	IDCategory string `json:"idcategory"`
}

type DutyWorkers struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FirstName  string
	LastName   string
	Email      string
	IDCategory primitive.ObjectID
}

func NewDutyWorkers() *DutyWorkers {
	return &DutyWorkers{}
}

func ConvertDutyWorkers(data DutyWorkersData) *DutyWorkers {
	newWorker := NewDutyWorkers()
	newWorker.FirstName = data.FirstName
	newWorker.LastName = data.LastName
	newWorker.Email = data.Email
	newWorker.IDCategory, _ = primitive.ObjectIDFromHex(data.IDCategory)
	return newWorker
}
