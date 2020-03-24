package api

import (
	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

//GetFields 获取页面中的所有字段
func GetFields(c *models.ReqContext, form dtos.GetFieldsForm) Response {
	query := models.GetFieldsQuery{CardId: form.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	for _, cd := range query.Result {
		lQuery := models.GetLocaleQuery{Name: cd.Name, VId: cd.Id, Type: "field"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		cd.Locale = lQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//GetFieldByID 获取字段
func GetFieldByID(c *models.ReqContext, form dtos.GetFieldByIdForm) Response {
	query := models.GetFieldQuery{FieldId: form.FieldId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get field", err)
	}
	lQuery := models.GetLocaleQuery{Name: query.Result.Name, VId: query.Result.Id, Type: "field"}
	if err := bus.Dispatch(&lQuery); err != nil {
		return Error(500, "Failed to get locale", err)
	}
	query.Result.Locale = lQuery.Result
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//CreateField 添加字段
func CreateField(c *models.ReqContext, form dtos.CreateFieldForm, lang dtos.LocaleForm) Response {
	//calculate seq
	query := models.GetFieldsQuery{CardId: form.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	seq := int32(1)
	if len(query.Result) > 0 {
		seq = query.Result[len(query.Result)-1].Seq + 2
	}

	//create field
	cmd := models.CreateFieldCommand{
		CardId:    form.CardId,
		Name:      form.Name,
		Seq:       seq,
		IsVisible: form.IsVisible,
		Type:      form.Type,
		Default:   form.Default,
		Filter:    form.Filter,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to create  field", err)
	}

	//create parameter
	aquery := models.GetViewActionsQuery{CardId: form.CardId}
	if err := bus.Dispatch(&aquery); err != nil {
		return Error(500, "Failed to get actions", err)
	}
	for _, act := range aquery.Result {
		paCmd := models.CreateParameterCommand{
			ActionId:   act.Id,
			FieldId:    cmd.Result,
			IsVisible:  true,
			IsEditable: true,
			Default:    form.Default,
		}
		if err2 := bus.Dispatch(&paCmd); err2 != nil {
			return Error(500, "Failed to create parameter", err2)
		}
	}

	// set locale
	localeCmd := models.SetLocaleCommand{
		VId:      cmd.Result,
		Type:     "field",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)

	//get new field list
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get fields", err)
	}

	//get field locale
	for _, cd := range query.Result {
		lQuery := models.GetLocaleQuery{Name: cd.Name, VId: cd.Id, Type: "field"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		cd.Locale = lQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//UpdateField 修改字段信息
func UpdateField(c *models.ReqContext, form dtos.UpdateFieldForm, lang dtos.LocaleForm) Response {
	cmd := models.UpdateFieldCommand{
		FieldId:   form.FieldId,
		Name:      form.Name,
		IsVisible: form.IsVisible,
		Type:      form.Type,
		Default:   form.Default,
		Filter:    form.Filter,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to update field", err)
	}
	query := models.GetFieldsQuery{CardId: cmd.Result.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	for idx, cd := range query.Result {
		if cd.Seq != int32(idx*2+1) {
			cd.Seq = int32(idx*2 + 1)
			seqCmd := models.UpdateFieldSeqCommand{
				FieldId: cd.Id,
				Seq:     cd.Seq,
			}
			if err := bus.Dispatch(&seqCmd); err != nil {
				return Error(500, "Failed to update field", err)
			}
		}
	}
	localeCmd := models.SetLocaleCommand{
		VId:      form.FieldId,
		Type:     "field",
		Text:     form.Text,
		Language: lang.Language,
	}
	if err := bus.Dispatch(&localeCmd); err != nil {
		return Error(500, "Failed to update locale", err)
	}

	for _, cd := range query.Result {
		lQuery := models.GetLocaleQuery{Name: cd.Name, VId: cd.Id, Type: "field"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		cd.Locale = lQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//UpdateField 修改字段顺序
func UpdateFieldSeq(c *models.ReqContext, form dtos.UpdateFieldSeqForm, lang dtos.LocaleForm) Response {
	if form.Source == 0 || form.Target == 0 {
		result := new(CommonResult)
		result.Success = false
		return JSON(200, result)
	}
	squery := models.GetFieldQuery{FieldId: form.Source}
	if err := bus.Dispatch(&squery); err != nil {
		return Error(500, "Failed to get field", err)
	}
	tquery := models.GetFieldQuery{FieldId: form.Target}
	if err := bus.Dispatch(&tquery); err != nil {
		return Error(500, "Failed to get field", err)
	}

	cmd := models.UpdateFieldSeqCommand{
		FieldId: squery.Result.Id,
		Seq:     tquery.Result.Seq + form.Position*2 - 3,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to update field", err)
	}
	query := models.GetFieldsQuery{CardId: squery.Result.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	for idx, cd := range query.Result {
		if cd.Seq != int32(idx*2+1) {
			cd.Seq = int32(idx*2 + 1)
			seqCmd := models.UpdateFieldSeqCommand{
				FieldId: cd.Id,
				Seq:     cd.Seq,
			}
			if err := bus.Dispatch(&seqCmd); err != nil {
				return Error(500, "Failed to update field", err)
			}
		}
	}

	for _, cd := range query.Result {
		lQuery := models.GetLocaleQuery{Name: cd.Name, VId: cd.Id, Type: "field"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		cd.Locale = lQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//DeleteField 删除字段
func DeleteField(c *models.ReqContext, form dtos.DeleteFieldForm) Response {
	cmd := models.DeleteFieldCommand{
		FieldId: form.FieldId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to delete field", err)
	}
	query := models.GetFieldsQuery{CardId: cmd.Result.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	for _, cd := range query.Result {
		lQuery := models.GetLocaleQuery{Name: cd.Name, VId: cd.Id, Type: "field"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		cd.Locale = lQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}
