package tables

type Duties struct {
	IDduties     int `json:"idduties"`
	MonthNumber  int `json:"monthnumber"`
	DutiesDay    int `json:"dutiesday"`
	DutiesDtaff  int `json:"dutiesstaff"`
	DutiesDayoff int `json:"dutiesdayoff"`
}

func NewDuties() *Duties {
	return &Duties{}
}
