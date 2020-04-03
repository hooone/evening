package folder

import (
	"errors"
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetFolder)
	bus.AddHandler("sql", GetFolderByID)
	bus.AddHandler("sql", CreateFolder)
	bus.AddHandler("sql", UpdateFolder)
	bus.AddHandler("sql", UpdateFolderSeq)
	bus.AddHandler("sql", DeleteFolder)
}

var (
	ErrFolderNotFound = errors.New("folder data not found")
)

//GetFolder 获得文件夹列表
func GetFolder(query *GetFolderQuery) error {
	folders := make([]*Folder, 0)
	err := sqlstore.X.Where("org_id = ?", query.OrgId).Find(&folders)
	if err != nil {
		return err
	}
	query.Result = folders
	return nil
}

//GetFolderByID 获得文件夹
func GetFolderByID(query *GetFolderByIDQuery) error {
	folder := Folder{
		Id:    query.FolderId,
		OrgId: query.OrgId,
	}
	success, err := sqlstore.X.Where("id=? and org_id = ?", query.FolderId, query.OrgId).Get(&folder)
	if err != nil {
		return err
	}
	if !success {
		return ErrFolderNotFound
	}
	query.Result = &folder
	return nil
}

//CreateFolder 添加文件夹
func CreateFolder(cmd *CreateFolderCommand) error {
	folder := Folder{
		Name:     cmd.Name,
		Seq:      cmd.Seq,
		IsFolder: cmd.IsFolder,
		OrgId:    cmd.OrgId,
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
	}
	_, err := sqlstore.X.Insert(&folder)
	if err != nil {
		return err
	}
	cmd.Result = folder.Id
	return nil
}

//UpdateFolder 修改文件夹信息
func UpdateFolder(cmd *UpdateFolderCommand) error {
	folder := Folder{
		Id:       cmd.FolderId,
		OrgId:    cmd.OrgId,
		Name:     cmd.Name,
		UpdateAt: time.Now(),
	}
	count, err := sqlstore.X.Where("id=? and org_id = ?", cmd.FolderId, cmd.OrgId).Cols("name").Cols("update_at").Update(&folder)
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}

//UpdateFolderSeq 修改文件夹信息
func UpdateFolderSeq(cmd *UpdateFolderSeqCommand) error {
	folder := Folder{
		Id:       cmd.FolderId,
		OrgId:    cmd.OrgId,
		Seq:      cmd.Seq,
		UpdateAt: time.Now(),
	}
	count, err := sqlstore.X.Where("id=? and org_id = ?", cmd.FolderId, cmd.OrgId).Cols("seq").Cols("update_at").Update(&folder)
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}

//DeleteFolder 删除文件夹
func DeleteFolder(cmd *DeleteFolderCommand) error {
	folder := Folder{
		Id:    cmd.Id,
		OrgId: cmd.OrgId,
	}
	len, err := sqlstore.X.Where("id=? and org_id = ?", cmd.Id, cmd.OrgId).Delete(&folder)
	if err != nil {
		return err
	}
	cmd.Result = len
	return nil
}
