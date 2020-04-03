package data

import (
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetTestData)
	bus.AddHandler("sql", InsertData)
	bus.AddHandler("sql", UpdateData)
	bus.AddHandler("sql", DeleteData)
}

//GetTestData 获得数据
func GetTestData(query *GetTestDataQuery) error {
	datas := make([]*TestData, 0)
	err := sqlstore.X.Where("card_id = ?", query.CardId).Find(&datas)
	if err != nil {
		return err
	}
	query.Result = datas
	return nil
}

//InsertData 填入数据
func InsertData(cmd *InsertTestDataCommand) error {
	key := parseKey(cmd.CardId, TimeNow())
	data := TestData{
		Id:     key,
		CardId: cmd.CardId,
		Value:  cmd.Value,
	}
	_, err := sqlstore.X.Insert(&data)
	if err != nil {
		return err
	}
	cmd.Result = key
	return nil
}

//UpdateData 修改数据
func UpdateData(cmd *UpdateTestDataCommand) error {
	data := TestData{
		Id:    cmd.Key,
		Value: cmd.Value,
	}
	_, err := sqlstore.X.Id(cmd.Key).Cols("value").Update(&data)
	if err != nil {
		return err
	}
	cmd.Result = 0
	return nil
}

//DeleteData 删除数据
func DeleteData(cmd *DeleteTestDataCommand) error {
	data := TestData{
		Id: cmd.Key,
	}
	_, err := sqlstore.X.Id(cmd.Key).Delete(&data)
	if err != nil {
		return err
	}
	cmd.Result = 0
	return nil
}

func parseKey(cardId int64, time JsonTime) int64 {
	var rst int64
	rst = rst + (cardId << 48)
	rst = rst + (time.UnixNano() >> 16)
	return rst
}
