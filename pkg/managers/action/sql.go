package action

import (
	"errors"
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetViewAction)
	bus.AddHandler("sql", GetViewActions)
	bus.AddHandler("sql", CreateViewAction)
	bus.AddHandler("sql", UpdateViewAction)
	bus.AddHandler("sql", UpdateViewActionSeq)
	bus.AddHandler("sql", DeleteViewAction)
}

var (
	ErrViewActionNotFound = errors.New("view action data not found")
)

//GetViewAction 获得操作
func GetViewAction(query *GetViewActionQuery) error {
	viewAction := ViewAction{Id: query.ViewActionId}
	success, err := sqlstore.X.Where("id= ? and org_id= ?", query.ViewActionId, query.OrgId).Get(&viewAction)
	if err != nil {
		return err
	}
	if !success {
		return ErrViewActionNotFound
	}
	query.Result = &viewAction
	return nil
}

//GetViewActions 获得页面所有操作
func GetViewActions(query *GetViewActionsQuery) error {
	viewActions := make([]*ViewAction, 0)
	err := sqlstore.X.Where("card_id = ? and org_id= ?", query.CardId, query.OrgId).Find(&viewActions)
	if err != nil {
		return err
	}
	query.Result = viewActions
	return nil
}

//CreateViewAction 添加操作
func CreateViewAction(cmd *CreateViewActionCommand) error {
	viewAction := ViewAction{
		CardId:      cmd.CardId,
		OrgId:       cmd.OrgId,
		Name:        cmd.Name,
		Seq:         cmd.Seq,
		Type:        cmd.Type,
		DoubleCheck: cmd.DoubleCheck,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	_, err := sqlstore.X.Insert(&viewAction)
	if err != nil {
		return err
	}
	cmd.Result = viewAction.Id
	return nil
}

//UpdateViewAction 修改操作
func UpdateViewAction(cmd *UpdateViewActionCommand) error {
	viewAction := ViewAction{
		Id:          cmd.ViewActionId,
		OrgId:       cmd.OrgId,
		Name:        cmd.Name,
		Type:        cmd.Type,
		DoubleCheck: cmd.DoubleCheck,
		UpdateAt:    time.Now(),
	}
	_, err := sqlstore.X.Id(cmd.ViewActionId).Cols("name").Cols("type").Cols("double_check").Cols("update_at").Update(&viewAction)
	if err != nil {
		return err
	}
	sqlstore.X.Id(viewAction.Id).Get(&viewAction)
	cmd.Result = &viewAction
	return nil
}

//UpdateViewActionSeq 修改操作顺序
func UpdateViewActionSeq(cmd *UpdateViewActionSeqCommand) error {
	viewAction := ViewAction{
		Id:       cmd.ViewActionId,
		OrgId:    cmd.OrgId,
		Seq:      cmd.Seq,
		UpdateAt: time.Now(),
	}
	_, err := sqlstore.X.Id(cmd.ViewActionId).Cols("seq").Cols("update_at").Update(&viewAction)
	if err != nil {
		return err
	}
	sqlstore.X.Id(viewAction.Id).Get(&viewAction)
	cmd.Result = &viewAction
	return nil
}

//DeleteViewAction 删除操作
func DeleteViewAction(cmd *DeleteViewActionCommand) error {
	viewAction := ViewAction{
		Id:    cmd.ViewActionId,
		OrgId: cmd.OrgId,
	}
	count, err := sqlstore.X.Where("id = ? and org_id= ?", cmd.ViewActionId, cmd.OrgId).Delete(&viewAction)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrViewActionNotFound
	}
	return nil
}
