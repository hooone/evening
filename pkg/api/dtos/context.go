package dtos

import (
	"strings"

	"github.com/hooone/evening/pkg/log"
	"github.com/hooone/evening/pkg/managers/user"
	"github.com/hooone/evening/pkg/setting"
	"gopkg.in/macaron.v1"
)

type ReqContext struct {
	*macaron.Context
	*user.SignedInUser
	UserToken *user.UserToken

	IsSignedIn bool
	Logger     log.Logger
}

func (ctx *ReqContext) IsApiRequest() bool {
	return strings.HasPrefix(ctx.Req.URL.Path, "/api")
}

func (ctx *ReqContext) JsonApiErr(status int, message string, err error) {
	resp := make(map[string]interface{})

	if err != nil {
		ctx.Logger.Error(message, "error", err)
		if setting.Env != setting.PROD {
			resp["error"] = err.Error()
		}
	}

	switch status {
	case 404:
		resp["message"] = "Not Found"
	case 500:
		resp["message"] = "Internal Server Error"
	}

	if message != "" {
		resp["message"] = message
	}

	ctx.JSON(status, resp)
}
