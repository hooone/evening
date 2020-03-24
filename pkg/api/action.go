package api

import (
	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

//GetViewActions 获取页面中的所有操作
func GetViewActions(c *models.ReqContext, form dtos.GetViewActionsForm) Response {
	query := models.GetViewActionsQuery{CardId: form.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get view actions", err)
	}
	for _, act := range query.Result {
		lQuery := models.GetLocaleQuery{Name: act.Name, VId: act.Id, Type: "action"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		act.Locale = lQuery.Result
		//get parameters
		paraQuery := models.GetParametersQuery{ActionId: act.Id}
		if err := bus.Dispatch(&paraQuery); err != nil {
			return Error(500, "Failed to get parameters", err)
		}
		act.Parameters = paraQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//GetViewActionByID 获取操作
func GetViewActionByID(c *models.ReqContext, form dtos.GetViewActionByIdForm) Response {
	query := models.GetViewActionQuery{ViewActionId: form.ViewActionId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get view action", err)
	}
	lQuery := models.GetLocaleQuery{Name: query.Result.Name, VId: query.Result.Id, Type: "action"}
	if err := bus.Dispatch(&lQuery); err != nil {
		return Error(500, "Failed to get locale", err)
	}
	query.Result.Locale = lQuery.Result
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//CreateViewAction 添加操作
func CreateViewAction(c *models.ReqContext, form dtos.CreateViewActionForm, lang dtos.LocaleForm) Response {
	//calculate seq
	query := models.GetViewActionsQuery{CardId: form.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get view actions", err)
	}
	seq := int32(1)
	if len(query.Result) > 0 {
		seq = query.Result[len(query.Result)-1].Seq + 2
	}

	//create action
	cmd := models.CreateViewActionCommand{
		CardId:      form.CardId,
		Name:        form.Name,
		Seq:         seq,
		DoubleCheck: form.DoubleCheck,
		Type:        form.Type,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to create  view action", err)
	}

	//create parameter
	fquery := models.GetFieldsQuery{CardId: form.CardId}
	if err := bus.Dispatch(&fquery); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	for _, field := range fquery.Result {
		paCmd := models.CreateParameterCommand{
			ActionId:   cmd.Result,
			FieldId:    field.Id,
			IsVisible:  true,
			IsEditable: true,
			Default:    field.Default,
		}
		if err2 := bus.Dispatch(&paCmd); err2 != nil {
			return Error(500, "Failed to create parameter", err2)
		}
	}

	//set locale
	localeCmd := models.SetLocaleCommand{
		VId:      cmd.Result,
		Type:     "action",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)

	//get new action list
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get view actions", err)
	}
	for _, act := range query.Result {
		lQuery := models.GetLocaleQuery{Name: act.Name, VId: act.Id, Type: "action"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		act.Locale = lQuery.Result
		//get parameters
		paraQuery := models.GetParametersQuery{ActionId: act.Id}
		if err := bus.Dispatch(&paraQuery); err != nil {
			return Error(500, "Failed to get parameters", err)
		}
		act.Parameters = paraQuery.Result
	}

	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//UpdateViewAction 修改操作信息
func UpdateViewAction(c *models.ReqContext, form dtos.UpdateViewActionForm, lang dtos.LocaleForm) Response {
	cmd := models.UpdateViewActionCommand{
		ViewActionId: form.ViewActionId,
		Name:         form.Name,
		DoubleCheck:  form.DoubleCheck,
		Type:         form.Type,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to update view action", err)
	}
	query := models.GetViewActionsQuery{CardId: cmd.Result.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get view actions", err)
	}
	for idx, cd := range query.Result {
		if cd.Seq != int32(idx*2+1) {
			cd.Seq = int32(idx*2 + 1)
			seqCmd := models.UpdateViewActionSeqCommand{
				ViewActionId: cd.Id,
				Seq:          cd.Seq,
			}
			if err := bus.Dispatch(&seqCmd); err != nil {
				return Error(500, "Failed to update view action", err)
			}
		}
	}
	localeCmd := models.SetLocaleCommand{
		VId:      form.ViewActionId,
		Type:     "action",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)

	for _, act := range query.Result {
		lQuery := models.GetLocaleQuery{Name: act.Name, VId: act.Id, Type: "action"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		act.Locale = lQuery.Result
		//get parameters
		paraQuery := models.GetParametersQuery{ActionId: act.Id}
		if err := bus.Dispatch(&paraQuery); err != nil {
			return Error(500, "Failed to get parameters", err)
		}
		act.Parameters = paraQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//UpdateViewAction 修改操作顺序
func UpdateViewActionSeq(c *models.ReqContext, form dtos.UpdateViewActionSeqForm, lang dtos.LocaleForm) Response {
	if form.Source == 0 || form.Target == 0 {
		result := new(CommonResult)
		result.Success = false
		return JSON(200, result)
	}
	squery := models.GetViewActionQuery{ViewActionId: form.Source}
	if err := bus.Dispatch(&squery); err != nil {
		return Error(500, "Failed to get view action", err)
	}
	tquery := models.GetViewActionQuery{ViewActionId: form.Target}
	if err := bus.Dispatch(&tquery); err != nil {
		return Error(500, "Failed to get view action", err)
	}

	cmd := models.UpdateViewActionSeqCommand{
		ViewActionId: squery.Result.Id,
		Seq:          tquery.Result.Seq + form.Position*2 - 3,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to update view action", err)
	}
	query := models.GetViewActionsQuery{CardId: squery.Result.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get view actions", err)
	}
	for idx, cd := range query.Result {
		if cd.Seq != int32(idx*2+1) {
			cd.Seq = int32(idx*2 + 1)
			seqCmd := models.UpdateViewActionSeqCommand{
				ViewActionId: cd.Id,
				Seq:          cd.Seq,
			}
			if err := bus.Dispatch(&seqCmd); err != nil {
				return Error(500, "Failed to update view action", err)
			}
		}
	}

	for _, act := range query.Result {
		lQuery := models.GetLocaleQuery{Name: act.Name, VId: act.Id, Type: "action"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		act.Locale = lQuery.Result
		//get parameters
		paraQuery := models.GetParametersQuery{ActionId: act.Id}
		if err := bus.Dispatch(&paraQuery); err != nil {
			return Error(500, "Failed to get parameters", err)
		}
		act.Parameters = paraQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//DeleteViewAction 删除操作
func DeleteViewAction(c *models.ReqContext, form dtos.DeleteViewActionForm) Response {
	cmd := models.DeleteViewActionCommand{
		ViewActionId: form.ViewActionId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to delete view action", err)
	}
	query := models.GetViewActionsQuery{CardId: cmd.Result.CardId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get view actions", err)
	}
	for _, act := range query.Result {
		lQuery := models.GetLocaleQuery{Name: act.Name, VId: act.Id, Type: "action"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		act.Locale = lQuery.Result
		//get parameters
		paraQuery := models.GetParametersQuery{ActionId: act.Id}
		if err := bus.Dispatch(&paraQuery); err != nil {
			return Error(500, "Failed to get parameters", err)
		}
		act.Parameters = paraQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}
