package locale

type Locale struct {
	Name    string
	Default string
	ZH_CN   string `json:"zh-CN"`
	EN_US   string `json:"en-US"`
}

type International struct {
	Id     int64
	Name   string
	Text   string
	Locale *Locale
}
