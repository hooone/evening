package sqlstore

import (
	"errors"
	"sort"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetField)
	bus.AddHandler("sql", GetFields)
	bus.AddHandler("sql", CreateField)
	bus.AddHandler("sql", UpdateField)
	bus.AddHandler("sql", UpdateFieldSeq)
	bus.AddHandler("sql", DeleteField)
}

//GetField 获得字段
func GetField(query *models.GetFieldQuery) error {
	field := models.Field{Id: query.FieldId}
	success, err := x.Id(query.FieldId).Get(&field)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("now field row found")
	}
	query.Result = &field
	return nil
}

//GetFields 获得页面所有字段
func GetFields(query *models.GetFieldsQuery) error {
	fields := make([]*models.Field, 0)
	err := x.Where("card_id = ?", query.CardId).Find(&fields)
	if err != nil {
		return err
	}
	sort.Sort(models.FieldSlice(fields))
	query.Result = fields
	return nil
}

//CreateField 添加字段
func CreateField(cmd *models.CreateFieldCommand) error {
	field := models.Field{
		CardId:    cmd.CardId,
		Name:      cmd.Name,
		Seq:       cmd.Seq,
		IsVisible: cmd.IsVisible,
		Type:      cmd.Type,
		Default:   cmd.Default,
		Filter:    cmd.Filter,
	}
	_, err := x.Insert(&field)
	if err != nil {
		return err
	}
	cmd.Result = field.Id
	return nil
}

//UpdateField 修改字段
func UpdateField(cmd *models.UpdateFieldCommand) error {
	field := models.Field{
		Id:        cmd.FieldId,
		Name:      cmd.Name,
		IsVisible: cmd.IsVisible,
		Type:      cmd.Type,
		Default:   cmd.Default,
		Filter:    cmd.Filter,
	}
	_, err := x.Id(cmd.FieldId).Cols("name").Cols("is_visible").Cols("default").Cols("type").Cols("filter").Update(&field)
	if err != nil {
		return err
	}
	success, err2 := x.Id(field.Id).Get(&field)
	if err2 != nil {
		return err
	}
	if !success {
		return errors.New("no field found")
	}
	cmd.Result = &field
	return nil
}

//UpdateFieldSeq 修改字段顺序
func UpdateFieldSeq(cmd *models.UpdateFieldSeqCommand) error {
	field := models.Field{
		Id:  cmd.FieldId,
		Seq: cmd.Seq,
	}
	_, err := x.Id(cmd.FieldId).Cols("seq").Update(&field)
	if err != nil {
		return err
	}
	success, err2 := x.Id(cmd.FieldId).Get(&field)
	if err2 != nil {
		return err2
	}
	if !success {
		return errors.New("now field row found")
	}
	cmd.Result = &field
	return nil
}

//DeleteField 删除字段
func DeleteField(cmd *models.DeleteFieldCommand) error {
	field := models.Field{
		Id: cmd.FieldId,
	}
	success, err := x.Id(cmd.FieldId).Get(&field)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("no row found")
	}
	_, err2 := x.Id(cmd.FieldId).Delete(&field)
	if err2 != nil {
		return err2
	}

	paraCmd := models.DeleteParameterByFieldCommand{FieldId: field.Id}
	if err3 := bus.Dispatch(&paraCmd); err != nil {
		return err3
	}
	cmd.Result = &field
	return nil
}
