package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type DutiesData struct {
	IDDutyDay    string `json:"iddutyday"`
	IDDutyWorker string `json:"iddutyworker"`
}

type Duties struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	IDDutyDay    primitive.ObjectID
	IDDutyWorker primitive.ObjectID
}

func NewDuties() *Duties {
	return &Duties{}
}

func ConvertDuties(data DutiesData) *Duties {
	newDuty := NewDuties()
	newDuty.IDDutyDay, _ = primitive.ObjectIDFromHex(data.IDDutyDay)
	newDuty.IDDutyWorker, _ = primitive.ObjectIDFromHex(data.IDDutyWorker)
	return newDuty
}
