package sqlstore

import (
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetStyles)
	bus.AddHandler("sql", UpdateStyle)
}

func GetStyles(query *models.GetStylesQuery) error {
	settings := make([]*models.StyleSet, 0)
	err := x.Table("style_set").Where("type=?", query.Style).Find(&settings)
	if err != nil {
		return err
	}
	styles := make([]*models.Style, 0)
	for _, setting := range settings {
		style := models.Style{}
		_, err2 := x.Where("type = ? and card_id = ? and property = ?", query.Style, query.CardId, setting.Property).Get(&style)
		if err2 != nil {
			return err2
		}
		if style.Id == 0 {
			style := models.Style{
				CardId:     query.CardId,
				FieldId:    0,
				Type:       query.Style,
				Property:   setting.Property,
				Value:      setting.Value,
				Relation:   setting.Relation,
				MustNumber: setting.MustNumber,
			}
			_, err3 := x.Insert(&style)
			if err3 != nil {
				return err3
			}
		}
		styles = append(styles, &style)
	}
	for _, style := range styles {
		if style.FieldId != 0 {
			Fquery := models.GetFieldQuery{
				FieldId: style.FieldId,
			}
			if err4 := bus.Dispatch(&Fquery); err != nil {
				return err4
			}
			lQuery := models.GetLocaleQuery{
				VId:  Fquery.Result.Id,
				Type: "field",
			}
			if err4 := bus.Dispatch(&lQuery); err != nil {
				return err4
			}
			Fquery.Result.Locale = lQuery.Result
			style.Field = Fquery.Result
		}
	}
	query.Result = styles
	return nil
}

func UpdateStyle(cmd *models.UpdateStyleCommand) error {
	style := models.Style{
		Id:      cmd.Id,
		FieldId: cmd.FieldId,
		Value:   cmd.Value,
	}
	_, err := x.Id(cmd.Id).Cols("field_id").Cols("value").Update(&style)
	if err != nil {
		return err
	}
	x.Id(cmd.Id).Get(&style)
	cmd.Result = &style
	return nil
}
