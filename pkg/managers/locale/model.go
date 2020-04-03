package locale

import "time"

type LocaleSet struct {
	Id       int64
	VId      int64
	Name     string
	Type     string
	Language string
	Text     string
	OrgId    int64
	CreateAt time.Time
	UpdateAt time.Time
}

func (l LocaleSet) TableName() string {
	return "locale"
}
