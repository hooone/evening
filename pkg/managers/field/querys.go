package field

type GetFieldsQuery struct {
	CardId int64
	OrgId  int64
	Result []*Field
}
type GetFieldQuery struct {
	FieldId int64
	OrgId   int64
	Result  *Field
}
type CreateFieldCommand struct {
	OrgId     int64
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
	OrgId     int64
	FieldId   int64
	Name      string
	IsVisible bool
	Type      string
	Filter    string
	Default   string
	Result    *Field
}
type UpdateFieldSeqCommand struct {
	OrgId   int64
	FieldId int64
	Seq     int32
	Result  *Field
}
type DeleteFieldCommand struct {
	OrgId   int64
	FieldId int64
	Result  *Field
}
