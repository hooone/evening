package sqlstore

import (
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetLocale)
	bus.AddHandler("sql", SetLocale)
}

func GetLocale(query *models.GetLocaleQuery) error {
	query.Result = models.Locale{Name: query.Name, Default: query.Name}
	locales := make([]*models.LocaleSet, 0)
	err := x.Where("v_id = ? and type = ?", query.VId, query.Type).Find(&locales)
	if err != nil {
		return err
	}
	for _, l := range locales {
		if l.Language == "zh-CN" {
			query.Result.ZH_CN = l.Text
		} else if l.Language == "en-US" {
			query.Result.EN_US = l.Text
		}
	}
	return nil
}
func SetLocale(cmd *models.SetLocaleCommand) error {
	locale := models.LocaleSet{
		VId:      cmd.VId,
		Type:     cmd.Type,
		Language: cmd.Language,
		Text:     cmd.Text,
	}
	if _, err := x.Where("v_id = ? and type = ? and language = ?", cmd.VId, cmd.Type, cmd.Language).Delete(&locale); err != nil {
		return err
	}
	if _, err := x.Insert(&locale); err != nil {
		return err
	}
	return nil
}
