package dtos

type CreateFolderForm struct {
	Name string `form:"Name"`
	Text string `form:"Text"`
}

type UpdateFolderForm struct {
	Name     string `form:"Name"`
	Text     string `form:"Text"`
	FolderId int64  `form:"FolderId"`
}

type DeleteFolderForm struct {
	Id int64 `form:"FolderId"`
}

type CreateTreePageForm struct {
	Name string `form:"Name"`
	Text string `form:"Text"`
}
type CreateNodePageForm struct {
	Name     string `form:"Name"`
	Text     string `form:"Text"`
	FolderId int64  `form:"FolderId"`
}

type UpdatePageForm struct {
	Name   string `form:"Name"`
	Text   string `form:"Text"`
	PageID int64  `form:"PageID"`
}
type DeletePageForm struct {
	PageID int64 `form:"PageID"`
}

type NavMoveForm struct {
	SourceFolder int64 `form:"SourceFolder"`
	SourcePage   int64 `form:"SourcePage"`
	TargetFolder int64 `form:"TargetFolder"`
	TargetPage   int64 `form:"TargetPage"`
	Position     int   `form:"Position"`
}
