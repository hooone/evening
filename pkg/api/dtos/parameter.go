package dtos

type UpdateParameterForm struct {
	CardId int64  `form:"CardId"`
	Data   string `form:"Data"`
}
