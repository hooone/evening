package card

import (
	"errors"
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetCard)
	bus.AddHandler("sql", GetCards)
	bus.AddHandler("sql", CreateCard)
	bus.AddHandler("sql", UpdateCard)
	bus.AddHandler("sql", UpdateCardSeq)
	bus.AddHandler("sql", DeleteCard)
}

var (
	ErrCardNotFound = errors.New("card not found")
)

//GetCard 获得窗体
func GetCard(query *GetCardQuery) error {
	card := Card{
		Id:    query.CardId,
		OrgId: query.OrgId,
	}
	success, err := sqlstore.X.Where("id=? and org_id = ?", query.CardId, query.OrgId).Get(&card)
	if err != nil {
		return err
	}
	if !success {
		return ErrCardNotFound
	}
	query.Result = &card
	return nil
}

//GetCards 获得页面所有窗体
func GetCards(query *GetCardsQuery) error {
	cards := make([]*Card, 0)
	err := sqlstore.X.Where("page_id = ? and org_id = ?", query.PageId, query.OrgId).Find(&cards)
	if err != nil {
		return err
	}
	query.Result = cards
	return nil
}

//CreateCard 添加窗体
func CreateCard(cmd *CreateCardCommand) error {
	card := Card{
		PageId:   cmd.PageId,
		OrgId:    cmd.OrgId,
		Name:     cmd.Name,
		Style:    cmd.Style,
		Seq:      cmd.Seq,
		Pos:      cmd.Pos,
		Width:    cmd.Width,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	_, err := sqlstore.X.Insert(&card)
	if err != nil {
		return err
	}
	cmd.Result = card.Id
	return nil
}

//UpdateCard 修改窗体
func UpdateCard(cmd *UpdateCardCommand) error {
	card := Card{
		Id:       cmd.CardId,
		OrgId:    cmd.OrgId,
		Name:     cmd.Name,
		Style:    cmd.Style,
		Pos:      cmd.Pos,
		Width:    cmd.Width,
		UpdateAt: time.Now(),
	}
	_, err := sqlstore.X.Where("id=? and org_id=?", cmd.CardId, cmd.OrgId).Cols("name").Cols("style").Cols("pos").Cols("width").Cols("update_at").Update(&card)
	if err != nil {
		return err
	}
	sqlstore.X.Id(card.Id).Get(&card)
	cmd.Result = &card
	return nil
}

//UpdateCardSeq 修改窗体顺序
func UpdateCardSeq(cmd *UpdateCardSeqCommand) error {
	card := Card{
		Id:       cmd.CardId,
		OrgId:    cmd.OrgId,
		Seq:      cmd.Seq,
		UpdateAt: time.Now(),
	}
	_, err := sqlstore.X.Where("id=? and org_id=?", cmd.CardId, cmd.OrgId).Cols("seq").Cols("update_at").Update(&card)
	if err != nil {
		return err
	}
	cmd.Result = 0
	return nil
}

//DeleteCard 删除窗体
func DeleteCard(cmd *DeleteCardCommand) error {
	card := Card{
		Id:    cmd.CardId,
		OrgId: cmd.OrgId,
	}
	_, err := sqlstore.X.Where("id=? and org_id=?", cmd.CardId, cmd.OrgId).Delete(&card)
	if err != nil {
		return err
	}
	cmd.Result = 0
	return nil
}
