package models

import (
	"encoding/binary"
	"errors"
	"strconv"
	"strings"
	"time"
)

//Card view card is base element in evening
type TestData struct {
	Id     int64
	CardId int64
	Value  string
}

type GetTestDataQuery struct {
	CardId int64
	Result []*TestData
}

type InsertTestDataCommand struct {
	CardId int64
	Value  string
	Result int64
}

type UpdateTestDataCommand struct {
	Key    int64
	Value  string
	Result int64
}

type DeleteTestDataCommand struct {
	Key    int64
	Result int64
}

type JsonDate time.Time
type JsonTime time.Time

func (p *JsonDate) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)
	*p = JsonDate(local)
	return err
}

func (p *JsonTime) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*p = JsonTime(local)
	return err
}

func (c JsonDate) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, "2006-01-02")
	data = append(data, '"')
	return data, nil
}

func (c JsonTime) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, "2006-01-02 15:04:05")
	data = append(data, '"')
	return data, nil
}

func TimeNow() JsonTime {
	return JsonTime(time.Now())
}

func (c JsonDate) String() string {
	return time.Time(c).Format("2006-01-02")
}

func (c JsonTime) String() string {
	return time.Time(c).Format("2006-01-02 15:04:05")
}

func (c JsonTime) UnixNano() int64 {
	return time.Time(c).UnixNano()
}
func Todate(in string) (out time.Time, err error) {
	out, err = time.Parse("2006-01-02", in)
	return out, err
}

func Todatetime(in string) (out time.Time, err error) {
	out, err = time.ParseInLocation("2006-01-02 15:04:05", in, time.Local)
	return out, err
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func ParseKey(cardId int64, time JsonTime) int64 {
	var rst int64
	rst = rst + (cardId << 48)
	rst = rst + (time.UnixNano() >> 16)
	return rst
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
func Compare(operation string, typec string, left interface{}, right interface{}) bool {
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
			return true
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
