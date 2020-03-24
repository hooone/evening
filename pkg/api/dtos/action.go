package dtos

type GetViewActionsForm struct {
	CardId int64 `form:"CardId"`
}
type GetViewActionByIdForm struct {
	ViewActionId int64 `form:"ViewActionId"`
}

type CreateViewActionForm struct {
	Name        string `form:"Name"`
	Text        string `form:"Text"`
	Type        string `form:"Type"`
	DoubleCheck bool   `form:"DoubleCheck"`
	CardId      int64  `form:"CardId"`
}
type UpdateViewActionForm struct {
	ViewActionId int64  `form:"Id"`
	Name         string `form:"Name"`
	Text         string `form:"Text"`
	DoubleCheck  bool   `form:"DoubleCheck"`
	Type         string `form:"Type"`
}
type UpdateViewActionSeqForm struct {
	Source   int64 `form:"Source"`
	Target   int64 `form:"Target"`
	Position int32 `form:"Position"`
}
type DeleteViewActionForm struct {
	ViewActionId int64 `form:"Id"`
}
