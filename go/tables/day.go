package tables

type Day struct {
	IDDay     int `json:"idday"`
	DayNumber int `json:"daynumber"`
	IsDayoff  int `json:"isdayoff"`
}

func NewDay() *Day {
	return &Day{}
}
