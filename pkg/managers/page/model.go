package page

import "time"

type Page struct {
	Id       int64
	FolderId int64
	OrgId    int64
	Name     string
	Seq      int32
	CreateAt time.Time
	UpdateAt time.Time
}
