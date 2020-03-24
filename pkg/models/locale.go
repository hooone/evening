package models

//Locale i18n
type Locale struct {
	Name    string
	Default string
	ZH_CN   string `json:"zh-CN"`
	EN_US   string `json:"en-US"`
}

type LocaleSet struct {
	Id       int64
	VId      int64
	Type     string
	Language string
	Text     string
}

func (l LocaleSet) TableName() string {
	return "locale"
}

type GetLocaleQuery struct {
	Name   string
	VId    int64
	Type   string
	Result Locale
}

type SetLocaleCommand struct {
	VId      int64
	Type     string
	Language string
	Text     string
	Result   int
}
