package parameter

import (
	"sort"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/parameter"
	"github.com/hooone/evening/pkg/services/field"
	"github.com/hooone/evening/pkg/services/locale"

	"github.com/hooone/evening/pkg/registry"
)

func init() {
	registry.RegisterService(&ParameterService{})
}

type ParameterService struct {
	Locale       *locale.LocaleService `inject:""`
	FieldService *field.FieldService   `inject:""`
}

func (s *ParameterService) Init() error {
	return nil
}

//GetParameters 查找参数
func (s *ParameterService) GetParameters(actionId int64, orgId int64, lang string) ([]*Parameter, error) {
	result := make([]*Parameter, 0)
	query := parameter.GetParametersQuery{
		ActionId: actionId,
		OrgId:    orgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return result, err
	}
	for _, p := range query.Result {
		pa := Parameter{
			Id:         p.Id,
			ActionId:   p.ActionId,
			FieldId:    p.FieldId,
			IsVisible:  p.IsVisible,
			IsEditable: p.IsEditable,
			Default:    p.Default,
		}
		fd, err := s.FieldService.GetFieldByID(pa.FieldId, orgId, lang)
		if err != nil {
			return result, err
		}
		pa.Field = &fd
		result = append(result, &pa)
	}
	sort.Sort(ParameterSlice(result))
	return result, nil
}

//UpdateParameter 修改参数
func (s *ParameterService) UpdateParameter(form Parameter, orgId int64) error {
	cmd := parameter.UpdateParameterCommand{
		OrgId:       orgId,
		ParameterId: form.Id,
		IsVisible:   form.IsVisible,
		IsEditable:  form.IsEditable,
		Default:     form.Default,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	return nil
}
