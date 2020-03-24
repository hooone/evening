package api

import (
	"encoding/json"
	"strconv"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

//ReadTestData 获取表单的测试数据
func ReadTestData(c *models.ReqContext, form dtos.ReadTestDataForm) Response {
	fQuery := models.GetFieldsQuery{CardId: form.CardId}
	if err := bus.Dispatch(&fQuery); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	//get parameters
	Parameters := make([]*models.Parameter, 0)
	for _, fd := range fQuery.Result {
		if fd.Filter == "EQUAL" {
			param := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "eq",
				Field:      fd,
			}
			Parameters = append(Parameters, &param)
		} else if fd.Filter == "MAX" {
			param := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "le",
				Field:      fd,
			}
			Parameters = append(Parameters, &param)
		} else if fd.Filter == "MIN" {
			param := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "ge",
				Field:      fd,
			}
			Parameters = append(Parameters, &param)
		} else if fd.Filter == "RANGE" {
			param2 := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "ge",
				Field:      fd,
			}
			Parameters = append(Parameters, &param2)
			param := models.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "le",
				Field:      fd,
			}
			Parameters = append(Parameters, &param)
		}
	}
	for idx, param := range Parameters {
		param.Value = c.Query("param" + strconv.Itoa(idx))
	}
	//get data
	query := models.GetTestDataQuery{
		CardId: form.CardId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get fields", err)
	}
	//filter
	data := make([]interface{}, 0)
	for _, p := range query.Result {
		m := make(map[string]interface{})
		json.Unmarshal([]byte(p.Value), &m)
		match := true
		for _, para := range Parameters {
			_, ok := m[para.Field.Name]
			if ok {
				match = match && models.Compare(para.Compare, para.Field.Type, m[para.Field.Name], para.Value)
			} else {
				match = false
			}
		}
		if match {
			var mp map[string]interface{}
			json.Unmarshal([]byte(p.Value), &mp)
			mp["__Key"] = strconv.FormatInt(p.Id, 10)
			mp["__CardId"] = p.CardId
			data = append(data, mp)
		}
	}

	result := new(CommonResult)
	result.Data = data
	result.Success = true
	return JSON(200, result)
}

func CreateTestData(c *models.ReqContext, form dtos.CreateTestDataForm) Response {
	//get parameters
	query := models.GetParametersQuery{ActionId: form.ActionId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get paramters", err)
	}

	//get form data
	mp := make(map[string]interface{})
	for _, parm := range query.Result {
		pvalue := c.Query(parm.Field.Name)
		mp[parm.Field.Name] = pvalue
	}
	data, _ := json.Marshal(mp)

	//insert data
	cmd := models.InsertTestDataCommand{
		CardId: form.CardId,
		Value:  string(data),
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to insert test data", err)
	}

	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

func UpdateTestData(c *models.ReqContext, form dtos.UpdateTestDataForm) Response {
	//get parameters
	query := models.GetParametersQuery{ActionId: form.ActionId}
	if err := bus.Dispatch(&query); err != nil {
		return Error(500, "Failed to get paramters", err)
	}

	//get form data
	mp := make(map[string]interface{})
	for _, parm := range query.Result {
		pvalue := c.Query(parm.Field.Name)
		mp[parm.Field.Name] = pvalue
	}
	data, _ := json.Marshal(mp)

	//update data
	cmd := models.UpdateTestDataCommand{
		Key:   form.Key,
		Value: string(data),
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to update test data", err)
	}

	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}

func DeleteTestData(c *models.ReqContext, form dtos.DeleteTestDataForm) Response {
	cmd := models.DeleteTestDataCommand{
		Key: form.Key,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return Error(500, "Failed to delete test data", err)
	}

	result := new(CommonResult)
	result.Data = cmd.Result
	result.Success = true
	return JSON(200, result)
}
