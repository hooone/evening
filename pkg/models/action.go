package models

//ViewAction view ViewAction is base element in evening
type ViewAction struct {
	Id          int64
	CardId      int64
	Name        string
	Type        string //CRUD
	Seq         int32
	DoubleCheck bool
	Locale      Locale       `xorm:"-"`
	Parameters  []*Parameter `xorm:"-"`
}
type ViewActionSlice []*ViewAction

func (s ViewActionSlice) Len() int           { return len(s) }
func (s ViewActionSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ViewActionSlice) Less(i, j int) bool { return s[i].Seq < s[j].Seq }

type GetViewActionQuery struct {
	ViewActionId int64
	Result       *ViewAction
}
type GetViewActionsQuery struct {
	CardId int64
	Result []*ViewAction
}
type CreateViewActionCommand struct {
	CardId      int64
	Name        string
	Seq         int32
	Type        string
	DoubleCheck bool
	Result      int64
}
type UpdateViewActionCommand struct {
	ViewActionId int64
	Name         string
	Type         string //CRUD
	DoubleCheck  bool
	Result       *ViewAction
}
type UpdateViewActionSeqCommand struct {
	ViewActionId int64
	Seq          int32
	Result       *ViewAction
}
type DeleteViewActionCommand struct {
	ViewActionId int64
	Result       *ViewAction
}
