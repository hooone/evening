package parameter

import (
	"errors"
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetParameter)
	bus.AddHandler("sql", GetParameters)
	bus.AddHandler("sql", CreateParameter)
	bus.AddHandler("sql", UpdateParameter)
	bus.AddHandler("sql", DeleteParameterByField)
	bus.AddHandler("sql", DeleteParameterByAction)
}

var (
	ErrParameterNotFound = errors.New("parameter data not found")
)

//GetParameter 获得参数
func GetParameter(query *GetParameterQuery) error {
	parameter := Parameter{
		Id:    query.ParameterId,
		OrgId: query.OrgId,
	}
	success, err := sqlstore.X.Where("id=? and org_id=?", query.ParameterId, query.OrgId).Get(&parameter)
	if err != nil {
		return err
	}
	if !success {
		return ErrParameterNotFound
	}
	query.Result = &parameter
	return nil
}

//GetParameters 获得操作的所有字段
func GetParameters(query *GetParametersQuery) error {
	parameters := make([]*Parameter, 0)
	err := sqlstore.X.Where("action_id = ? and org_id=?", query.ActionId, query.OrgId).Find(&parameters)
	if err != nil {
		return err
	}
	query.Result = parameters
	return nil
}

//CreateParameter 添加参数
func CreateParameter(cmd *CreateParameterCommand) error {
	parameter := Parameter{
		ActionId:   cmd.ActionId,
		OrgId:      cmd.OrgId,
		FieldId:    cmd.FieldId,
		IsVisible:  cmd.IsVisible,
		IsEditable: cmd.IsEditable,
		Default:    cmd.Default,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}
	_, err := sqlstore.X.Insert(&parameter)
	if err != nil {
		return err
	}
	cmd.Result = parameter.Id
	return nil
}

//UpdateParameter 修改参数
func UpdateParameter(cmd *UpdateParameterCommand) error {
	parameter := Parameter{
		Id:         cmd.ParameterId,
		OrgId:      cmd.OrgId,
		IsVisible:  cmd.IsVisible,
		IsEditable: cmd.IsEditable,
		Default:    cmd.Default,
		UpdateAt:   time.Now(),
	}
	_, err := sqlstore.X.Where("id = ? and org_id=?", cmd.ParameterId, cmd.OrgId).Cols("is_visible").Cols("is_editable").Cols("default").Cols("update_at").Update(&parameter)
	if err != nil {
		return err
	}
	sqlstore.X.Id(parameter.Id).Get(&parameter)
	cmd.Result = &parameter
	return nil
}

//DeleteParameterByField 删除参数
func DeleteParameterByField(cmd *DeleteParameterByFieldCommand) error {
	count, err := sqlstore.X.Where("field_id = ? and org_id=?", cmd.FieldId, cmd.OrgId).Delete(new(Parameter))
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}

//DeleteParameterByAction 删除参数
func DeleteParameterByAction(cmd *DeleteParameterByActionCommand) error {
	count, err := sqlstore.X.Where("action_id = ? and org_id=?", cmd.ActionId, cmd.OrgId).Delete(new(Parameter))
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}
