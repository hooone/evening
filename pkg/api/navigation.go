package api

import (
	"fmt"
	"strconv"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/services/navigation"
)

//GetNavigation get all navigation
func (hs *HTTPServer) GetNavigation(c *dtos.ReqContext, lang dtos.LocaleForm) Response {
	data, err := hs.NavigationService.GetNavigation(c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	result.Message = c.SignedInUser.Name
	result.Data = data
	result.Success = true
	return JSON(200, result)
}

//RenderData get all pages/cards/data
func (hs *HTTPServer) RenderData(c *dtos.ReqContext, lang dtos.LocaleForm) Response {
	data, err := hs.NavigationService.GetNavigation(c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}

	for _, folder := range data {
		for _, page := range folder.Pages {
			cards, _ := hs.CardService.GetCards(page.Id, "", c.OrgId, lang.Language)
			for _, card := range cards {
				for _, act := range card.Actions {
					if act.Type == "READ" {
						for _, param := range act.Parameters {
							if param.Field.Type == "date" || param.Field.Type == "datetime" {
								nt64, err3 := strconv.ParseInt(param.Default, 10, 64)
								if err3 != nil {
									param.Value = 0
								} else {
									param.Value = nt64
								}
							} else {
								param.Value = param.Default
							}
						}
						testdata, _ := getTestData(card.Id, act.Parameters)
						card.Data = testdata
					}
				}
			}
			page.Cards = cards
		}
	}

	result.Message = c.SignedInUser.Name
	result.Data = data
	result.Success = true
	return JSON(200, result)
}

//CreateFolder 添加文件夹
func (hs *HTTPServer) CreateFolder(c *dtos.ReqContext, form dtos.CreateFolderForm, lang dtos.LocaleForm) Response {
	f := navigation.Folder{
		Name: form.Name,
		Text: form.Text,
	}
	err := hs.NavigationService.CreateFolder(f, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = 0
	result.Success = true
	return JSON(200, result)
}

//UpdateFolder 修改文件夹信息
func (hs *HTTPServer) UpdateFolder(c *dtos.ReqContext, form dtos.UpdateFolderForm, lang dtos.LocaleForm) Response {
	f := navigation.Folder{
		Id:   form.FolderId,
		Name: form.Name,
		Text: form.Text,
	}
	err := hs.NavigationService.UpdateFolder(f, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = 0
	result.Success = true
	return JSON(200, result)
}

//DeleteFolder 删除文件夹
func (hs *HTTPServer) DeleteFolder(c *dtos.ReqContext, form dtos.DeleteFolderForm, lang dtos.LocaleForm) Response {
	f := navigation.Folder{
		Id: form.Id,
	}
	err := hs.NavigationService.DeleteFolder(f, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = 0
	result.Success = true
	return JSON(200, result)
}

//CreateTreePage 添加一级页面
func (hs *HTTPServer) CreateTreePage(c *dtos.ReqContext, form dtos.CreateTreePageForm, lang dtos.LocaleForm) Response {
	f := navigation.Page{
		Name: form.Name,
		Text: form.Text,
	}
	err := hs.NavigationService.CreateTreePage(f, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = "/" + form.Name
	result.Success = true
	return JSON(200, result)
}

//CreateNodePage 添加子页面
func (hs *HTTPServer) CreateNodePage(c *dtos.ReqContext, form dtos.CreateNodePageForm, lang dtos.LocaleForm) Response {
	f := navigation.Page{
		Name:     form.Name,
		Text:     form.Text,
		FolderId: form.FolderId,
	}
	data, err := hs.NavigationService.CreateNodePage(f, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = data
	result.Success = true
	return JSON(200, result)
}

//UpdatePage 修改页面信息
func (hs *HTTPServer) UpdatePage(c *dtos.ReqContext, form dtos.UpdatePageForm, lang dtos.LocaleForm) Response {
	f := navigation.Page{
		Id:   form.PageID,
		Name: form.Name,
		Text: form.Text,
	}
	err := hs.NavigationService.UpdatePage(f, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = 0
	result.Success = true
	return JSON(200, result)
}

//DeletePage 删除页面
func (hs *HTTPServer) DeletePage(c *dtos.ReqContext, form dtos.DeletePageForm, lang dtos.LocaleForm) Response {
	f := navigation.Page{
		Id: form.PageID,
	}
	err := hs.NavigationService.DeletePage(f, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = 0
	result.Success = true
	return JSON(200, result)
}

//NavMove 导航栏拖动
func (hs *HTTPServer) NavMove(c *dtos.ReqContext, form dtos.NavMoveForm, lang dtos.LocaleForm) Response {
	err := hs.NavigationService.NavMove(form.SourceFolder, form.SourcePage,
		form.TargetFolder, form.TargetPage, form.Position, c.OrgId)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		hs.log.Error("NavMove err", "msg", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.NavigationService.GetNavigation(c.OrgId, lang.Language)
	if err2 != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		hs.log.Error("GetNavigation err", "msg", err)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = data
	result.Success = true
	return JSON(200, result)

}
