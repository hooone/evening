package sqlstore

import (
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetTestData)
	bus.AddHandler("sql", InsertData)
	bus.AddHandler("sql", UpdateData)
	bus.AddHandler("sql", DeleteData)
}

//GetTestData 获得数据
func GetTestData(query *models.GetTestDataQuery) error {
	datas := make([]*models.TestData, 0)
	err := x.Where("card_id = ?", query.CardId).Find(&datas)
	if err != nil {
		return err
	}
	query.Result = datas
	return nil
}

//InsertData 填入数据
func InsertData(cmd *models.InsertTestDataCommand) error {
	key := models.ParseKey(cmd.CardId, models.TimeNow())
	data := models.TestData{
		Id:     key,
		CardId: cmd.CardId,
		Value:  cmd.Value,
	}
	_, err := x.Insert(&data)
	if err != nil {
		return err
	}
	cmd.Result = key
	return nil
}

//UpdateData 修改数据
func UpdateData(cmd *models.UpdateTestDataCommand) error {
	data := models.TestData{
		Id:    cmd.Key,
		Value: cmd.Value,
	}
	_, err := x.Id(cmd.Key).Cols("value").Update(&data)
	if err != nil {
		return err
	}
	cmd.Result = 0
	return nil
}

//DeleteData 删除数据
func DeleteData(cmd *models.DeleteTestDataCommand) error {
	data := models.TestData{
		Id: cmd.Key,
	}
	_, err := x.Id(cmd.Key).Delete(&data)
	if err != nil {
		return err
	}
	cmd.Result = 0
	return nil
}
