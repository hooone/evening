package models

//Field 数据字段
type Field struct {
	Id        int64
	CardId    int64
	Name      string
	Seq       int32
	IsVisible bool
	Type      string
	Filter    string
	Default   string
	Locale    Locale `xorm:"-"`
}

type FieldSlice []*Field

func (s FieldSlice) Len() int           { return len(s) }
func (s FieldSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s FieldSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }

type GetFieldsQuery struct {
	CardId int64
	Result []*Field
}
type GetFieldQuery struct {
	FieldId int64
	Result  *Field
}
type CreateFieldCommand struct {
	CardId    int64
	Name      string
	Seq       int32
	IsVisible bool
	Type      string
	Filter    string
	Default   string
	Result    int64
}
type UpdateFieldCommand struct {
	FieldId   int64
	Name      string
	IsVisible bool
	Type      string
	Filter    string
	Default   string
	Result    *Field
}
type UpdateFieldSeqCommand struct {
	FieldId int64
	Seq     int32
	Result  *Field
}
type DeleteFieldCommand struct {
	FieldId int64
	Result  *Field
}
