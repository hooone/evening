package action

type GetViewActionQuery struct {
	ViewActionId int64
	OrgId        int64
	Result       *ViewAction
}
type GetViewActionsQuery struct {
	CardId int64
	OrgId  int64
	Result []*ViewAction
}
type CreateViewActionCommand struct {
	OrgId       int64
	CardId      int64
	Name        string
	Seq         int32
	Type        string
	DoubleCheck bool
	Result      int64
}
type UpdateViewActionCommand struct {
	OrgId        int64
	ViewActionId int64
	Name         string
	Type         string //CRUD
	DoubleCheck  bool
	Result       *ViewAction
}
type UpdateViewActionSeqCommand struct {
	OrgId        int64
	ViewActionId int64
	Seq          int32
	Result       *ViewAction
}
type DeleteViewActionCommand struct {
	OrgId        int64
	ViewActionId int64
	Result       *ViewAction
}
