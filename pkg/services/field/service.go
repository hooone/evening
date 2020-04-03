package field

import (
	"sort"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/action"
	"github.com/hooone/evening/pkg/managers/field"
	"github.com/hooone/evening/pkg/managers/parameter"
	"github.com/hooone/evening/pkg/services/locale"

	"github.com/hooone/evening/pkg/registry"
)

func init() {
	registry.RegisterService(&FieldService{})
}

type FieldService struct {
	Locale *locale.LocaleService `inject:""`
}

func (s *FieldService) Init() error {
	return nil
}

//GetFields 获取页面中的所有字段
func (s *FieldService) GetFields(cardId int64, orgId int64, lang string) ([]*Field, error) {
	result := make([]*Field, 0)
	query := field.GetFieldsQuery{
		CardId: cardId,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return result, err
	}
	for _, f := range query.Result {
		fl := Field{
			Id:        f.Id,
			CardId:    f.CardId,
			Name:      f.Name,
			Seq:       f.Seq,
			IsVisible: f.IsVisible,
			Type:      f.Type,
			Filter:    f.Filter,
			Default:   f.Default,
		}
		s.Locale.GetLocale(&fl, lang)
		result = append(result, &fl)
	}
	sort.Sort(FieldSlice(result))

	//arrange seq
	for idx, cd := range result {
		if cd.Seq != int32(idx*2+1) {
			cd.Seq = int32(idx*2 + 1)
			seqCmd := field.UpdateFieldSeqCommand{
				FieldId: cd.Id,
				OrgId:   orgId,
				Seq:     cd.Seq,
			}
			if err := bus.Dispatch(&seqCmd); err != nil {
				return result, err
			}
		}
	}

	return result, nil
}

//GetFieldByID 获取字段
func (s *FieldService) GetFieldByID(FieldId int64, orgId int64, lang string) (Field, error) {
	query := field.GetFieldQuery{
		FieldId: FieldId,
		OrgId:   orgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return Field{}, err
	}
	fl := Field{
		Id:        query.Result.Id,
		CardId:    query.Result.CardId,
		Name:      query.Result.Name,
		Seq:       query.Result.Seq,
		IsVisible: query.Result.IsVisible,
		Type:      query.Result.Type,
		Filter:    query.Result.Filter,
		Default:   query.Result.Default,
	}
	s.Locale.GetLocale(&fl, lang)
	return fl, nil
}

//CreateField 添加字段
func (s *FieldService) CreateField(form Field, orgId int64, lang string) error {
	//calculate seq
	query := field.GetFieldsQuery{
		CardId: form.CardId,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return err
	}
	seq := int32(1)
	for _, temp := range query.Result {
		if temp.Seq > seq {
			seq = temp.Seq
		}
	}
	seq = seq + 2

	//create field
	cmd := field.CreateFieldCommand{
		OrgId:     orgId,
		CardId:    form.CardId,
		Name:      form.Name,
		Seq:       seq,
		IsVisible: form.IsVisible,
		Type:      form.Type,
		Default:   form.Default,
		Filter:    form.Filter,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	form.Id = cmd.Result

	//create parameter
	aquery := action.GetViewActionsQuery{
		CardId: form.CardId,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&aquery); err != nil {
		return err
	}
	for _, act := range aquery.Result {
		paCmd := parameter.CreateParameterCommand{
			OrgId:      orgId,
			ActionId:   act.Id,
			FieldId:    cmd.Result,
			IsVisible:  true,
			IsEditable: true,
			Default:    form.Default,
		}
		if err2 := bus.Dispatch(&paCmd); err2 != nil {
			return err2
		}
	}

	// set locale
	s.Locale.SetLocale(form, lang)

	return nil
}

//UpdateField 修改字段信息
func (s *FieldService) UpdateField(form Field, orgId int64, lang string) error {
	//update field
	cmd := field.UpdateFieldCommand{
		FieldId:   form.Id,
		OrgId:     orgId,
		Name:      form.Name,
		IsVisible: form.IsVisible,
		Type:      form.Type,
		Default:   form.Default,
		Filter:    form.Filter,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}

	//set locale
	s.Locale.SetLocale(form, lang)

	return nil
}

//UpdateFieldSeq 修改字段顺序
func (s *FieldService) UpdateFieldSeq(Source int64, Target int64, Position int32, orgId int64, lang string) error {
	if Source == 0 || Target == 0 {
		return nil
	}
	//get fields
	squery := field.GetFieldQuery{
		FieldId: Source,
		OrgId:   orgId,
	}
	if err := bus.Dispatch(&squery); err != nil {
		return err
	}
	tquery := field.GetFieldQuery{
		FieldId: Target,
		OrgId:   orgId,
	}
	if err := bus.Dispatch(&tquery); err != nil {
		return err
	}

	//update seq
	cmd := field.UpdateFieldSeqCommand{
		FieldId: squery.Result.Id,
		OrgId:   orgId,
		Seq:     tquery.Result.Seq + Position*2 - 3,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}

	return nil
}

//DeleteField 删除字段
func (s *FieldService) DeleteField(form Field, orgId int64, lang string) error {
	//delete field
	cmd := field.DeleteFieldCommand{
		FieldId: form.Id,
		OrgId:   orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}

	//delete locale
	s.Locale.DeleteLocale(form, lang)

	//delete parameter
	pCmd := parameter.DeleteParameterByFieldCommand{
		OrgId:   orgId,
		FieldId: form.Id,
	}
	if err := bus.Dispatch(&pCmd); err != nil {
		return err
	}
	return nil
}
