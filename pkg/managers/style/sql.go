package style

import (
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/sqlstore"
)

func init() {
	bus.AddHandler("sql", GetStyles)
	bus.AddHandler("sql", CreateStyles)
	bus.AddHandler("sql", UpdateStyle)
	bus.AddHandler("sql", DeleteStyles)
}

//GetStyles get style,if not exists then create
func GetStyles(query *GetStylesQuery) error {
	settings := make([]*StyleSet, 0)
	err := sqlstore.X.Table("style_set").Where("type=?", query.Style).Find(&settings)
	if err != nil {
		return err
	}
	styles := make([]*Style, 0)
	for _, setting := range settings {
		style := Style{}
		_, err2 := sqlstore.X.Where("type = ? and card_id = ? and property = ?", query.Style, query.CardId, setting.Property).Get(&style)
		if err2 != nil {
			return err2
		}
		if style.Id == 0 {
			style := Style{
				CardId:     query.CardId,
				FieldId:    0,
				Type:       query.Style,
				Property:   setting.Property,
				Value:      setting.Value,
				Relation:   setting.Relation,
				MustNumber: setting.MustNumber,
				OrgId:      query.OrgId,
				CreateAt:   time.Now(),
				UpdateAt:   time.Now(),
			}
			_, err3 := sqlstore.X.Insert(&style)
			if err3 != nil {
				return err3
			}
		}
		styles = append(styles, &style)
	}

	query.Result = styles
	return nil
}

func CreateStyles(cmd *CreateStylesCommand) error {
	settings := make([]*StyleSet, 0)
	err := sqlstore.X.Table("style_set").Find(&settings)
	if err != nil {
		return err
	}
	for _, setting := range settings {
		style := Style{}
		_, err2 := sqlstore.X.Where("type = ? and card_id = ? and property = ?", setting.Type, cmd.CardId, setting.Property).Get(&style)
		if err2 != nil {
			return err2
		}
		if style.Id == 0 {
			style := Style{
				CardId:     cmd.CardId,
				FieldId:    0,
				Type:       setting.Type,
				Property:   setting.Property,
				Value:      setting.Value,
				Relation:   setting.Relation,
				MustNumber: setting.MustNumber,
				OrgId:      cmd.OrgId,
				CreateAt:   time.Now(),
				UpdateAt:   time.Now(),
			}
			_, err3 := sqlstore.X.Insert(&style)
			if err3 != nil {
				return err3
			}
		}
	}

	return nil
}
func UpdateStyle(cmd *UpdateStyleCommand) error {
	style := Style{
		Id:       cmd.Id,
		FieldId:  cmd.FieldId,
		Value:    cmd.Value,
		UpdateAt: time.Now(),
	}
	_, err := sqlstore.X.Id(cmd.Id).Cols("field_id").Cols("value").Cols("update_at").Update(&style)
	if err != nil {
		return err
	}
	sqlstore.X.Id(cmd.Id).Get(&style)
	cmd.Result = &style
	return nil
}

func DeleteStyles(cmd *DeleteStylesCommand) error {
	count, err := sqlstore.X.Where("card_id = ?", cmd.CardId).Delete(new(Style))
	if err != nil {
		return err
	}
	cmd.Result = count
	return nil
}
