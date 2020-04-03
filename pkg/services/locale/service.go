package locale

import (
	"reflect"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/locale"
	"github.com/hooone/evening/pkg/registry"
)

func init() {
	registry.RegisterService(&LocaleService{})
}

type LocaleService struct {
}

func (s *LocaleService) Init() error {
	return nil
}

func (s *LocaleService) GetLocale(f interface{}, lang string) error {
	v := reflect.ValueOf(f)
	typeName := reflect.TypeOf(f).Elem().Name()
	v = v.Elem()
	id, s1 := v.FieldByName("Id").Interface().(int64)
	name, s2 := v.FieldByName("Name").Interface().(string)
	loc := Locale{
		Name:    name,
		Default: name,
	}
	if s1 && s2 {
		query := locale.GetLocaleQuery{
			VId:  id,
			Type: typeName,
		}
		if err := bus.Dispatch(&query); err != nil {
			return err
		}
		for _, l := range query.Result {
			if l.Language == lang {
				loc.Default = l.Text
			}
			if l.Language == "zh-CN" {
				loc.ZH_CN = l.Text
			} else if l.Language == "en-US" {
				loc.EN_US = l.Text
			}
		}
	}
	locT := v.FieldByName("Locale")
	if locT.Kind() == reflect.Ptr {
		locT.Set(reflect.ValueOf(&loc))
	} else {
		locT.Set(reflect.ValueOf(loc))
	}
	return nil
}

func (s *LocaleService) SetLocale(f interface{}, lang string) error {
	v := reflect.ValueOf(f)
	typeName := reflect.TypeOf(f).Name()
	id, s1 := v.FieldByName("Id").Interface().(int64)
	name, s2 := v.FieldByName("Name").Interface().(string)
	text, s3 := v.FieldByName("Text").Interface().(string)
	if s1 && s2 && s3 {
		cmd := locale.SetLocaleCommand{
			VId:      id,
			Type:     typeName,
			Name:     name,
			Language: lang,
			Text:     text,
		}
		if err := bus.Dispatch(&cmd); err != nil {
			return err
		}
	}
	return nil
}

func (s *LocaleService) DeleteLocale(f interface{}, lang string) error {
	v := reflect.ValueOf(f)
	typeName := reflect.TypeOf(f).Name()
	id, s1 := v.FieldByName("Id").Interface().(int64)
	if s1 {
		cmd := locale.DeleteLocaleCommand{
			VId:  id,
			Type: typeName,
		}
		if err := bus.Dispatch(&cmd); err != nil {
			return err
		}
	}
	return nil
}
