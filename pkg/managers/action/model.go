package action

import "time"

type ViewAction struct {
	Id          int64
	CardId      int64
	Name        string
	Type        string //CRUD
	Seq         int32
	DoubleCheck bool
	OrgId       int64
	CreateAt    time.Time
	UpdateAt    time.Time
}
