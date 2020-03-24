package sqlstore

import (
	"errors"
	"sort"
	"strconv"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetParameter)
	bus.AddHandler("sql", GetParameters)
	bus.AddHandler("sql", CreateParameter)
	bus.AddHandler("sql", UpdateParameter)
	bus.AddHandler("sql", DeleteParameterByField)
	bus.AddHandler("sql", DeleteParameterByAction)
}

//GetParameter 获得参数
func GetParameter(query *models.GetParameterQuery) error {
	parameter := models.Parameter{Id: query.ParameterId}
	success, err := x.Id(query.ParameterId).Get(&parameter)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("now parameter row found")
	}
	success, err = x.Table("field").Id(parameter.FieldId).Get(parameter.Field)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("can not find relation field " + strconv.Itoa(int(parameter.FieldId)))
	}
	query.Result = &parameter
	return nil
}

//GetParameters 获得操作的所有字段
func GetParameters(query *models.GetParametersQuery) error {
	parameters := make([]*models.Parameter, 0)
	err := x.Where("action_id = ?", query.ActionId).Find(&parameters)
	if err != nil {
		return err
	}
	for _, p := range parameters {
		p.Field = &models.Field{}
		success, err2 := x.Table("field").Id(p.FieldId).Get(p.Field)
		if err2 != nil {
			return err2
		}
		if !success {
			return errors.New("can not find relation field " + strconv.Itoa(int(p.FieldId)))
		}
		lQuery := models.GetLocaleQuery{Name: p.Field.Name, VId: p.Field.Id, Type: "field"}
		if err := bus.Dispatch(&lQuery); err != nil {
			return err
		}
		p.Field.Locale = lQuery.Result
	}
	sort.Sort(models.ParameterSlice(parameters))
	query.Result = parameters
	return nil
}

//CreateParameter 添加参数
func CreateParameter(cmd *models.CreateParameterCommand) error {
	parameter := models.Parameter{
		ActionId:   cmd.ActionId,
		FieldId:    cmd.FieldId,
		IsVisible:  cmd.IsVisible,
		IsEditable: cmd.IsEditable,
		Default:    cmd.Default,
	}
	_, err := x.Insert(&parameter)
	if err != nil {
		return err
	}
	cmd.Result = parameter.Id
	return nil
}

//UpdateParameter 修改参数
func UpdateParameter(cmd *models.UpdateParameterCommand) error {
	parameter := models.Parameter{
		Id:         cmd.ParameterId,
		IsVisible:  cmd.IsVisible,
		IsEditable: cmd.IsEditable,
		Default:    cmd.Default,
	}
	_, err := x.Id(cmd.ParameterId).Cols("is_visible").Cols("is_editable").Cols("default").Update(&parameter)
	if err != nil {
		return err
	}
	x.Id(parameter.Id).Get(&parameter)
	cmd.Result = &parameter
	return nil
}

//DeleteParameterByField 删除参数
func DeleteParameterByField(cmd *models.DeleteParameterByFieldCommand) error {
	count, err := x.Where("field_id = ?", cmd.FieldId).Delete(new(models.Parameter))
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}

//DeleteParameterByAction 删除参数
func DeleteParameterByAction(cmd *models.DeleteParameterByActionCommand) error {
	count, err := x.Where("action_id = ?", cmd.ActionId).Delete(new(models.Parameter))
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}
