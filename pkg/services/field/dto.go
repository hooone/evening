package field

import "github.com/hooone/evening/pkg/services/locale"

type Field struct {
	Id        int64
	CardId    int64
	Name      string
	Text      string
	Seq       int32
	IsVisible bool
	Type      string
	Filter    string
	Default   string
	Locale    locale.Locale
}

type FieldSlice []*Field

func (s FieldSlice) Len() int           { return len(s) }
func (s FieldSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s FieldSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }
