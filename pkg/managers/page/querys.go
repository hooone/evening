package page

type GetPagesQuery struct {
	FolderId int64
	OrgId    int64
	Result   []*Page
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
	OrgId    int64
	Name     string
	Seq      int32
	FolderId int64
	Result   int64
}

type UpdatePageCommand struct {
	OrgId  int64
	PageId int64
	Name   string
	Result *Page
}

type UpdatePageSeqCommand struct {
	OrgId    int64
	PageId   int64
	FolderId int64
	Seq      int32
	Result   int64
}

type DeletePageCommand struct {
	OrgId  int64
	PageId int64
	Result *Page
}
type DeletePageByFolderCommand struct {
	OrgId    int64
	FolderId int64
	Result   int64
}
