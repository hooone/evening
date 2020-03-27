package api

import (
	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

//GetCards 获取页面中的所有卡片
func GetCards(c *models.ReqContext, form dtos.GetCardsForm) Response {
	cards := make([]*models.Card, 0)
	if form.PageId != 0 {
		query := models.GetCardsQuery{
			PageId: form.PageId,
			OrgId:  c.OrgId,
		}
		if err := bus.Dispatch(&query); err != nil {
			return Error(500, "Failed to get cards", err)
		}
		cards = query.Result
	} else {
		pageQuery := models.GetPageByNameQuery{
			PageName: form.PageName,
			OrgId:    c.OrgId,
		}
		if err := bus.Dispatch(&pageQuery); err != nil {
			if err == models.ErrPageNotFound {
				result404 := new(CommonResult)
				result404.Data = make([]*models.Card, 0)
				result404.Success = false
				result404.Message = "404"
				return JSON(200, result404)
			}
			return Error(500, "Failed to get page", err)
		}
		query := models.GetCardsQuery{
			PageId: pageQuery.Result.Id,
			OrgId:  c.OrgId,
		}
		if err := bus.Dispatch(&query); err != nil {
			return Error(500, "Failed to get cards", err)
		}
		cards = query.Result
	}

	for _, cd := range cards {
		lQuery := models.GetLocaleQuery{Name: cd.Name, VId: cd.Id, Type: "card"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		cd.Locale = lQuery.Result
		fQuery := models.GetFieldsQuery{CardId: cd.Id}
		if err := bus.Dispatch(&fQuery); err != nil {
			return Error(500, "Failed to get fields", err)
		}
		cd.Fields = fQuery.Result
		for _, fd := range cd.Fields {
			lfQuery := models.GetLocaleQuery{Name: fd.Name, VId: fd.Id, Type: "field"}
			if err := bus.Dispatch(&lfQuery); err != nil {
				return Error(500, "Failed to get locale", err)
			}
			fd.Locale = lfQuery.Result
		}
		//get styles
		sQuery := models.GetStylesQuery{
			CardId: cd.Id,
			Style:  cd.Style,
		}
		if err := bus.Dispatch(&sQuery); err != nil {
			return Error(500, "Failed to get styles", err)
		}
		cd.Styles = sQuery.Result
		//get actions of card
		aQuery := models.GetViewActionsQuery{CardId: cd.Id}
		if err := bus.Dispatch(&aQuery); err != nil {
			return Error(500, "Failed to get actions", err)
		}
		cd.Actions = aQuery.Result
		for _, act := range cd.Actions {
			lfQuery := models.GetLocaleQuery{Name: act.Name, VId: act.Id, Type: "action"}
			if err := bus.Dispatch(&lfQuery); err != nil {
				return Error(500, "Failed to get locale", err)
			}
			act.Locale = lfQuery.Result
			paraQuery := models.GetParametersQuery{ActionId: act.Id}
			if err := bus.Dispatch(&paraQuery); err != nil {
				return Error(500, "Failed to get parameters 1", err)
			}
			act.Parameters = paraQuery.Result
		}

		//append reader
		reader := models.ViewAction{
			CardId:      cd.Id,
			Name:        "Read",
			Type:        "READ",
			Seq:         0,
			DoubleCheck: false,
		}
		reader.Locale = models.Locale{}
		reader.Parameters = make([]*models.Parameter, 0)
		for _, fd := range cd.Fields {
			if fd.Filter == "EQUAL" {
				param := models.Parameter{
					FieldId:    fd.Id,
					IsVisible:  fd.IsVisible,
					IsEditable: fd.IsVisible,
					Default:    fd.Default,
					Compare:    "eq",
					Field:      fd,
				}
				reader.Parameters = append(reader.Parameters, &param)
			} else if fd.Filter == "MAX" {
				param := models.Parameter{
					FieldId:    fd.Id,
					IsVisible:  fd.IsVisible,
					IsEditable: fd.IsVisible,
					Default:    fd.Default,
					Compare:    "le",
					Field:      fd,
				}
				reader.Parameters = append(reader.Parameters, &param)
			} else if fd.Filter == "MIN" {
				param := models.Parameter{
					FieldId:    fd.Id,
					IsVisible:  fd.IsVisible,
					IsEditable: fd.IsVisible,
					Default:    fd.Default,
					Compare:    "ge",
					Field:      fd,
				}
				reader.Parameters = append(reader.Parameters, &param)
			} else if fd.Filter == "RANGE" {
				param2 := models.Parameter{
					FieldId:    fd.Id,
					IsVisible:  fd.IsVisible,
					IsEditable: fd.IsVisible,
					Default:    fd.Default,
					Compare:    "ge",
					Field:      fd,
				}
				reader.Parameters = append(reader.Parameters, &param2)
				param := models.Parameter{
					FieldId:    fd.Id,
					IsVisible:  fd.IsVisible,
					IsEditable: fd.IsVisible,
					Default:    fd.Default,
					Compare:    "le",
					Field:      fd,
				}
				reader.Parameters = append(reader.Parameters, &param)
			}
		}
		cd.Actions = append(cd.Actions, &reader)
	}
	result := new(CommonResult)
	result.Data = cards
	result.Success = true
	return JSON(200, result)
}

//GetCardByID 获取卡片
func GetCardByID(c *models.ReqContext, form dtos.GetCardByIdForm) Response {
	query := models.GetCardQuery{
		CardId: form.CardId,
		OrgId:  c.OrgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get card", err)
	}
	lQuery := models.GetLocaleQuery{Name: query.Result.Name, VId: query.Result.Id, Type: "card"}
	if err := bus.Dispatch(&lQuery); err != nil {
		return Error(500, "Failed to get locale", err)
	}
	query.Result.Locale = lQuery.Result
	fQuery := models.GetFieldsQuery{CardId: query.Result.Id}
	if err := bus.Dispatch(&fQuery); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	for _, fd := range fQuery.Result {
		lfQuery := models.GetLocaleQuery{Name: fd.Name, VId: fd.Id, Type: "field"}
		if err := bus.Dispatch(&lfQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		fd.Locale = lfQuery.Result
	}
	query.Result.Fields = fQuery.Result

	//get styles
	sQuery := models.GetStylesQuery{
		CardId: query.Result.Id,
		Style:  query.Result.Style,
	}
	if err := bus.Dispatch(&sQuery); err != nil {
		return Error(500, "Failed to get styles", err)
	}
	query.Result.Styles = sQuery.Result

	aQuery := models.GetViewActionsQuery{CardId: query.Result.Id}
	if err := bus.Dispatch(&aQuery); err != nil {
		return Error(500, "Failed to get view actions", err)
	}
	for _, act := range aQuery.Result {
		lfQuery := models.GetLocaleQuery{Name: act.Name, VId: act.Id, Type: "action"}
		if err := bus.Dispatch(&lfQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		act.Locale = lfQuery.Result

		paraQuery := models.GetParametersQuery{ActionId: act.Id}
		if err := bus.Dispatch(&paraQuery); err != nil {
			return Error(500, "Failed to get parameters 2", err)
		}
		act.Parameters = paraQuery.Result
	}
	query.Result.Actions = aQuery.Result
	//append reader
	reader := models.ViewAction{
		CardId:      form.CardId,
		Name:        "Read",
		Type:        "READ",
		Seq:         0,
		DoubleCheck: false,
	}
	reader.Locale = models.Locale{}
	reader.Parameters = make([]*models.Parameter, 0)
	for _, fd := range query.Result.Fields {
		if fd.Filter == "EQUAL" {
			param := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "eq",
				Field:      fd,
			}
			reader.Parameters = append(reader.Parameters, &param)
		} else if fd.Filter == "MAX" {
			param := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "le",
				Field:      fd,
			}
			reader.Parameters = append(reader.Parameters, &param)
		} else if fd.Filter == "MIN" {
			param := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "ge",
				Field:      fd,
			}
			reader.Parameters = append(reader.Parameters, &param)
		} else if fd.Filter == "RANGE" {
			param2 := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "ge",
				Field:      fd,
			}
			reader.Parameters = append(reader.Parameters, &param2)
			param := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "le",
				Field:      fd,
			}
			reader.Parameters = append(reader.Parameters, &param)
		}
	}
	query.Result.Actions = append(query.Result.Actions, &reader)
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//CreateCard 添加卡片
func CreateCard(c *models.ReqContext, form dtos.CreateCardForm, lang dtos.LocaleForm) Response {
	//page
	page := new(models.Page)
	if form.PageId != 0 {
		pageQuery := models.GetPageByIDQuery{
			PageId: form.PageId,
			OrgId:  c.OrgId,
		}
		if err := bus.Dispatch(&pageQuery); err != nil {
			return Error(500, "Failed to get cards", err)
		}
		page = pageQuery.Result
	} else {
		pageQuery := models.GetPageByNameQuery{
			PageName: form.PageName,
			OrgId:    c.OrgId,
		}
		if err := bus.Dispatch(&pageQuery); err != nil {
			return Error(500, "Failed to get page", err)
		}
		page = pageQuery.Result
	}

	query := models.GetCardsQuery{
		PageId: page.Id,
		OrgId:  c.OrgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get cards", err)
	}
	seq := int32(1)
	if len(query.Result) > 0 {
		seq = query.Result[len(query.Result)-1].Seq + 2
	}
	cmd := models.CreateCardCommand{
		PageId: page.Id,
		Name:   form.Name,
		Style:  "TABLE",
		Seq:    seq,
		Pos:    0,
		Width:  12,
		OrgId:  c.OrgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to create  card", err)
	}
	localeCmd := models.SetLocaleCommand{
		VId:      cmd.Result,
		Type:     "card",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)
	result := new(CommonResult)
	result.Data = query.Result
	result.Success = true
	return JSON(200, result)
}

