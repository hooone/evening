package dtos

type LoginCommand struct {
	User     string `form:"User" `
	Password string `form:"Password"`
	Remember bool   `form:"Remember"`
}

type SignUpStep2Form struct {
	Username string `form:"User"`
	Password string `form:"Password"`
}
