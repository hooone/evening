package sqlstore

import (
	"errors"
	"sort"
	"strconv"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetNavigation)
	bus.AddHandler("sql", GetFolder)
	bus.AddHandler("sql", GetFolderByID)
	bus.AddHandler("sql", CreateFolder)
	bus.AddHandler("sql", UpdateFolder)
	bus.AddHandler("sql", UpdateFolderSeq)
	bus.AddHandler("sql", DeleteFolder)
	bus.AddHandler("sql", GetPagesByFolder)
	bus.AddHandler("sql", GetPageByID)
	bus.AddHandler("sql", GetPageByName)
	bus.AddHandler("sql", CreatePage)
	bus.AddHandler("sql", UpdatePage)
	bus.AddHandler("sql", UpdatePageSeq)
	bus.AddHandler("sql", DeletePage)
	bus.AddHandler("sql", DeletePageByFolder)
}

//GetNavigation 获得完整导航栏
func GetNavigation(query *models.GetNavigationQuery) error {
	folders := make([]*models.Folder, 0)
	err := x.Where("org_id = ?", query.OrgId).Find(&folders)
	if err != nil {
		return err
	}
	sort.Sort(models.FolderSlice(folders))
	for _, fd := range folders {
		err = x.Where("folder_id = ?", fd.Id).Find(&fd.Pages)
		if err != nil {
			return err
		}
		sort.Sort(models.PageSlice(fd.Pages))
	}
	query.Result = folders
	return nil
}

//GetFolder 获得文件夹列表
func GetFolder(query *models.GetFolderQuery) error {
	folders := make([]*models.Folder, 0)
	err := x.Where("org_id = ?", query.OrgId).Find(&folders)
	if err != nil {
		return err
	}
	sort.Sort(models.FolderSlice(folders))
	query.Result = folders
	return nil
}

//GetFolderByID 获得文件夹
func GetFolderByID(query *models.GetFolderByIDQuery) error {
	folder := models.Folder{
		Id:    query.FolderId,
		OrgId: query.OrgId,
	}
	success, err := x.ID(query.FolderId).Get(&folder)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("no folder found")
	}
	query.Result = &folder
	return nil
}

//CreateFolder 添加文件夹
func CreateFolder(cmd *models.CreateFolderCommand) error {
	folder := models.Folder{
		Name:     cmd.Name,
		Seq:      cmd.Seq,
		IsFolder: cmd.IsFolder,
		OrgId:    cmd.OrgId,
	}
	_, err := x.Insert(&folder)
	if err != nil {
		return err
	}
	cmd.Result = folder.Id
	return nil
}

//UpdateFolder 修改文件夹信息
func UpdateFolder(cmd *models.UpdateFolderCommand) error {
	folder := models.Folder{
		Id:   cmd.FolderId,
		Name: cmd.Name,
	}
	count, err := x.Id(cmd.FolderId).Cols("name").Update(&folder)
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}

//UpdateFolderSeq 修改文件夹信息
func UpdateFolderSeq(cmd *models.UpdateFolderSeqCommand) error {
	folder := models.Folder{
		Id:  cmd.FolderId,
		Seq: cmd.Seq,
	}
	count, err := x.Id(cmd.FolderId).Cols("seq").Update(&folder)
	if err != nil {
		return err
	}

	folders := make([]*models.Folder, 0)
	err = x.Find(&folders)
	if err != nil {
		return err
	}
	sort.Sort(models.FolderSlice(folders))
	for idx, fd := range folders {
		if fd.Seq != int32(idx*2+1) {
			fd.Seq = int32(idx*2 + 1)
			x.Id(fd.Id).Cols("seq").Update(fd)
		}
	}

	cmd.Result = count
	return nil
}

//DeleteFolder 删除文件夹
func DeleteFolder(cmd *models.DeleteFolderCommand) error {
	folder := models.Folder{
		Id: cmd.Id,
	}
	len, err := x.Id(cmd.Id).Delete(&folder)
	if err != nil {
		return err
	}
	cmd.Result = len
	return nil
}

//GetPagesByFolder 获取文件夹里的所有页面
func GetPagesByFolder(query *models.GetPagesQuery) error {
	pages := make([]*models.Page, 0)
	err := x.Where("folder_id = ? and org_id = ?", query.FolderId, query.OrgId).Find(&pages)
	if err != nil {
		return err
	}
	sort.Sort(models.PageSlice(pages))
	query.Result = pages
	return nil
}

//GetPageByID 获取页面
func GetPageByID(query *models.GetPageByIDQuery) error {
	page := models.Page{
		Id:    query.PageId,
		OrgId: query.OrgId,
	}
	success, err := x.ID(query.PageId).Get(&page)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("no page found;ID:" + strconv.Itoa(int(query.PageId)))
	}
	query.Result = &page
	return nil
}

//GetPageByName 获取页面
func GetPageByName(query *models.GetPageByNameQuery) error {
	page := models.Page{}
	success, err := x.Where("name = ? and org_id = ?", query.PageName, query.OrgId).Get(&page)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("no page found;Name:" + query.PageName)
	}
	query.Result = &page
	return nil
}

//CreatePage 添加页面
func CreatePage(cmd *models.CreatePageCommand) error {
	page := models.Page{
		Name:     cmd.Name,
		Seq:      cmd.Seq,
		FolderId: cmd.FolderId,
		OrgId:    cmd.OrgId,
	}
	_, err := x.Insert(&page)
	if err != nil {
		return err
	}
	cmd.Result = page.Id
	return nil
}

//UpdatePage 修改页面信息
func UpdatePage(cmd *models.UpdatePageCommand) error {
	page := models.Page{}
	success, err := x.Id(cmd.PageId).Get(&page)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("no row updated")
	}
	page.Name = cmd.Name
	_, err = x.Id(page.Id).Cols("name").Update(&page)
	if err != nil {
		return err
	}
	cmd.Result = &page
	return nil
}

//UpdatePageSeq 修改页面顺序
func UpdatePageSeq(cmd *models.UpdatePageSeqCommand) error {
	page := models.Page{
		Id:       cmd.PageId,
		FolderId: cmd.FolderId,
		Seq:      cmd.Seq,
	}
	count, err := x.Id(cmd.PageId).Cols("seq").Cols("folder_id").Update(&page)
	if err != nil {
		return err
	}

	pages := make([]*models.Page, 0)
	err = x.Where("folder_id = ?", cmd.FolderId).Find(&pages)
	if err != nil {
		return err
	}
	sort.Sort(models.PageSlice(pages))
	for idx, fd := range pages {
		if fd.Seq != int32(idx*2+1) {
			fd.Seq = int32(idx*2 + 1)
			x.Id(fd.Id).Cols("seq").Update(fd)
		}
	}

	cmd.Result = count
	return nil
}

//DeletePage 删除页面
func DeletePage(cmd *models.DeletePageCommand) error {
	page := models.Page{}
	success, err := x.Id(cmd.PageId).Get(&page)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("no row delete")
	}
	_, err = x.Id(page.Id).Delete(&page)
	if err != nil {
		return err
	}
	cmd.Result = &page
	return nil
}

//DeletePageByFolder 删除文件夹下所有页面
func DeletePageByFolder(cmd *models.DeletePageByFolderCommand) error {
	count, err := x.Where("folder_id = ?", cmd.FolderId).Delete(new(models.Page))
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}
