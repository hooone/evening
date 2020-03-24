package models

type Style struct {
	Id         int64
	CardId     int64
	FieldId    int64
	Type       string
	Property   string
	Value      string
	Relation   bool
	MustNumber bool
	Field      *Field `xorm:"-"`
}
type StyleSet struct {
	Id         int64
	Type       string
	Property   string
	Value      string
	MustNumber bool
	Relation   bool
}

type GetStylesQuery struct {
	CardId int64
	Style  string
	Result []*Style
}
type UpdateStyleCommand struct {
	Id      int64
	FieldId int64
	Value   string
	Result  *Style
}
