package action

import (
	"github.com/hooone/evening/pkg/services/locale"
	"github.com/hooone/evening/pkg/services/parameter"
)

type ViewAction struct {
	Id          int64
	CardId      int64
	Name        string
	Type        string //CRUD
	Text        string
	Seq         int32
	DoubleCheck bool
	Locale      locale.Locale
	Parameters  []*parameter.Parameter
}
type ViewActionSlice []*ViewAction

func (s ViewActionSlice) Len() int           { return len(s) }
func (s ViewActionSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ViewActionSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }
