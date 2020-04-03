package data

type GetTestDataQuery struct {
	CardId int64
	Result []*TestData
}

type InsertTestDataCommand struct {
	CardId int64
	Value  string
	Result int64
}

type UpdateTestDataCommand struct {
	Key    int64
	Value  string
	Result int64
}

type DeleteTestDataCommand struct {
	Key    int64
	Result int64
}