//UpdateCard 修改卡片信息
func UpdateCard(c *models.ReqContext, form dtos.UpdateCardForm, lang dtos.LocaleForm) Response {

	localeCmd := models.SetLocaleCommand{
		VId:      form.CardId,
		Type:     "card",
		Text:     form.Text,
		Language: lang.Language,
	}
	bus.Dispatch(&localeCmd)
	cmd := models.UpdateCardCommand{
		CardId: form.CardId,
		Name:   form.Name,
		Style:  form.Style,
		Width:  form.Width,
		Pos:    form.Pos,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to update card", err)
	}

	lQuery := models.GetLocaleQuery{Name: cmd.Result.Name, VId: cmd.Result.Id, Type: "card"}
	if err := bus.Dispatch(&lQuery); err != nil {
		return Error(500, "Failed to get locale", err)
	}
	cmd.Result.Locale = lQuery.Result

	//get styles
	sQuery := models.GetStylesQuery{
		CardId: form.CardId,
		Style:  form.Style,
	}
	if err := bus.Dispatch(&sQuery); err != nil {
		return Error(500, "Failed to get styles", err)
	}
	cmd.Result.Styles = sQuery.Result

	//get fields
	fQuery := models.GetFieldsQuery{CardId: form.CardId}
	if err := bus.Dispatch(&fQuery); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	for _, fd := range fQuery.Result {
		lfQuery := models.GetLocaleQuery{Name: fd.Name, VId: fd.Id, Type: "field"}
		if err := bus.Dispatch(&lfQuery); err != nil {
			return Error(500, "Failed to get locale", err)
		}
		fd.Locale = lfQuery.Result
	}
	cmd.Result.Fields = fQuery.Result

	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

//UpdateCardSeq 修改卡片顺序
func UpdateCardSeq(c *models.ReqContext, form dtos.UpdateCardSeqForm) Response {
	if form.Source == 0 || form.Target == 0 {
		result := new(CommonResult)
		result.Success = false
		return JSON(200, result)
	}
	squery := models.GetCardQuery{
		CardId: form.Source,
		OrgId:  c.OrgId,
	}
	if err := bus.Dispatch(&squery); err != nil {
		return Error(500, "Failed to get card", err)
	}
	tquery := models.GetCardQuery{
		CardId: form.Target,
		OrgId:  c.OrgId,
	}
	if err := bus.Dispatch(&tquery); err != nil {
		return Error(500, "Failed to get card", err)
	}

	cmd := models.UpdateCardSeqCommand{
		CardId: squery.Result.Id,
		Seq:    tquery.Result.Seq + form.Position*2 - 3,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to update card", err)
	}
	query := models.GetCardsQuery{
		PageId: squery.Result.PageId,
		OrgId:  c.OrgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get cards", err)
	}
	for idx, cd := range query.Result {
		if cd.Seq != int32(idx*2+1) {
			cd.Seq = int32(idx*2 + 1)
			seqCmd := models.UpdateCardSeqCommand{
				CardId: cd.Id,
				Seq:    cd.Seq,
			}
			if err := bus.Dispatch(&seqCmd); err != nil {
				return Error(500, "Failed to update card seq", err)
			}
		}
	}

	for _, cd := range query.Result {
		lQuery := models.GetLocaleQuery{Name: cd.Name, VId: cd.Id, Type: "card"}
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

//DeleteCard 删除卡片
func DeleteCard(c *models.ReqContext, form dtos.DeleteCardForm) Response {
	cmd := models.DeleteCardCommand{
		CardId: form.CardId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to delete card", err)
	}
	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}
