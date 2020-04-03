package parameter

import "github.com/hooone/evening/pkg/services/field"

type Parameter struct {
	Id         int64
	ActionId   int64
	FieldId    int64
	IsVisible  bool
	IsEditable bool
	Default    string
	Compare    string //use for read. eg. eq,ne,le,ge,lt,gt
	Field      *field.Field
	Value      interface{}
}
type ParameterSlice []*Parameter

func (s ParameterSlice) Len() int           { return len(s) }
func (s ParameterSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ParameterSlice) Less(i, j int) bool { return s[i].Field.Seq < s[j].Field.Seq }
