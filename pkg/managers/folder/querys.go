package folder

type GetFolderQuery struct {
	OrgId  int64
	Result []*Folder
}
type GetFolderByIDQuery struct {
	OrgId    int64
	FolderId int64
	Result   *Folder
}

type CreateFolderCommand struct {
	Name     string
	Seq      int32
	IsFolder bool
	OrgId    int64
	Result   int64
}
type UpdateFolderCommand struct {
	OrgId    int64
	FolderId int64
	Name     string
	Result   int64
}

type UpdateFolderSeqCommand struct {
	OrgId    int64
	FolderId int64
	Seq      int32
	Result   int64
}

type DeleteFolderCommand struct {
	OrgId  int64
	Id     int64
	Result int64
}
