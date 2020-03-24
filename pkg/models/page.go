package models

type Page struct {
	Id       int64
	FolderId int64
	OrgId    int64
	Name     string
	Seq      int32
	Locale   Locale `xorm:"-"`
}

type PageSlice []*Page

func (s PageSlice) Len() int           { return len(s) }
func (s PageSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PageSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }

type GetPagesQuery struct {
	FolderId int64
	OrgId    int64
	Result   PageSlice
}

type GetPageByIDQuery struct {
	PageId int64
	OrgId  int64
	Result *Page
}
type GetPageByNameQuery struct {
	PageName string
	OrgId    int64
	Result   *Page
}

type CreatePageCommand struct {
	Name     string
	Seq      int32
	FolderId int64
	OrgId    int64
	Result   int64
}

type UpdatePageCommand struct {
	PageId int64
	Name   string
	Result *Page
}

type UpdatePageSeqCommand struct {
	PageId   int64
	FolderId int64
	Seq      int32
	Result   int64
}

type DeletePageCommand struct {
	PageId int64
	Result *Page
}
type DeletePageByFolderCommand struct {
	FolderId int64
	Result   int64
}
