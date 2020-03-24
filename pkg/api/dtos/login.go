package dtos

type LoginCommand struct {
	User     string `form:"User" `
	Password string `form:"Password"`
	Remember bool   `form:"Remember"`
}
