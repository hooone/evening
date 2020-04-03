package field

import "time"

type Field struct {
	Id        int64
	CardId    int64
	Name      string
	Seq       int32
	IsVisible bool
	Type      string
	Filter    string
	Default   string
	OrgId     int64
	CreateAt  time.Time
	UpdateAt  time.Time
}
