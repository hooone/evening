package api

import (
	"errors"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
	m "github.com/hooone/evening/pkg/models"
)

//GetNavigation get all navigation
func GetNavigation(c *models.ReqContext) Response {
	query := m.GetNavigationQuery{OrgId: c.OrgId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get folders", err)
	}
	for _, fd := range query.Result {
		lQuery := m.GetLocaleQuery{Name: fd.Name, VId: fd.Id, Type: "folder"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		fd.Locale = lQuery.Result
		for _, pg := range fd.Pages {
			l2Query := m.GetLocaleQuery{Name: pg.Name, VId: pg.Id, Type: "page"}
			if err := bus.Dispatch(&l2Query); err != nil {
				return Error(500, "Failed to get locale", err)
			}
			pg.Locale = l2Query.Result
		}
	}
	result := new(CommonResult)
	result.Message = c.SignedInUser.Name
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//CreateFolder 添加文件夹
func CreateFolder(c *models.ReqContext, form dtos.CreateFolderForm, lang dtos.LocaleForm) Response {
	query := models.GetFolderQuery{OrgId: c.OrgId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get folders", err)
	}
	var seq int32
	seq = 1
	if len(query.Result) > 0 {
		seq = query.Result[len(query.Result)-1].Seq + 2
	}
	cmd := models.CreateFolderCommand{
		Name:     form.Name,
		Seq:      seq,
		IsFolder: true,
		OrgId:    c.OrgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to create folder", err)
	}
	localeCmd := models.SetLocaleCommand{
		VId:      cmd.Result,
		Type:     "folder",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)
	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

//UpdateFolder 修改文件夹信息
func UpdateFolder(c *models.ReqContext, form dtos.UpdateFolderForm, lang dtos.LocaleForm) Response {
	cmd := models.UpdateFolderCommand{
		FolderId: form.FolderId,
		Name:     form.Name,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to update folder", err)
	}
	if cmd.Result == 0 {
		return Error(500, "Failed to update folder ", errors.New("no row updated"))
	}
	localeCmd := models.SetLocaleCommand{
		VId:      form.FolderId,
		Type:     "folder",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)
	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

//DeleteFolder 删除文件夹
func DeleteFolder(c *models.ReqContext, form dtos.DeleteFolderForm) Response {
	cmd := models.DeleteFolderCommand{
		Id: form.Id,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to delete folder", err)
	}
	pageCmd := models.DeletePageByFolderCommand{
		FolderId: form.Id,
	}
	if err := bus.Dispatch(&pageCmd); err != nil {
		return Error(500, "Failed to delete pages", err)
	}
	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

//CreateTreePage 添加一级页面
func CreateTreePage(c *models.ReqContext, form dtos.CreateTreePageForm, lang dtos.LocaleForm) Response {
	//重名判断
	oldPquery := models.GetPageByNameQuery{
		PageName: form.Name,
		OrgId:    c.OrgId,
	}
	errP := bus.Dispatch(&oldPquery)
	if errP != models.ErrPageNotFound {
		resultN := new(CommonResult)
		resultN.Data = 0
		resultN.Success = false
		resultN.Message = "Page name conflict"
		return JSON(200, resultN)
	}

	query := models.GetFolderQuery{OrgId: c.OrgId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get folder", err)
	}
	var seq int32
	seq = 1
	if len(query.Result) > 0 {
		seq = query.Result[len(query.Result)-1].Seq + 2
	}
	cmd := models.CreateFolderCommand{
		Name:     form.Name,
		Seq:      seq,
		IsFolder: false,
		OrgId:    c.OrgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to create folder", err)
	}
	pageCmd := models.CreatePageCommand{
		Name:     form.Name,
		Seq:      1,
		FolderId: cmd.Result,
		OrgId:    c.OrgId,
	}
	if err := bus.Dispatch(&pageCmd); err != nil {
		return Error(500, "Failed to create page", err)
	}
	localeCmd := models.SetLocaleCommand{
		VId:      pageCmd.Result,
		Type:     "page",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)
	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

//CreateNodePage 添加子页面
func CreateNodePage(c *models.ReqContext, form dtos.CreateNodePageForm, lang dtos.LocaleForm) Response {
	//重名判断
	oldPquery := models.GetPageByNameQuery{
		PageName: form.Name,
		OrgId:    c.OrgId,
	}
	if err := bus.Dispatch(&oldPquery); err == nil {
		resultN := new(CommonResult)
		resultN.Data = 0
		resultN.Success = false
		resultN.Message = "Page name conflict"
		return JSON(200, resultN)
	}

	query := models.GetPagesQuery{FolderId: form.FolderId, OrgId: c.OrgId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get pages", err)
	}
	var seq int32
	seq = 1
	if len(query.Result) > 0 {
		seq = query.Result[len(query.Result)-1].Seq + 2
	}
	cmd := models.CreatePageCommand{
		Name:     form.Name,
		Seq:      seq,
		FolderId: form.FolderId,
		OrgId:    c.OrgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to create folder", err)
	}
	localeCmd := models.SetLocaleCommand{
		VId:      cmd.Result,
		Type:     "page",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)
	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

//UpdatePage 修改页面信息
func UpdatePage(c *models.ReqContext, form dtos.UpdatePageForm, lang dtos.LocaleForm) Response {
	cmd := models.UpdatePageCommand{
		Name:   form.Name,
		PageId: form.PageID,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to update page", err)
	}
	query := models.GetFolderByIDQuery{
		OrgId:    c.OrgId,
		FolderId: cmd.Result.FolderId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get folder", err)
	}
	if !query.Result.IsFolder {
		fdCmd := models.UpdateFolderCommand{
			FolderId: query.Result.Id,
			Name:     cmd.Result.Name,
		}
		if err := bus.Dispatch(&fdCmd); err != nil {
			return Error(500, "Failed to update folder", err)
		}
	}
	localeCmd := models.SetLocaleCommand{
		VId:      form.PageID,
		Type:     "page",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)
	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

//DeletePage 删除页面
func DeletePage(c *models.ReqContext, form dtos.DeletePageForm) Response {
	cmd := models.DeletePageCommand{
		PageId: form.PageID,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to delete page", err)
	}
	query := models.GetFolderByIDQuery{
		OrgId:    c.OrgId,
		FolderId: cmd.Result.FolderId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get folder", err)
	}
	if !query.Result.IsFolder {
		fdCmd := models.DeleteFolderCommand{
			Id: query.Result.Id,
		}
		if err := bus.Dispatch(&fdCmd); err != nil {
			return Error(500, "Failed to delete folder", err)
		}
	}
	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

//NavMove 导航栏拖动
func NavMove(c *models.ReqContext, form dtos.NavMoveForm) Response {
	result := new(CommonResult)
	//get all need data
	sFquery := models.GetFolderByIDQuery{
		OrgId:    c.OrgId,
		FolderId: form.SourceFolder,
	}
	if form.SourceFolder != 0 {
		if err := bus.Dispatch(&sFquery); err != nil {
			return Error(500, "Failed to get folder", err)
		}
	}
	tFquery := models.GetFolderByIDQuery{
		OrgId:    c.OrgId,
		FolderId: form.TargetFolder,
	}
	if form.TargetFolder != 0 {
		if err := bus.Dispatch(&tFquery); err != nil {
			return Error(500, "Failed to get folder", err)
		}
	}
	sPquery := models.GetPageByIDQuery{
		PageId: form.SourcePage,
		OrgId:  c.OrgId,
	}
	if form.SourcePage != 0 {
		if err := bus.Dispatch(&sPquery); err != nil {
			return Error(500, "Failed to get page", err)
		}
		sFquery.FolderId = sPquery.Result.FolderId
		if err := bus.Dispatch(&sFquery); err != nil {
			return Error(500, "Failed to get folder", err)
		}
	}
	tPquery := models.GetPageByIDQuery{
		PageId: form.TargetPage,
		OrgId:  c.OrgId,
	}
	if form.TargetPage != 0 {
		if err := bus.Dispatch(&tPquery); err != nil {
			return Error(500, "Failed to get page", err)
		}
		tFquery.FolderId = tPquery.Result.FolderId
		if err := bus.Dispatch(&tFquery); err != nil {
			return Error(500, "Failed to get folder", err)
		}
	}
	//move
	if form.SourceFolder != 0 && form.TargetFolder != 0 {
		//folder to folder
		if sFquery.Result.IsFolder && tFquery.Result.IsFolder {
			moveCmd := models.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
			}
			if form.Position == 1 {
				moveCmd.Seq = tFquery.Result.Seq - 1
			}
			if form.Position == 3 {
				moveCmd.Seq = tFquery.Result.Seq + 1
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return Error(500, "Failed to move folder", err)
			}
		}
	} else if form.SourceFolder != 0 && form.TargetPage != 0 {
		if tFquery.Result.IsFolder {
			//folder to node page
			moveCmd := models.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
			}
			moveCmd.Seq = tFquery.Result.Seq + 1
			if err := bus.Dispatch(&moveCmd); err != nil {
				return Error(500, "Failed to move folder", err)
			}
		} else {
			//folder to tree page
			moveCmd := models.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
			}
			if form.Position == 1 {
				moveCmd.Seq = tFquery.Result.Seq - 1
			}
			if form.Position == 3 {
				moveCmd.Seq = tFquery.Result.Seq + 1
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return Error(500, "Failed to move folder", err)
			}
		}
	} else if form.SourcePage != 0 && form.TargetFolder != 0 {
		if sFquery.Result.IsFolder {
			//node page to folder
			if form.Position == 2 {
				if sFquery.Result.Id != tFquery.Result.Id {
					moveCmd := models.UpdatePageSeqCommand{
						PageId:   sPquery.Result.Id,
						FolderId: tFquery.Result.Id,
						Seq:      99999,
					}
					if err := bus.Dispatch(&moveCmd); err != nil {
						return Error(500, "Failed to move page", err)
					}
				}
			} else {
				//create fake folder
				createCmd := models.CreateFolderCommand{
					Name:     sPquery.Result.Name,
					IsFolder: false,
					OrgId:    c.OrgId,
				}
				if form.Position == 1 {
					createCmd.Seq = tFquery.Result.Seq - 1
				}
				if form.Position == 3 {
					createCmd.Seq = tFquery.Result.Seq + 1
				}
				if err := bus.Dispatch(&createCmd); err != nil {
					return Error(500, "Failed to create folder", err)
				}
				//move page
				moveCmd := models.UpdatePageSeqCommand{
					PageId:   sPquery.Result.Id,
					FolderId: createCmd.Result,
					Seq:      1,
				}
				if err := bus.Dispatch(&moveCmd); err != nil {
					return Error(500, "Failed to move page", err)
				}
				//arrange folder
				moveFCmd := models.UpdateFolderSeqCommand{
					FolderId: createCmd.Result,
					Seq:      createCmd.Seq,
				}
				if err := bus.Dispatch(&moveFCmd); err != nil {
					return Error(500, "Failed to move folder", err)
				}
			}
		} else {
			//tree page to folder
			if form.Position == 2 {
				if sFquery.Result.Id != tFquery.Result.Id {
					moveCmd := models.UpdatePageSeqCommand{
						PageId:   sPquery.Result.Id,
						FolderId: tFquery.Result.Id,
						Seq:      99999,
					}
					if err := bus.Dispatch(&moveCmd); err != nil {
						return Error(500, "Failed to move page", err)
					}
					deleteCmd := models.DeleteFolderCommand{
						Id: sFquery.Result.Id,
					}
					if err := bus.Dispatch(&deleteCmd); err != nil {
						return Error(500, "Failed to delete fake folder", err)
					}
				}
			} else {
				//move folder
				moveCmd := models.UpdateFolderSeqCommand{
					FolderId: sFquery.Result.Id,
				}
				if form.Position == 1 {
					moveCmd.Seq = tFquery.Result.Seq - 1
				}
				if form.Position == 3 {
					moveCmd.Seq = tFquery.Result.Seq + 1
				}
				if err := bus.Dispatch(&moveCmd); err != nil {
					return Error(500, "Failed to move folder", err)
				}
			}
		}
	} else if form.SourcePage != 0 && form.TargetPage != 0 {
		if sFquery.Result.IsFolder && tFquery.Result.IsFolder {
			//node page to node page
			moveCmd := models.UpdatePageSeqCommand{
				PageId:   sPquery.Result.Id,
				FolderId: tFquery.Result.Id,
			}
			if form.Position == 1 {
				moveCmd.Seq = tPquery.Result.Seq - 1
			}
			if form.Position == 3 {
				moveCmd.Seq = tPquery.Result.Seq + 1
			}

			if err := bus.Dispatch(&moveCmd); err != nil {
				return Error(500, "Failed to move page", err)
			}
		} else if sFquery.Result.IsFolder && !tFquery.Result.IsFolder {
			//node page to tree page
			//create fake folder
			createCmd := models.CreateFolderCommand{
				Name:     sPquery.Result.Name,
				IsFolder: false,
				OrgId:    c.OrgId,
			}
			if form.Position == 1 {
				createCmd.Seq = tFquery.Result.Seq - 1
			}
			if form.Position == 3 {
				createCmd.Seq = tFquery.Result.Seq + 1
			}
			if err := bus.Dispatch(&createCmd); err != nil {
				return Error(500, "Failed to create folder", err)
			}
			//move page
			moveCmd := models.UpdatePageSeqCommand{
				PageId:   sPquery.Result.Id,
				FolderId: createCmd.Result,
				Seq:      1,
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return Error(500, "Failed to move page", err)
			}
			//arrange folder
			moveFCmd := models.UpdateFolderSeqCommand{
				FolderId: createCmd.Result,
				Seq:      createCmd.Seq,
			}
			if err := bus.Dispatch(&moveFCmd); err != nil {
				return Error(500, "Failed to move folder", err)
			}
		} else if !sFquery.Result.IsFolder && tFquery.Result.IsFolder {
			//tree page to node page
			if sFquery.Result.Id != tFquery.Result.Id {
				moveCmd := models.UpdatePageSeqCommand{
					PageId:   sPquery.Result.Id,
					FolderId: tFquery.Result.Id,
				}
				if form.Position == 1 {
					moveCmd.Seq = tPquery.Result.Seq - 1
				}
				if form.Position == 3 {
					moveCmd.Seq = tPquery.Result.Seq + 1
				}
				if err := bus.Dispatch(&moveCmd); err != nil {
					return Error(500, "Failed to move page", err)
				}
				deleteCmd := models.DeleteFolderCommand{
					Id: sFquery.Result.Id,
				}
				if err := bus.Dispatch(&deleteCmd); err != nil {
					return Error(500, "Failed to delete fake folder", err)
				}
			}
		} else if !sFquery.Result.IsFolder && !tFquery.Result.IsFolder {
			//tree page to tree page
			moveCmd := models.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
			}
			if form.Position == 1 {
				moveCmd.Seq = tFquery.Result.Seq - 1
			}
			if form.Position == 3 {
				moveCmd.Seq = tFquery.Result.Seq + 1
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return Error(500, "Failed to move folder", err)
			}
		}
	} else if form.SourceFolder != 0 && form.TargetFolder == 0 && form.TargetPage == 0 {
		//folder to empty
		moveCmd := models.UpdateFolderSeqCommand{
			FolderId: sFquery.Result.Id,
			Seq:      99999,
		}
		if err := bus.Dispatch(&moveCmd); err != nil {
			return Error(500, "Failed to move folder", err)
		}
	} else if form.SourcePage != 0 && form.TargetFolder == 0 && form.TargetPage == 0 {
		if sFquery.Result.IsFolder {
			//node page to empty
			//create fake folder
			createCmd := models.CreateFolderCommand{
				Name:     sPquery.Result.Name,
				IsFolder: false,
				Seq:      99999,
				OrgId:    c.OrgId,
			}
			if err := bus.Dispatch(&createCmd); err != nil {
				return Error(500, "Failed to create folder", err)
			}
			//move page
			moveCmd := models.UpdatePageSeqCommand{
				PageId:   sPquery.Result.Id,
				FolderId: createCmd.Result,
				Seq:      1,
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return Error(500, "Failed to move page", err)
			}
			//arrange folder
			moveFCmd := models.UpdateFolderSeqCommand{
				FolderId: createCmd.Result,
				Seq:      createCmd.Seq,
			}
			if err := bus.Dispatch(&moveFCmd); err != nil {
				return Error(500, "Failed to move folder", err)
			}
		} else {
			//tree page to empty
			moveCmd := models.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
				Seq:      99999,
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return Error(500, "Failed to move folder", err)
			}
		}
	}
	result.Success = true
	return JSON(200, result)
}
