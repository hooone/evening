package navigation

import (
	"errors"

	"github.com/hooone/evening/pkg/services/card"
	"github.com/hooone/evening/pkg/services/locale"
)

var (
	ErrPageNameConflict = errors.New("Page name aleady exist")
)

//Folder Data Transfer Object
type Folder struct {
	Id       int64
	OrgId    int64
	Name     string
	Text     string
	IsFolder bool
	Seq      int32
	Locale   locale.Locale
	Pages    []*Page
}
type Page struct {
	Id       int64
	FolderId int64
	OrgId    int64
	Name     string
	Text     string
	Seq      int32
	Cards    []*card.Card
	Locale   locale.Locale
}

type FolderSlice []*Folder

func (s FolderSlice) Len() int           { return len(s) }
func (s FolderSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s FolderSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }

type PageSlice []*Page

func (s PageSlice) Len() int           { return len(s) }
func (s PageSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PageSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }

type NavMoveForm struct {
	SourceFolder int64 `form:"SourceFolder"`
	SourcePage   int64 `form:"SourcePage"`
	TargetFolder int64 `form:"TargetFolder"`
	TargetPage   int64 `form:"TargetPage"`
	Position     int   `form:"Position"`
}
