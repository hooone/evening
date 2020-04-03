package api

import (
	"fmt"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/services/field"
)

//GetFields 获取页面中的所有字段
func (hs *HTTPServer) GetFields(c *dtos.ReqContext, form dtos.GetFieldsForm, lang dtos.LocaleForm) Response {
	data, err := hs.FieldService.GetFields(form.CardId, c.OrgId, lang.Language)
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

//GetFieldByID 获取字段
func (hs *HTTPServer) GetFieldByID(c *dtos.ReqContext, form dtos.GetFieldByIdForm, lang dtos.LocaleForm) Response {
	data, err := hs.FieldService.GetFieldByID(form.FieldId, c.OrgId, lang.Language)
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

//CreateField 添加字段
func (hs *HTTPServer) CreateField(c *dtos.ReqContext, form dtos.CreateFieldForm, lang dtos.LocaleForm) Response {
	fieldT := field.Field{
		Name:      form.Name,
		IsVisible: form.IsVisible,
		Type:      form.Type,
		Default:   form.Default,
		Filter:    form.Filter,
		CardId:    form.CardId,
		Text:      form.Text,
	}
	err := hs.FieldService.CreateField(fieldT, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.FieldService.GetFields(form.CardId, c.OrgId, lang.Language)
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

//UpdateField 修改字段信息
func (hs *HTTPServer) UpdateField(c *dtos.ReqContext, form dtos.UpdateFieldForm, lang dtos.LocaleForm) Response {
	fieldT := field.Field{
		Id:        form.FieldId,
		Name:      form.Name,
		Text:      form.Text,
		IsVisible: form.IsVisible,
		Type:      form.Type,
		Default:   form.Default,
		Filter:    form.Filter,
	}
	err := hs.FieldService.UpdateField(fieldT, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.FieldService.GetFields(form.CardId, c.OrgId, lang.Language)
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

//UpdateField 修改字段顺序
func (hs *HTTPServer) UpdateFieldSeq(c *dtos.ReqContext, form dtos.UpdateFieldSeqForm, lang dtos.LocaleForm) Response {
	err := hs.FieldService.UpdateFieldSeq(form.Source, form.Target, form.Position, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.FieldService.GetFields(form.CardId, c.OrgId, lang.Language)
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

//DeleteField 删除字段
func (hs *HTTPServer) DeleteField(c *dtos.ReqContext, form dtos.DeleteFieldForm, lang dtos.LocaleForm) Response {
	fieldT := field.Field{
		Id: form.FieldId,
	}
	err := hs.FieldService.DeleteField(fieldT, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.FieldService.GetFields(form.CardId, c.OrgId, lang.Language)
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
