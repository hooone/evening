package api

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	dataMg "github.com/hooone/evening/pkg/managers/data"
	paramSvr "github.com/hooone/evening/pkg/services/parameter"
)

//ReadTestData 获取表单的测试数据
func (hs *HTTPServer) ReadTestData(c *dtos.ReqContext, form dtos.ReadTestDataForm, lang dtos.LocaleForm) Response {
	fds, _ := hs.FieldService.GetFields(form.CardId, c.OrgId, lang.Language)
	//get parameters
	Parameters := make([]*paramSvr.Parameter, 0)
	for _, fd := range fds {
		if fd.Filter == "EQUAL" {
			param := paramSvr.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "eq",
				Field:      fd,
			}
			Parameters = append(Parameters, &param)
		} else if fd.Filter == "MAX" {
			param := paramSvr.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "le",
				Field:      fd,
			}
			Parameters = append(Parameters, &param)
		} else if fd.Filter == "MIN" {
			param := paramSvr.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "ge",
				Field:      fd,
			}
			Parameters = append(Parameters, &param)
		} else if fd.Filter == "RANGE" {
			param2 := paramSvr.Parameter{
				FieldId:    fd.Id,
				IsVisible:  fd.IsVisible,
				IsEditable: fd.IsVisible,
				Default:    fd.Default,
				Compare:    "ge",
				Field:      fd,
			}
			Parameters = append(Parameters, &param2)
			param := paramSvr.Parameter{
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
	data, err := getTestData(form.CardId, Parameters)
	if err != nil {
		return Error(500, "Failed to get fields", err)
	}
	result := new(CommonResult)
	result.Data = data
	result.Success = true
	return JSON(200, result)
}

func getTestData(cardId int64, Parameters []*paramSvr.Parameter) ([]interface{}, error) {
	data := make([]interface{}, 0)
	//get data
	query := dataMg.GetTestDataQuery{
		CardId: cardId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return data, err
	}
	//filter
	for _, p := range query.Result {
		m := make(map[string]interface{})
		json.Unmarshal([]byte(p.Value), &m)
		match := true
		for _, para := range Parameters {
			_, ok := m[para.Field.Name]
			if ok {
				match = match && compare(para.Compare, para.Field.Type, m[para.Field.Name], para.Value)
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
	return data, nil
}

func (hs *HTTPServer) CreateTestData(c *dtos.ReqContext, form dtos.CreateTestDataForm, lang dtos.LocaleForm) Response {
	//get parameters
	pas, _ := hs.ParameterService.GetParameters(form.ActionId, c.OrgId, lang.Language)

	//get form data
	mp := make(map[string]interface{})
	for _, parm := range pas {
		pvalue := c.Query(parm.Field.Name)
		mp[parm.Field.Name] = pvalue
	}
	data, _ := json.Marshal(mp)

	//insert data
	cmd := dataMg.InsertTestDataCommand{
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

func (hs *HTTPServer) UpdateTestData(c *dtos.ReqContext, form dtos.UpdateTestDataForm, lang dtos.LocaleForm) Response {
	//get parameters
	pas, _ := hs.ParameterService.GetParameters(form.ActionId, c.OrgId, lang.Language)

	//get form data
	mp := make(map[string]interface{})
	for _, parm := range pas {
		pvalue := c.Query(parm.Field.Name)
		mp[parm.Field.Name] = pvalue
	}
	data, _ := json.Marshal(mp)

	//update data
	cmd := dataMg.UpdateTestDataCommand{
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

func DeleteTestData(c *dtos.ReqContext, form dtos.DeleteTestDataForm) Response {
	cmd := dataMg.DeleteTestDataCommand{
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

func Todatetime(in string) (out time.Time, err error) {
	out, err = time.ParseInLocation("2006-01-02 15:04:05", in, time.Local)
	return out, err
}

func toInt(in interface{}) (int64, error) {
	switch in.(type) {
	case float64:
		return int64(in.(float64)), nil
		break
	case float32:
		return int64(in.(float32)), nil
		break
	case int:
		return int64(in.(int)), nil
		break
	case int64:
		return in.(int64), nil
		break
	case string:
		return strconv.ParseInt(in.(string), 10, 64)
		break
	default:
		return 0, errors.New("type mismatch")
		break
	}
	return 0, errors.New("type mismatch")
}

func toDateTime(in interface{}) (time.Time, error) {
	switch in.(type) {
	case string:
		return Todatetime(in.(string))
		break
	case int64:
		now := time.Now()
		m := 0 - time.Duration(in.(int64)*int64(time.Millisecond))
		return now.Add(m), nil
		break
	case int32:
		now := time.Now()
		m := 0 - time.Duration(int64(in.(int32))*int64(time.Millisecond))
		return now.Add(m), nil
		break
	default:
		return time.Now(), errors.New("type mismatch")
		break
	}
	return time.Now(), errors.New("type mismatch")
}
func toDate(in interface{}) (time.Time, error) {
	switch in.(type) {
	case string:
		sp := strings.Split(in.(string), " ")
		return Todatetime(sp[0] + " 00:00:00")
		break
	case int64:
		now := time.Now()
		m := 0 - time.Duration(in.(int64)*int64(time.Millisecond))
		return now.Add(m), nil
		break
	case int32:
		now := time.Now()
		m := 0 - time.Duration(int64(in.(int32))*int64(time.Millisecond))
		return now.Add(m), nil
		break
	default:
		return time.Now(), errors.New("type mismatch")
		break
	}
	return time.Now(), errors.New("type mismatch")
}
func toString(in interface{}) (string, error) {
	switch in.(type) {
	case string:
		return in.(string), nil
		break
	default:
		return "", errors.New("type mismatch")
		break
	}
	return "", errors.New("type mismatch")
}
func compare(operation string, typec string, left interface{}, right interface{}) bool {
	if typec == "int" {
		leftvalue, ok1 := toInt(left)
		rightvalue, ok2 := toInt(right)
		if ok1 != nil || ok2 != nil {
			return true
		}
		if operation == "lt" {
			return leftvalue < rightvalue
		} else if operation == "le" {
			return leftvalue <= rightvalue
		} else if operation == "gt" {
			return leftvalue > rightvalue
		} else if operation == "ge" {
			return leftvalue >= rightvalue
		} else if operation == "eq" {
			return leftvalue == rightvalue
		} else if operation == "ne" {
			return leftvalue != rightvalue
		}
	} else if typec == "datetime" {
		leftvalue, ok1 := toDateTime(left)
		rightvalue, ok2 := toDateTime(right)
		if ok1 != nil || ok2 != nil {
			return false
		}
		if operation == "lt" {
			return leftvalue.UnixNano() < rightvalue.UnixNano()
		} else if operation == "le" {
			return leftvalue.UnixNano() <= rightvalue.UnixNano()
		} else if operation == "gt" {
			return leftvalue.UnixNano() > rightvalue.UnixNano()
		} else if operation == "ge" {
			return leftvalue.UnixNano() >= rightvalue.UnixNano()
		} else if operation == "eq" {
			return leftvalue.UnixNano() == rightvalue.UnixNano()
		} else if operation == "ne" {
			return leftvalue.UnixNano() != rightvalue.UnixNano()
		}
	} else if typec == "date" {
		leftvalue, ok1 := toDate(left)
		rightvalue, ok2 := toDate(right)
		if ok1 != nil || ok2 != nil {
			return false
		}
		if operation == "lt" {
			return leftvalue.UnixNano() < rightvalue.UnixNano()
		} else if operation == "le" {
			return leftvalue.UnixNano() <= rightvalue.UnixNano()
		} else if operation == "gt" {
			return leftvalue.UnixNano() > rightvalue.UnixNano()
		} else if operation == "ge" {
			return leftvalue.UnixNano() >= rightvalue.UnixNano()
		} else if operation == "eq" {
			return leftvalue.UnixNano() == rightvalue.UnixNano()
		} else if operation == "ne" {
			return leftvalue.UnixNano() != rightvalue.UnixNano()
		}
	} else if typec == "string" {
		leftvalue, ok1 := toString(left)
		rightvalue, ok2 := toString(right)
		if ok1 != nil || ok2 != nil {
			return false
		}
		if operation == "lt" {
			return strings.Index(leftvalue, rightvalue) >= 0
		} else if operation == "le" {
			return strings.Index(leftvalue, rightvalue) >= 0
		} else if operation == "gt" {
			return strings.Index(leftvalue, rightvalue) >= 0
		} else if operation == "ge" {
			return strings.Index(leftvalue, rightvalue) >= 0
		} else if operation == "eq" {
			return leftvalue == rightvalue
		} else if operation == "ne" {
			return leftvalue != rightvalue
		}
	}
	return false
}
