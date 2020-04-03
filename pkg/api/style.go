package api

import (
	"encoding/json"
	"fmt"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/services/style"
)

//UpdateStyles 更新Style
func (hs *HTTPServer) UpdateStyles(c *dtos.ReqContext, form dtos.UpdateStylesForm, lang dtos.LocaleForm) Response {
	result := new(CommonResult)
	styles := make([]style.Style, 0)
	errJson := json.Unmarshal([]byte(form.Data), &styles)
	if errJson != nil {
		result.Message = fmt.Sprintf("%s", errJson)
		result.Success = false
		return JSON(200, result)
	}

	for _, styleT := range styles {
		err := hs.StyleService.UpdateStyles(styleT, c.OrgId, lang.Language)
		if err != nil {
			result.Data = 1
			result.Message = fmt.Sprintf("%s", err)
			result.Success = false
			return JSON(200, result)
		}
	}
	result.Data = 0
	result.Success = true
	return JSON(200, result)
}
