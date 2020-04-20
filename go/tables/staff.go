package tables

type DbStaff struct {
	IddbStaff     int    `json:"iddbstaff"`
	DbSurname     string `json:"dbsurname"`
	DbName        string `json:"dbname"`
	DbStaffKindID int    `json:"dbstaffkindid"`
}

func NewDbStaff() *DbStaff {
	return &DbStaff{}
}
