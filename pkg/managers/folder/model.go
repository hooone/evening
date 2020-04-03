package folder

import "time"

//Folder view folder
type Folder struct {
	Id       int64
	OrgId    int64
	Name     string
	IsFolder bool
	Seq      int32
	CreateAt time.Time
	UpdateAt time.Time
}
