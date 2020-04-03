package style

type GetStylesQuery struct {
	CardId int64
	OrgId  int64
	Style  string
	Result []*Style
}
type UpdateStyleCommand struct {
	Id      int64
	OrgId   int64
	FieldId int64
	Value   string
	Result  *Style
}

type CreateStylesCommand struct {
	OrgId  int64
	CardId int64
	Result int64
}

type DeleteStylesCommand struct {
	OrgId  int64
	CardId int64
	Result int64
}
