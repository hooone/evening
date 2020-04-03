package style

import (
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/style"
	"github.com/hooone/evening/pkg/services/field"
	"github.com/hooone/evening/pkg/services/locale"

	"github.com/hooone/evening/pkg/registry"
)

func init() {
	registry.RegisterService(&StyleService{})
}

type StyleService struct {
	Locale       *locale.LocaleService `inject:""`
	FieldService *field.FieldService   `inject:""`
}

func (s *StyleService) Init() error {
	return nil
}

//UpdateStyles 新建Style
func (s *StyleService) CreateStyles(cardId int64, orgId int64, lang string) error {
	cmd := style.CreateStylesCommand{
		OrgId:  orgId,
		CardId: cardId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	return nil
}

//UpdateStyles 新建Style
func (s *StyleService) GetStyles(cardId int64, styleName string, orgId int64, lang string) ([]*Style, error) {
	result := make([]*Style, 0)
	query := style.GetStylesQuery{
		OrgId:  orgId,
		CardId: cardId,
		Style:  styleName,
	}
	if err := bus.Dispatch(&query); err != nil {
		return result, err
	}
	for _, st := range query.Result {
		sty := Style{
			Id:         st.Id,
			CardId:     st.CardId,
			FieldId:    st.FieldId,
			Type:       st.Type,
			Property:   st.Property,
			Value:      st.Value,
			Relation:   st.Relation,
			MustNumber: st.MustNumber,
		}
		if sty.Relation && sty.FieldId != 0 {
			fd, err := s.FieldService.GetFieldByID(sty.FieldId, orgId, lang)
			if err != nil {
				return result, err
			}
			sty.Field = &fd
		}
		result = append(result, &sty)
	}
	return result, nil
}

//UpdateStyles 更新Style
func (s *StyleService) UpdateStyles(form Style, orgId int64, lang string) error {
	cmd := style.UpdateStyleCommand{
		Id:      form.Id,
		OrgId:   orgId,
		FieldId: form.FieldId,
		Value:   form.Value,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	return nil
}

//UpdateStyles 更新Style
func (s *StyleService) DeleteStyles(cardId int64, orgId int64, lang string) error {
	cmd := style.DeleteStylesCommand{
		CardId: cardId,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	return nil
}
