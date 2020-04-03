package api

import (
	"fmt"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/services/card"
)

//GetCards 获取页面中的所有窗体
func (hs *HTTPServer) GetCards(c *dtos.ReqContext, form dtos.GetCardsForm, lang dtos.LocaleForm) Response {
	data, err := hs.CardService.GetCards(form.PageId, form.PageName, c.OrgId, lang.Language)
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

//GetCardByID 获取窗体
func (hs *HTTPServer) GetCardByID(c *dtos.ReqContext, form dtos.GetCardByIdForm, lang dtos.LocaleForm) Response {
	data, err := hs.CardService.GetCardByID(form.CardId, c.OrgId, lang.Language)
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

//CreateCard 添加窗体
func (hs *HTTPServer) CreateCard(c *dtos.ReqContext, form dtos.CreateCardForm, lang dtos.LocaleForm) Response {
	cardT := card.Card{
		PageId: form.PageId,
		Name:   form.Name,
		Text:   form.Text,
	}
	err := hs.CardService.CreateCard(form.PageId, form.PageName, cardT, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.CardService.GetCards(form.PageId, form.PageName, c.OrgId, lang.Language)
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

//UpdateCard 修改窗体信息
func (hs *HTTPServer) UpdateCard(c *dtos.ReqContext, form dtos.UpdateCardForm, lang dtos.LocaleForm) Response {
	cardT := card.Card{
		Id:    form.CardId,
		Name:  form.Name,
		Style: form.Style,
		Text:  form.Text,
		Width: form.Width,
		Pos:   form.Pos,
	}
	err := hs.CardService.UpdateCard(cardT, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.CardService.GetCardByID(form.CardId, c.OrgId, lang.Language)
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

//UpdateCardSeq 修改窗体顺序
func (hs *HTTPServer) UpdateCardSeq(c *dtos.ReqContext, form dtos.UpdateCardSeqForm, lang dtos.LocaleForm) Response {
	err := hs.CardService.UpdateCardSeq(form.Source, form.Target, form.Position, c.OrgId, lang.Language)
	result := new(CommonResult)
	if err != nil {
		result.Data = 1
		result.Message = fmt.Sprintf("%s", err)
		result.Success = false
		return JSON(200, result)
	}
	data, err2 := hs.CardService.GetCards(form.PageId, "", c.OrgId, lang.Language)
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

//DeleteCard 删除窗体
func (hs *HTTPServer) DeleteCard(c *dtos.ReqContext, form dtos.DeleteCardForm, lang dtos.LocaleForm) Response {
	cardT := card.Card{
		Id: form.CardId,
	}
	err := hs.CardService.DeleteCard(cardT, c.OrgId, lang.Language)
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
