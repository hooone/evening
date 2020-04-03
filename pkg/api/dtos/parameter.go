package dtos

type GetParametersForm struct {
	ActionId int64 `form:"ActionId"`
}

type UpdateParameterForm struct {
	CardId int64  `form:"CardId"`
	Data   string `form:"Data"`
}
