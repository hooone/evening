package sqlstore

import (
	"errors"
	"sort"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetCard)
	bus.AddHandler("sql", GetCards)
	bus.AddHandler("sql", CreateCard)
	bus.AddHandler("sql", UpdateCard)
	bus.AddHandler("sql", UpdateCardSeq)
	bus.AddHandler("sql", DeleteCard)
}

//GetCard 获得卡片
func GetCard(query *models.GetCardQuery) error {
	card := models.Card{
		Id:    query.CardId,
		OrgId: query.OrgId,
	}
	success, err := x.Id(query.CardId).Get(&card)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("now card row found")
	}
	query.Result = &card
	return nil
}

//GetCards 获得页面所有卡片
func GetCards(query *models.GetCardsQuery) error {
	cards := make([]*models.Card, 0)
	err := x.Where("page_id = ? and org_id = ?", query.PageId, query.OrgId).Find(&cards)
	if err != nil {
		return err
	}
	sort.Sort(models.CardSlice(cards))
	query.Result = cards
	return nil
}

//CreateCard 添加卡片
func CreateCard(cmd *models.CreateCardCommand) error {
	card := models.Card{
		PageId: cmd.PageId,
		Name:   cmd.Name,
		Style:  cmd.Style,
		Seq:    cmd.Seq,
		Pos:    cmd.Pos,
		Width:  cmd.Width,
		OrgId:  cmd.OrgId,
	}
	_, err := x.Insert(&card)
	if err != nil {
		return err
	}
	cmd.Result = card.Id
	return nil
}

//UpdateCard 修改卡片
func UpdateCard(cmd *models.UpdateCardCommand) error {
	card := models.Card{
		Id:    cmd.CardId,
		Name:  cmd.Name,
		Style: cmd.Style,
		Pos:   cmd.Pos,
		Width: cmd.Width,
	}
	_, err := x.Id(cmd.CardId).Cols("name").Cols("style").Cols("pos").Cols("width").Update(&card)
	if err != nil {
		return err
	}
	x.Id(card.Id).Get(&card)
	cmd.Result = &card
	return nil
}

//UpdateCardSeq 修改卡片顺序
func UpdateCardSeq(cmd *models.UpdateCardSeqCommand) error {
	card := models.Card{
		Id:  cmd.CardId,
		Seq: cmd.Seq,
	}
	_, err := x.Id(cmd.CardId).Cols("seq").Update(&card)
	if err != nil {
		return err
	}
	cmd.Result = 0
	return nil
}

//DeleteCard 删除卡片
func DeleteCard(cmd *models.DeleteCardCommand) error {
	card := models.Card{
		Id: cmd.CardId,
	}
	_, err := x.Id(cmd.CardId).Delete(&card)
	if err != nil {
		return err
	}
	cmd.Result = 0
	return nil
}
