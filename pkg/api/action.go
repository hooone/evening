package api

import (
	"fmt"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/services/action"
)

//GetViewActions 获取页面中的所有操作
func (hs *HTTPServer) GetViewActions(c *dtos.ReqContext, form dtos.GetViewActionsForm, lang dtos.LocaleForm) Response {
	data, err := hs.ActionService.GetViewActions(form.CardId, c.OrgId, lang.Language)
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

//GetViewActionByID 获取操作
func (hs *HTTPServer) GetViewActionByID(c *dtos.ReqContext, form dtos.GetViewActionByIdForm, lang dtos.LocaleForm) Response {
	data, err := hs.ActionService.GetViewActionByID(form.ViewActionId, c.OrgId, lang.Language)
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

//CreateViewAction 添加操作
func (hs *HTTPServer) CreateViewAction(c *dtos.ReqContext, form dtos.CreateViewActionForm, lang dtos.LocaleForm) Response {
	actionT := action.ViewAction{
		Name:        form.Name,
		Text:        form.Text,
		Type:        form.Type,
		DoubleCheck: form.DoubleCheck,
		CardId:      form.CardId,
	}
	err := hs.ActionService.CreateViewAction(actionT, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.ActionService.GetViewActions(form.CardId, c.OrgId, lang.Language)
	if err2 != nil {
		result.Data = 2
		result.Message = fmt.Sprintf("%s", err2)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = data
	result.Success = true
	return JSON(200, result)
}

//UpdateViewAction 修改操作信息
func (hs *HTTPServer) UpdateViewAction(c *dtos.ReqContext, form dtos.UpdateViewActionForm, lang dtos.LocaleForm) Response {
	actionT := action.ViewAction{
		Id:          form.ViewActionId,
		Name:        form.Name,
		Text:        form.Text,
		DoubleCheck: form.DoubleCheck,
		Type:        form.Type,
	}
	err := hs.ActionService.UpdateViewAction(actionT, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.ActionService.GetViewActions(form.CardId, c.OrgId, lang.Language)
	if err2 != nil {
		result.Data = 2
		result.Message = fmt.Sprintf("%s", err2)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = data
	result.Success = true
	return JSON(200, result)
}

//UpdateViewActionSeq 修改操作顺序
func (hs *HTTPServer) UpdateViewActionSeq(c *dtos.ReqContext, form dtos.UpdateViewActionSeqForm, lang dtos.LocaleForm) Response {
	err := hs.ActionService.UpdateViewActionSeq(form.Source, form.Target, form.Position, c.OrgId)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.ActionService.GetViewActions(form.CardId, c.OrgId, lang.Language)
	if err2 != nil {
		result.Data = 2
		result.Message = fmt.Sprintf("%s", err2)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = data
	result.Success = true
	return JSON(200, result)
}

//DeleteViewAction 删除操作
func (hs *HTTPServer) DeleteViewAction(c *dtos.ReqContext, form dtos.DeleteViewActionForm, lang dtos.LocaleForm) Response {
	actionT := action.ViewAction{
		Id: form.ViewActionId,
	}
	err := hs.ActionService.DeleteViewAction(actionT, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.ActionService.GetViewActions(form.CardId, c.OrgId, lang.Language)
	if err2 != nil {
		result.Data = 2
		result.Message = fmt.Sprintf("%s", err2)
		result.Success = false
		return JSON(200, result)
	}
	result.Data = data
	result.Success = true
	return JSON(200, result)
}
