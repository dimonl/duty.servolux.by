package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type DutyDayData struct {
	Day      string `json:"day"`
	IsDayOff int    `json:"isdayoff"`
}

type IdDutyDayData struct {
	ID string `json:"id"`
}

type PatchDutyDayData struct {
	ID       string `json:"id"`
	Day      string `json:"day"`
	IsDayOff int    `json:"isdayoff"`
}

type DutyDay struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Day      string
	IsDayOff int
}

func NewDutyDay() *DutyDay {
	return &DutyDay{}
}

func ConvertDutyDay(data DutyDayData) *DutyDay {
	newDutyDay := NewDutyDay()
	newDutyDay.Day = data.Day
	newDutyDay.IsDayOff = data.IsDayOff
	return newDutyDay
}
