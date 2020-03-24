package dtos

type GetCardsForm struct {
	PageId   int64  `form:"PageId"`
	PageName string `form:"PageName"`
}
type GetCardByIdForm struct {
	CardId int64 `form:"CardId"`
}

type CreateCardForm struct {
	Name     string `form:"Name"`
	Text     string `form:"Text"`
	PageId   int64  `form:"PageID"`
	PageName string `form:"PageName"`
}
type UpdateCardForm struct {
	CardId int64  `form:"Id"`
	Name   string `form:"Name"`
	Style  string `form:"Style"`
	Text   string `form:"Text"`
	Width  int32  `form:"Width"`
	Pos    int32  `form:"Pos"`
}
type UpdateCardSeqForm struct {
	Source   int64 `form:"Source"`
	Target   int64 `form:"Target"`
	Position int32 `form:"Position"`
}
type DeleteCardForm struct {
	CardId int64 `form:"CardId"`
}
