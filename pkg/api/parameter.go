package api

import (
	"encoding/json"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

//UpdateParameter 修改参数
func UpdateParameter(c *models.ReqContext, form dtos.UpdateParameterForm) Response {
	parameters := make([]models.Parameter, 0)
	errJson := json.Unmarshal([]byte(form.Data), &parameters)
	if errJson != nil {
		result := new(CommonResult)
		result.Success = false
		return JSON(200, result)
	}
	for _, pa := range parameters {
		cmd := models.UpdateParameterCommand{
			ParameterId: pa.Id,
			IsVisible:   pa.IsVisible,
			IsEditable:  pa.IsEditable,
			Default:     pa.Default,
		}
		if err := bus.Dispatch(&cmd); err != nil {
			return Error(500, "Failed to update parameter", err)
		}
	}

	//return action list
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
			return Error(500, "Failed to get parameters 3", err)
		}
		act.Parameters = paraQuery.Result
	}
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}
