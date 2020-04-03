package field

import (
	"errors"
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetField)
	bus.AddHandler("sql", GetFields)
	bus.AddHandler("sql", CreateField)
	bus.AddHandler("sql", UpdateField)
	bus.AddHandler("sql", UpdateFieldSeq)
	bus.AddHandler("sql", DeleteField)
}

var (
	ErrFieldNotFound = errors.New("field data not found")
)

//GetField 获得字段
func GetField(query *GetFieldQuery) error {
	field := Field{Id: query.FieldId}
	success, err := sqlstore.X.Where("id=? and org_id=?", query.FieldId, query.OrgId).Get(&field)
	if err != nil {
		return err
	}
	if !success {
		return ErrFieldNotFound
	}
	query.Result = &field
	return nil
}

//GetFields 获得页面所有字段
func GetFields(query *GetFieldsQuery) error {
	fields := make([]*Field, 0)
	err := sqlstore.X.Where("card_id = ? and org_id=?", query.CardId, query.OrgId).Find(&fields)
	if err != nil {
		return err
	}
	query.Result = fields
	return nil
}

//CreateField 添加字段
func CreateField(cmd *CreateFieldCommand) error {
	field := Field{
		CardId:    cmd.CardId,
		OrgId:     cmd.OrgId,
		Name:      cmd.Name,
		Seq:       cmd.Seq,
		IsVisible: cmd.IsVisible,
		Type:      cmd.Type,
		Default:   cmd.Default,
		Filter:    cmd.Filter,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}
	_, err := sqlstore.X.Insert(&field)
	if err != nil {
		return err
	}
	cmd.Result = field.Id
	return nil
}

//UpdateField 修改字段
func UpdateField(cmd *UpdateFieldCommand) error {
	field := Field{
		Id:        cmd.FieldId,
		OrgId:     cmd.OrgId,
		Name:      cmd.Name,
		IsVisible: cmd.IsVisible,
		Type:      cmd.Type,
		Default:   cmd.Default,
		Filter:    cmd.Filter,
		UpdateAt:  time.Now(),
	}
	_, err := sqlstore.X.Where("id=? and org_id=?", cmd.FieldId, cmd.OrgId).Cols("name").Cols("is_visible").Cols("default").Cols("type").Cols("filter").Cols("update_at").Update(&field)
	if err != nil {
		return err
	}
	success, err2 := sqlstore.X.Id(field.Id).Get(&field)
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
func UpdateFieldSeq(cmd *UpdateFieldSeqCommand) error {
	field := Field{
		Id:       cmd.FieldId,
		OrgId:    cmd.OrgId,
		Seq:      cmd.Seq,
		UpdateAt: time.Now(),
	}
	_, err := sqlstore.X.Where("id=? and org_id=?", cmd.FieldId, cmd.OrgId).Cols("seq").Cols("update_at").Update(&field)
	if err != nil {
		return err
	}
	success, err2 := sqlstore.X.Id(cmd.FieldId).Get(&field)
	if err2 != nil {
		return err2
	}
	if !success {
		return ErrFieldNotFound
	}
	cmd.Result = &field
	return nil
}

//DeleteField 删除字段
func DeleteField(cmd *DeleteFieldCommand) error {
	field := Field{
		Id:    cmd.FieldId,
		OrgId: cmd.OrgId,
	}
	success, err := sqlstore.X.Where("id=? and org_id=?", cmd.FieldId, cmd.OrgId).Get(&field)
	if err != nil {
		return err
	}
	if !success {
		return ErrFieldNotFound
	}
	_, err2 := sqlstore.X.Where("id=? and org_id=?", cmd.FieldId, cmd.OrgId).Delete(&field)
	if err2 != nil {
		return err2
	}
	cmd.Result = &field
	return nil
}
