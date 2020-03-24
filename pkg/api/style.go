package api

import (
	"encoding/json"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

//UpdateStyles 更新Style
func UpdateStyles(c *models.ReqContext, form dtos.UpdateStylesForm) Response {
	styles := make([]models.Style, 0)
	errJSON := json.Unmarshal([]byte(form.Data), &styles)
	if errJSON != nil {
		result := new(CommonResult)
		result.Success = false
		return JSON(200, result)
	}
	for _, st := range styles {
		cmd := models.UpdateStyleCommand{
			Id:      st.Id,
			FieldId: st.FieldId,
			Value:   st.Value,
		}
		if err := bus.Dispatch(&cmd); err != nil {
			return Error(500, "Failed to update styles", err)
		}
	}
	result := new(CommonResult)
	result.Data = len(styles)
	result.Success = true
	return JSON(200, result)
}
