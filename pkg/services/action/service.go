package action

import (
	"sort"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/action"
	"github.com/hooone/evening/pkg/managers/field"
	paramMng "github.com/hooone/evening/pkg/managers/parameter"
	"github.com/hooone/evening/pkg/services/locale"
	"github.com/hooone/evening/pkg/services/parameter"

	"github.com/hooone/evening/pkg/registry"
)

func init() {
	registry.RegisterService(&ActionService{})
}

type ActionService struct {
	Locale           *locale.LocaleService       `inject:""`
	ParameterService *parameter.ParameterService `inject:""`
}

func (s *ActionService) Init() error {
	return nil
}

//GetViewActions 获取页面中的所有操作
func (s *ActionService) GetViewActions(cardId int64, orgId int64, lang string) ([]*ViewAction, error) {
	result := make([]*ViewAction, 0)
	query := action.GetViewActionsQuery{
		CardId: cardId,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return result, err
	}
	for _, v := range query.Result {
		va := ViewAction{
			Id:          v.Id,
			CardId:      v.CardId,
			Name:        v.Name,
			Type:        v.Type, //CRUD
			Seq:         v.Seq,
			DoubleCheck: v.DoubleCheck,
		}
		pa, err := s.ParameterService.GetParameters(v.Id, orgId, lang)
		if err != nil {
			return result, err
		}
		va.Parameters = pa
		s.Locale.GetLocale(&va, lang)
		result = append(result, &va)
	}

	sort.Sort(ViewActionSlice(result))
	//arrange folder seq
	for idx, va := range result {
		if va.Seq != int32(idx*2+1) {
			va.Seq = int32(idx*2 + 1)
			seqCmd := action.UpdateViewActionSeqCommand{
				ViewActionId: va.Id,
				Seq:          va.Seq,
				OrgId:        orgId,
			}
			if err := bus.Dispatch(&seqCmd); err != nil {
				return result, err
			}
		}
	}
	return result, nil
}

//GetViewActionByID 获取操作
func (s *ActionService) GetViewActionByID(actionId int64, orgId int64, lang string) (ViewAction, error) {
	query := action.GetViewActionQuery{
		ViewActionId: actionId,
		OrgId:        orgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return ViewAction{}, err
	}
	va := ViewAction{
		Id:          query.Result.Id,
		CardId:      query.Result.CardId,
		Name:        query.Result.Name,
		Type:        query.Result.Type, //CRUD
		Seq:         query.Result.Seq,
		DoubleCheck: query.Result.DoubleCheck,
	}
	pa, err := s.ParameterService.GetParameters(va.Id, orgId, lang)
	if err != nil {
		return ViewAction{}, err
	}
	va.Parameters = pa
	s.Locale.GetLocale(&va, lang)
	return va, nil
}

//CreateViewAction 添加操作
func (s *ActionService) CreateViewAction(form ViewAction, orgId int64, lang string) error {
	//calculate seq
	query := action.GetViewActionsQuery{
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

	//create action
	cmd := action.CreateViewActionCommand{
		CardId:      form.CardId,
		OrgId:       orgId,
		Name:        form.Name,
		Seq:         seq,
		DoubleCheck: form.DoubleCheck,
		Type:        form.Type,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	form.Id = cmd.Result
	//create parameter
	fquery := field.GetFieldsQuery{
		CardId: form.CardId,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&fquery); err != nil {
		return err
	}
	show := true
	if form.Type == "DELETE" {
		show = false
	}
	for _, field := range fquery.Result {
		paCmd := paramMng.CreateParameterCommand{
			ActionId:   cmd.Result,
			OrgId:      orgId,
			FieldId:    field.Id,
			IsVisible:  show,
			IsEditable: show,
			Default:    field.Default,
		}
		if err2 := bus.Dispatch(&paCmd); err2 != nil {
			return err2
		}
	}

	//set locale
	s.Locale.SetLocale(form, lang)

	return nil
}

//UpdateViewAction 修改操作信息
func (s *ActionService) UpdateViewAction(form ViewAction, orgId int64, lang string) error {
	cmd := action.UpdateViewActionCommand{
		ViewActionId: form.Id,
		OrgId:        orgId,
		Name:         form.Name,
		DoubleCheck:  form.DoubleCheck,
		Type:         form.Type,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	s.Locale.SetLocale(form, lang)
	return nil
}

//DeleteViewAction 删除操作
func (s *ActionService) DeleteViewAction(form ViewAction, orgId int64, lang string) error {
	//delete
	cmd := action.DeleteViewActionCommand{
		ViewActionId: form.Id,
		OrgId:        orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	//delete locale
	s.Locale.DeleteLocale(form, lang)
	//delete parameter
	pCmd := paramMng.DeleteParameterByActionCommand{
		OrgId:    orgId,
		ActionId: form.Id,
	}
	if err := bus.Dispatch(&pCmd); err != nil {
		return err
	}
	return nil
}

//UpdateViewActionSeq 修改操作顺序
func (s *ActionService) UpdateViewActionSeq(Source int64, Target int64, Position int32, orgId int64) error {
	if Source == 0 || Target == 0 {
		return nil
	}
	squery := action.GetViewActionQuery{
		ViewActionId: Source,
		OrgId:        orgId,
	}
	if err := bus.Dispatch(&squery); err != nil {
		return err
	}
	tquery := action.GetViewActionQuery{
		ViewActionId: Target,
		OrgId:        orgId,
	}
	if err := bus.Dispatch(&tquery); err != nil {
		return err
	}

	cmd := action.UpdateViewActionSeqCommand{
		ViewActionId: squery.Result.Id,
		OrgId:        orgId,
		Seq:          tquery.Result.Seq + Position*2 - 3,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}

	return nil
}
