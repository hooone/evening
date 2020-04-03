package api

import (
	"encoding/json"
	"fmt"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/services/parameter"
)

//UpdateParameter 修改参数
func (hs *HTTPServer) GetParameters(c *dtos.ReqContext, form dtos.GetParametersForm, lang dtos.LocaleForm) Response {

	result := new(CommonResult)
	data, err2 := hs.ParameterService.GetParameters(form.ActionId, c.OrgId, lang.Language)
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

//UpdateParameter 修改参数
func (hs *HTTPServer) UpdateParameter(c *dtos.ReqContext, form dtos.UpdateParameterForm, lang dtos.LocaleForm) Response {

	result := new(CommonResult)
	parameters := make([]parameter.Parameter, 0)
	errJson := json.Unmarshal([]byte(form.Data), &parameters)
	if errJson != nil {
		result.Message = fmt.Sprintf("%s", errJson)
		result.Success = false
		return JSON(200, result)
	}

	for _, paramT := range parameters {
		err := hs.ParameterService.UpdateParameter(paramT, c.OrgId)
		if err != nil {
			result.Data = 1
			result.Message = fmt.Sprintf("%s", err)
			result.Success = false
			return JSON(200, result)
		}
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
