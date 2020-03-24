package dtos

type ReadTestDataForm struct {
	CardId int64 `form:"__CardId"`
}

type CreateTestDataForm struct {
	CardId   int64 `form:"__CardId"`
	ActionId int64 `form:"__ActionId"`
}

type UpdateTestDataForm struct {
	CardId   int64 `form:"__CardId"`
	Key      int64 `form:"__Key"`
	ActionId int64 `form:"__ActionId"`
}

type DeleteTestDataForm struct {
	Key int64 `form:"__Key"`
}
