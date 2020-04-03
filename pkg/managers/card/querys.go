package card

type GetCardQuery struct {
	CardId int64
	OrgId  int64
	Result *Card
}
type GetCardsQuery struct {
	PageId int64
	OrgId  int64
	Result []*Card
}
type CreateCardCommand struct {
	PageId int64
	OrgId  int64
	Name   string
	Style  string
	Seq    int32
	Pos    int32
	Width  int32
	Result int64
}
type UpdateCardCommand struct {
	CardId int64
	OrgId  int64
	Name   string
	Style  string
	Pos    int32
	Width  int32
	Result *Card
}
type UpdateCardSeqCommand struct {
	CardId int64
	OrgId  int64
	Seq    int32
	Result int64
}
type DeleteCardCommand struct {
	CardId int64
	OrgId  int64
	Result int64
}
