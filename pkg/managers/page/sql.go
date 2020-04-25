package page

import (
	"errors"
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetPagesByFolder)
	bus.AddHandler("sql", GetPageByID)
	bus.AddHandler("sql", GetPageByName)
	bus.AddHandler("sql", CreatePage)
	bus.AddHandler("sql", UpdatePage)
	bus.AddHandler("sql", UpdatePageSeq)
	bus.AddHandler("sql", DeletePage)
	bus.AddHandler("sql", DeletePageByFolder)
}

var (
	ErrPageNotFound = errors.New("Page not found")
)

//GetPagesByFolder 获取文件夹里的所有页面
func GetPagesByFolder(query *GetPagesQuery) error {
	pages := make([]*Page, 0)
	err := sqlstore.X.Where("folder_id = ? and org_id=?", query.FolderId, query.OrgId).Find(&pages)
	if err != nil {
		return err
	}
	query.Result = pages
	return nil
}

//GetPageByID 获取页面
func GetPageByID(query *GetPageByIDQuery) error {
	page := Page{
		Id: query.PageId,
	}
	success, err := sqlstore.X.Where("id=? and org_id=?", query.PageId, query.OrgId).Get(&page)
	if err != nil {
		return err
	}
	if !success {
		return ErrPageNotFound
	}
	query.Result = &page
	return nil
}

//GetPageByName 获取页面
func GetPageByName(query *GetPageByNameQuery) error {
	page := Page{}
	success, err := sqlstore.X.Where("name = ? and org_id=?", query.PageName, query.OrgId).Get(&page)
	if err != nil {
		return err
	}
	if !success {
		return ErrPageNotFound
	}
	query.Result = &page
	return nil
}

//CreatePage 添加页面
func CreatePage(cmd *CreatePageCommand) error {
	page := Page{
		Name:     cmd.Name,
		Seq:      cmd.Seq,
		FolderId: cmd.FolderId,
		OrgId:    cmd.OrgId,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	_, err := sqlstore.X.Insert(&page)
	if err != nil {
		return err
	}
	cmd.Result = page.Id
	return nil
}

//UpdatePage 修改页面信息
func UpdatePage(cmd *UpdatePageCommand) error {
	page := Page{
		Id:    cmd.PageId,
		OrgId: cmd.OrgId,
	}
	success, err := sqlstore.X.Where("id=? and org_id=?", cmd.PageId, cmd.OrgId).Get(&page)
	if err != nil {
		return err
	}
	if !success {
		return ErrPageNotFound
	}
	page.UpdateAt = time.Now()
	page.Name = cmd.Name
	_, err = sqlstore.X.Where("id=? and org_id=?", page.Id, cmd.OrgId).Cols("name").Cols("update_at").Update(&page)
	if err != nil {
		return err
	}
	cmd.Result = &page
	return nil
}

//UpdatePageSeq 修改页面顺序
func UpdatePageSeq(cmd *UpdatePageSeqCommand) error {
	page := Page{
		Id:       cmd.PageId,
		FolderId: cmd.FolderId,
		Seq:      cmd.Seq,
		OrgId:    cmd.OrgId,
		UpdateAt: time.Now(),
	}
	count, err := sqlstore.X.Where("id=? and org_id=?", cmd.PageId, cmd.OrgId).Cols("seq").Cols("folder_id").Cols("update_at").Update(&page)
	if err != nil {
		return err
	}

	cmd.Result = count
	return nil
}

//DeletePage 删除页面
func DeletePage(cmd *DeletePageCommand) error {
	page := Page{}
	success, err := sqlstore.X.Where("id=? and org_id=?", cmd.PageId, cmd.OrgId).Get(&page)
	if err != nil {
		return err
	}
	if !success {
		return ErrPageNotFound
	}
	_, err = sqlstore.X.Where("id=? and org_id=?", cmd.PageId, cmd.OrgId).Delete(&page)
	if err != nil {
		return err
	}
	cmd.Result = &page
	return nil
}

//DeletePageByFolder 删除文件夹下所有页面
func DeletePageByFolder(cmd *DeletePageByFolderCommand) error {
	count, err := sqlstore.X.Where("folder_id = ? and org_id=?", cmd.FolderId, cmd.OrgId).Delete(new(Page))
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}
