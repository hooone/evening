package dtos

type GetFieldsForm struct {
	CardId int64 `form:"CardId"`
}
type GetFieldByIdForm struct {
	FieldId int64 `form:"FieldId"`
}

type CreateFieldForm struct {
	Name      string `form:"Name"`
	Text      string `form:"Text"`
	IsVisible bool   `form:"IsVisible"`
	Type      string `form:"Type"`
	Default   string `form:"Default"`
	Filter    string `form:"Filter"`
	CardId    int64  `form:"CardId"`
}
type UpdateFieldForm struct {
	FieldId   int64  `form:"Id"`
	Name      string `form:"Name"`
	Text      string `form:"Text"`
	IsVisible bool   `form:"IsVisible"`
	Type      string `form:"Type"`
	Default   string `form:"Default"`
	Filter    string `form:"Filter"`
}
type UpdateFieldSeqForm struct {
	Source   int64 `form:"Source"`
	Target   int64 `form:"Target"`
	Position int32 `form:"Position"`
}
type DeleteFieldForm struct {
	FieldId int64 `form:"Id"`
}
