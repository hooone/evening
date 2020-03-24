package models

//Folder view folder
type Folder struct {
	Id       int64
	OrgId    int64
	Name     string
	IsFolder bool
	Seq      int32
	Locale   Locale  `xorm:"-"`
	Pages    []*Page `xorm:"-"`
}

type FolderSlice []*Folder

func (s FolderSlice) Len() int           { return len(s) }
func (s FolderSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s FolderSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }

type GetNavigationQuery struct {
	OrgId  int64
	Result FolderSlice
}

type GetFolderQuery struct {
	OrgId  int64
	Result FolderSlice
}
type GetFolderByIDQuery struct {
	OrgId    int64
	FolderId int64
	Result   *Folder
}

type CreateFolderCommand struct {
	Name     string
	Seq      int32
	Result   int64
	IsFolder bool
	OrgId    int64
}
type UpdateFolderCommand struct {
	FolderId int64
	Name     string
	Result   int64
}

type UpdateFolderSeqCommand struct {
	FolderId int64
	Seq      int32
	Result   int64
}

type DeleteFolderCommand struct {
	Id     int64
	Result int64
}
