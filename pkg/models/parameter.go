package models

//Parameter  is base element in evening
type Parameter struct {
	Id         int64
	ActionId   int64
	FieldId    int64
	IsVisible  bool
	IsEditable bool
	Default    string
	Compare    string      `xorm:"-"` //use for read. eg. eq,ne,le,ge,lt,gt
	Field      *Field      `xorm:"-"`
	Value      interface{} `xorm:"-"`
}
type ParameterSlice []*Parameter

func (s ParameterSlice) Len() int           { return len(s) }
func (s ParameterSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ParameterSlice) Less(i, j int) bool { return s[i].Field.Seq < s[j].Field.Seq }

type GetParameterQuery struct {
	ParameterId int64
	Result      *Parameter
}
type GetParametersQuery struct {
	ActionId int64
	Result   []*Parameter
}
type CreateParameterCommand struct {
	ActionId   int64
	FieldId    int64
	IsVisible  bool
	IsEditable bool
	Default    string
	Result     int64
}
type UpdateParameterCommand struct {
	ParameterId int64
	IsVisible   bool
	IsEditable  bool
	Default     string
	Result      *Parameter
}

type DeleteParameterByFieldCommand struct {
	FieldId int64
	Result  int64
}

type DeleteParameterByActionCommand struct {
	ActionId int64
	Result   int64
}
