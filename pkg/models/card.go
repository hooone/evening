package models

//Card view card is base element in evening
type Card struct {
	Id      int64
	PageId  int64
	OrgId   int64
	Name    string
	Seq     int32
	Pos     int32
	Width   int32
	Style   string
	Locale  Locale        `xorm:"-"`
	Fields  []*Field      `xorm:"-"`
	Actions []*ViewAction `xorm:"-"`
	Styles  []*Style      `xorm:"-"`
}
type CardSlice []*Card

func (s CardSlice) Len() int           { return len(s) }
func (s CardSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s CardSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }

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
	Name   string
	Style  string
	Pos    int32
	Width  int32
	Result *Card
}
type UpdateCardSeqCommand struct {
	CardId int64
	Seq    int32
	Result int64
}
type DeleteCardCommand struct {
	CardId int64
	Result int64
}
