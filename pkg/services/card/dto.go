package card

import (
	"github.com/hooone/evening/pkg/services/action"
	"github.com/hooone/evening/pkg/services/field"
	"github.com/hooone/evening/pkg/services/locale"
	"github.com/hooone/evening/pkg/services/style"
)

//Card view card is base element in evening
type Card struct {
	Id      int64
	PageId  int64
	OrgId   int64
	Name    string
	Text    string
	Seq     int32
	Pos     int32
	Width   int32
	Style   string
	Locale  locale.Locale
	Fields  []*field.Field
	Actions []*action.ViewAction
	Styles  []*style.Style
	Data    []interface{}
}
type CardSlice []*Card

func (s CardSlice) Len() int           { return len(s) }
func (s CardSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s CardSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }
