package tables

type DbStaffKind struct {
	IddbStaffKind   int    `json:"iddbstaffkind"`
	DbStaffKindName string `json:"dbstaffkindname"`
}

func NewDbStaffKind() *DbStaffKind {
	return &DbStaffKind{}
}
