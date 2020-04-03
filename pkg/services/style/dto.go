package style

import "github.com/hooone/evening/pkg/services/field"

type Style struct {
	Id         int64
	CardId     int64
	FieldId    int64
	Type       string
	Property   string
	Value      string
	Relation   bool
	MustNumber bool
	Field      *field.Field
}
