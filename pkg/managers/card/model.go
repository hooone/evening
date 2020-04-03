package card

import "time"

//Card view card is base element in evening
type Card struct {
	Id       int64
	PageId   int64
	OrgId    int64
	Name     string
	Seq      int32
	Pos      int32
	Width    int32
	Style    string
	CreateAt time.Time
	UpdateAt time.Time
}
