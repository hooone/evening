package card

import (
	"sort"
	"strings"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/card"
	"github.com/hooone/evening/pkg/managers/page"
	"github.com/hooone/evening/pkg/services/action"
	"github.com/hooone/evening/pkg/services/field"
	"github.com/hooone/evening/pkg/services/locale"
	"github.com/hooone/evening/pkg/services/parameter"
	"github.com/hooone/evening/pkg/services/style"

	"github.com/hooone/evening/pkg/registry"
)

func init() {
	registry.RegisterService(&CardService{})
}

type CardService struct {
	Locale        *locale.LocaleService `inject:""`
	FieldService  *field.FieldService   `inject:""`
	StyleService  *style.StyleService   `inject:""`
	ActionService *action.ActionService `inject:""`
}

func (s *CardService) Init() error {
	return nil
}

//GetCards 获取页面中的所有窗体
func (s *CardService) GetCards(pageId int64, pageName string, orgId int64, lang string) ([]*Card, error) {
	//get cards
	cards := make([]*Card, 0)
	if pageId != 0 {
		query := card.GetCardsQuery{
			PageId: pageId,
			OrgId:  orgId,
		}
		if err := bus.Dispatch(&query); err != nil {
			return cards, err
		}
		for _, cd := range query.Result {
			cad := Card{
				Id:     cd.Id,
				PageId: cd.PageId,
				OrgId:  cd.OrgId,
				Name:   cd.Name,
				Seq:    cd.Seq,
				Pos:    cd.Pos,
				Width:  cd.Width,
				Style:  cd.Style,
			}
			cards = append(cards, &cad)
		}
	} else {
		pageQuery := page.GetPageByNameQuery{
			PageName: pageName,
			OrgId:    orgId,
		}
		if err := bus.Dispatch(&pageQuery); err != nil {
			return cards, err
		}
		query := card.GetCardsQuery{
			PageId: pageQuery.Result.Id,
			OrgId:  orgId,
		}
		if err := bus.Dispatch(&query); err != nil {
			return cards, err
		}
		for _, cd := range query.Result {
			cad := Card{
				Id:     cd.Id,
				PageId: cd.PageId,
				OrgId:  cd.OrgId,
				Name:   cd.Name,
				Seq:    cd.Seq,
				Pos:    cd.Pos,
				Width:  cd.Width,
				Style:  cd.Style,
			}
			cards = append(cards, &cad)
		}
	}
	//set
	for _, cd := range cards {
		//locale
		s.Locale.GetLocale(cd, lang)

		//fields
		fa, err1 := s.FieldService.GetFields(cd.Id, orgId, lang)
		if err1 != nil {
			return cards, err1
		}
		cd.Fields = fa
		//styles
		st, err2 := s.StyleService.GetStyles(cd.Id, cd.Style, orgId, lang)
		if err2 != nil {
			return cards, err2
		}
		cd.Styles = st
		//actions
		ac, err3 := s.ActionService.GetViewActions(cd.Id, orgId, lang)
		if err3 != nil {
			return cards, err2
		}
		cd.Actions = ac

		//append reader
		reader := getReader(cd)
		cd.Actions = append(cd.Actions, &reader)
	}

	//arrange seq
	sort.Sort(CardSlice(cards))
	for idx, cd := range cards {
		if cd.Seq != int32(idx*2+1) {
			cd.Seq = int32(idx*2 + 1)
			seqCmd := card.UpdateCardSeqCommand{
				CardId: cd.Id,
				Seq:    cd.Seq,
				OrgId:  orgId,
			}
			if err := bus.Dispatch(&seqCmd); err != nil {
				return cards, err
			}
		}
	}
	return cards, nil
}
func getReader(card *Card) action.ViewAction {
	reader := action.ViewAction{
		CardId:      card.Id,
		Name:        "Read",
		Type:        "READ",
		Seq:         0,
		DoubleCheck: false,
	}
	reader.Parameters = make([]*parameter.Parameter, 0)
	for _, fd := range card.Fields {
		if fd.Filter == "EQUAL" {
			param := parameter.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "eq",
				Field:      fd,
			}
			reader.Parameters = append(reader.Parameters, &param)
		} else if fd.Filter == "MAX" {
			param := parameter.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "le",
				Field:      fd,
			}
			reader.Parameters = append(reader.Parameters, &param)
		} else if fd.Filter == "MIN" {
			param := parameter.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "ge",
				Field:      fd,
			}
			reader.Parameters = append(reader.Parameters, &param)
		} else if fd.Filter == "RANGE" {
			dftStr := strings.Split(fd.Default, "||")

			param2 := parameter.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    dftStr[0],
				Compare:    "ge",
				Field:      fd,
			}
			dftStr2 := fd.Default
			if len(dftStr) > 1 {
				dftStr2 = dftStr[1]
			}
			reader.Parameters = append(reader.Parameters, &param2)
			param := parameter.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    dftStr2,
				Compare:    "le",
				Field:      fd,
			}
			reader.Parameters = append(reader.Parameters, &param)
		}
	}
	return reader
}

