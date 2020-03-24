package sqlstore

import (
	"errors"
	"sort"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetViewAction)
	bus.AddHandler("sql", GetViewActions)
	bus.AddHandler("sql", CreateViewAction)
	bus.AddHandler("sql", UpdateViewAction)
	bus.AddHandler("sql", UpdateViewActionSeq)
	bus.AddHandler("sql", DeleteViewAction)
}

//GetViewAction 获得操作
func GetViewAction(query *models.GetViewActionQuery) error {
	viewAction := models.ViewAction{Id: query.ViewActionId}
	success, err := x.Id(query.ViewActionId).Get(&viewAction)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("now viewAction row found")
	}
	query.Result = &viewAction
	return nil
}

//GetViewActions 获得页面所有操作
func GetViewActions(query *models.GetViewActionsQuery) error {
	viewActions := make([]*models.ViewAction, 0)
	err := x.Where("card_id = ?", query.CardId).Find(&viewActions)
	if err != nil {
		return err
	}
	sort.Sort(models.ViewActionSlice(viewActions))
	query.Result = viewActions
	return nil
}

//CreateViewAction 添加操作
func CreateViewAction(cmd *models.CreateViewActionCommand) error {
	viewAction := models.ViewAction{
		CardId:      cmd.CardId,
		Name:        cmd.Name,
		Seq:         cmd.Seq,
		Type:        cmd.Type,
		DoubleCheck: cmd.DoubleCheck,
	}
	_, err := x.Insert(&viewAction)
	if err != nil {
		return err
	}
	cmd.Result = viewAction.Id
	return nil
}

//UpdateViewAction 修改操作
func UpdateViewAction(cmd *models.UpdateViewActionCommand) error {
	viewAction := models.ViewAction{
		Id:          cmd.ViewActionId,
		Name:        cmd.Name,
		Type:        cmd.Type,
		DoubleCheck: cmd.DoubleCheck,
	}
	_, err := x.Id(cmd.ViewActionId).Cols("name").Cols("type").Cols("double_check").Update(&viewAction)
	if err != nil {
		return err
	}
	x.Id(viewAction.Id).Get(&viewAction)
	cmd.Result = &viewAction
	return nil
}

//UpdateViewActionSeq 修改操作顺序
func UpdateViewActionSeq(cmd *models.UpdateViewActionSeqCommand) error {
	viewAction := models.ViewAction{
		Id:  cmd.ViewActionId,
		Seq: cmd.Seq,
	}
	_, err := x.Id(cmd.ViewActionId).Cols("seq").Update(&viewAction)
	if err != nil {
		return err
	}
	x.Id(viewAction.Id).Get(&viewAction)
	cmd.Result = &viewAction
	return nil
}

//DeleteViewAction 删除操作
func DeleteViewAction(cmd *models.DeleteViewActionCommand) error {
	viewAction := models.ViewAction{
		Id: cmd.ViewActionId,
	}
	success, err := x.Id(cmd.ViewActionId).Get(&viewAction)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("no row found")
	}

	_, err2 := x.Id(cmd.ViewActionId).Delete(&viewAction)
	if err2 != nil {
		return err2
	}

	paraCmd := models.DeleteParameterByActionCommand{ActionId: viewAction.Id}
	if err3 := bus.Dispatch(&paraCmd); err != nil {
		return err3
	}
	cmd.Result = &viewAction
	return nil
}
