package parameter

import "time"

type Parameter struct {
	Id         int64
	ActionId   int64
	FieldId    int64
	IsVisible  bool
	IsEditable bool
	Default    string
	OrgId      int64
	CreateAt   time.Time
	UpdateAt   time.Time
}
