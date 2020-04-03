package parameter

type GetParameterQuery struct {
	OrgId       int64
	ParameterId int64
	Result      *Parameter
}
type GetParametersQuery struct {
	OrgId    int64
	ActionId int64
	Result   []*Parameter
}
type CreateParameterCommand struct {
	OrgId      int64
	ActionId   int64
	FieldId    int64
	IsVisible  bool
	IsEditable bool
	Default    string
	Result     int64
}
type UpdateParameterCommand struct {
	OrgId       int64
	ParameterId int64
	IsVisible   bool
	IsEditable  bool
	Default     string
	Result      *Parameter
}

type DeleteParameterByFieldCommand struct {
	OrgId   int64
	FieldId int64
	Result  int64
}

type DeleteParameterByActionCommand struct {
	OrgId    int64
	ActionId int64
	Result   int64
}
