package locale

type GetLocaleQuery struct {
	VId    int64
	Type   string
	Result []*LocaleSet
}

type SetLocaleCommand struct {
	VId      int64
	Type     string
	Language string
	Name     string
	Text     string
	Result   int
}

type DeleteLocaleCommand struct {
	VId    int64
	Type   string
	Result int
}
