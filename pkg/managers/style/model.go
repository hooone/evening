package style

import "time"

type Style struct {
	Id         int64
	CardId     int64
	FieldId    int64
	Type       string
	Property   string
	Value      string
	Relation   bool
	MustNumber bool
	OrgId      int64
	CreateAt   time.Time
	UpdateAt   time.Time
}
type StyleSet struct {
	Id         int64
	Type       string
	Property   string
	Value      string
	MustNumber bool
	Relation   bool
}