//GetCardByID 获取窗体
func (s *CardService) GetCardByID(cardId int64, orgId int64, lang string) (*Card, error) {
	query := card.GetCardQuery{
		CardId: cardId,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return &Card{}, err
	}
	cd := &Card{
		Id:     query.Result.Id,
		PageId: query.Result.PageId,
		OrgId:  query.Result.OrgId,
		Name:   query.Result.Name,
		Seq:    query.Result.Seq,
		Pos:    query.Result.Pos,
		Width:  query.Result.Width,
		Style:  query.Result.Style,
	}
	//locale
	s.Locale.GetLocale(cd, lang)

	//fields
	fa, err1 := s.FieldService.GetFields(cd.Id, orgId, lang)
	if err1 != nil {
		return cd, err1
	}
	cd.Fields = fa
	//styles
	st, err2 := s.StyleService.GetStyles(cd.Id, cd.Style, orgId, lang)
	if err2 != nil {
		return cd, err2
	}
	cd.Styles = st
	//actions
	ac, err3 := s.ActionService.GetViewActions(cd.Id, orgId, lang)
	if err3 != nil {
		return cd, err2
	}
	cd.Actions = ac

	//append reader
	reader := getReader(cd)
	cd.Actions = append(cd.Actions, &reader)
	return cd, nil
}

//CreateCard 添加窗体
func (s *CardService) CreateCard(pageId int64, pageName string, form Card, orgId int64, lang string) error {
	//page
	pageT := new(page.Page)
	if pageId != 0 {
		pageQuery := page.GetPageByIDQuery{
			PageId: pageId,
			OrgId:  orgId,
		}
		if err := bus.Dispatch(&pageQuery); err != nil {
			return err
		}
		pageT = pageQuery.Result
	} else {
		pageQuery := page.GetPageByNameQuery{
			PageName: pageName,
			OrgId:    orgId,
		}
		if err := bus.Dispatch(&pageQuery); err != nil {
			return err
		}
		pageT = pageQuery.Result
	}

	//create
	cmd := card.CreateCardCommand{
		PageId: pageT.Id,
		Name:   form.Name,
		Style:  "TABLE",
		Seq:    99999,
		Pos:    0,
		Width:  12,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	form.Id = cmd.Result
	//create button
	s.ActionService.CreateViewAction(action.ViewAction{
		CardId:      cmd.Result,
		Name:        "create",
		Text:        "添加",
		Type:        "CREATE",
		DoubleCheck: false,
	}, orgId, "zh-CN")
	s.ActionService.CreateViewAction(action.ViewAction{
		CardId:      cmd.Result,
		Name:        "update",
		Text:        "修改",
		Type:        "UPDATE",
		DoubleCheck: false,
	}, orgId, "zh-CN")
	s.ActionService.CreateViewAction(action.ViewAction{
		CardId:      cmd.Result,
		Name:        "delete",
		Text:        "删除",
		Type:        "DELETE",
		DoubleCheck: false,
	}, orgId, "zh-CN")
	//create style
	s.StyleService.CreateStyles(form.Id, orgId, lang)
	//locale
	s.Locale.SetLocale(form, lang)
	return nil
}

//UpdateCard 修改窗体信息
func (s *CardService) UpdateCard(form Card, orgId int64, lang string) error {
	cmd := card.UpdateCardCommand{
		OrgId:  orgId,
		CardId: form.Id,
		Name:   form.Name,
		Style:  form.Style,
		Width:  form.Width,
		Pos:    form.Pos,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}

	s.Locale.SetLocale(form, lang)
	return nil
}

//UpdateCardSeq 修改窗体顺序
func (s *CardService) UpdateCardSeq(Source int64, Target int64, Position int32, orgId int64, lang string) error {
	if Source == 0 || Target == 0 {
		return nil
	}
	squery := card.GetCardQuery{
		CardId: Source,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&squery); err != nil {
		return err
	}
	tquery := card.GetCardQuery{
		CardId: Target,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&tquery); err != nil {
		return err
	}

	cmd := card.UpdateCardSeqCommand{
		CardId: squery.Result.Id,
		Seq:    tquery.Result.Seq + Position*2 - 3,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	return nil
}

//DeleteCard 删除窗体
func (s *CardService) DeleteCard(form Card, orgId int64, lang string) error {
	cmd := card.DeleteCardCommand{
		CardId: form.Id,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}

	//delete locale
	s.Locale.DeleteLocale(form, lang)

	//delete action
	if acs, err := s.ActionService.GetViewActions(form.Id, orgId, lang); err == nil {
		for _, ac := range acs {
			s.ActionService.DeleteViewAction(action.ViewAction{
				Id:          ac.Id,
				CardId:      ac.CardId,
				Name:        ac.Name,
				Type:        ac.Type, //CRUD
				Text:        ac.Text,
				Seq:         ac.Seq,
				DoubleCheck: ac.DoubleCheck,
			}, orgId, lang)
		}
	}

	//delete style
	s.StyleService.DeleteStyles(form.Id, orgId, lang)

	//delete field
	if fds, err := s.FieldService.GetFields(form.Id, orgId, lang); err == nil {
		for _, fd := range fds {
			s.FieldService.DeleteField(field.Field{
				Id:        fd.Id,
				CardId:    fd.CardId,
				Name:      fd.Name,
				Text:      fd.Text,
				Seq:       fd.Seq,
				IsVisible: fd.IsVisible,
				Type:      fd.Type,
				Filter:    fd.Filter,
				Default:   fd.Default,
			}, orgId, lang)
		}
	}

	return nil
}
