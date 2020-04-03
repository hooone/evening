package locale

import (
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetLocale)
	bus.AddHandler("sql", SetLocale)
	bus.AddHandler("sql", DeleteLocale)
}

func GetLocale(query *GetLocaleQuery) error {
	locales := make([]*LocaleSet, 0)
	err := sqlstore.X.Where("v_id = ? and type = ?", query.VId, query.Type).Find(&locales)
	if err != nil {
		return err
	}
	query.Result = locales
	return nil
}
func SetLocale(cmd *SetLocaleCommand) error {
	locale := LocaleSet{
		VId:      cmd.VId,
		Type:     cmd.Type,
		Language: cmd.Language,
		Name:     cmd.Name,
		Text:     cmd.Text,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	if _, err := sqlstore.X.Where("v_id = ? and type = ? and language = ?", cmd.VId, cmd.Type, cmd.Language).Delete(&LocaleSet{}); err != nil {
		return err
	}

	if _, err := sqlstore.X.Insert(&locale); err != nil {
		return err
	}
	return nil
}
func DeleteLocale(cmd *DeleteLocaleCommand) error {
	if _, err := sqlstore.X.Where("v_id = ? and type = ?", cmd.VId, cmd.Type).Delete(&LocaleSet{}); err != nil {
		return err
	}
	return nil
}
