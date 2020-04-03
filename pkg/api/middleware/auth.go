package middleware

import (
	"net/url"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/setting"
	macaron "gopkg.in/macaron.v1"
)

type AuthOptions struct {
	ReqSignedIn bool
}

func notAuthorized(c *dtos.ReqContext) {
	if c.IsApiRequest() {
		c.JsonApiErr(401, "Unauthorized", nil)
		return
	}

	WriteCookie(c.Resp, "redirect_to", url.QueryEscape(c.Req.RequestURI), 0, newCookieOptions)

	c.Redirect(setting.AppSubURL + "/login")
}

func accessForbidden(c *dtos.ReqContext) {
	if c.IsApiRequest() {
		c.JsonApiErr(403, "Permission denied", nil)
		return
	}

	c.Redirect(setting.AppSubURL + "/")
}

func Auth(options *AuthOptions) macaron.Handler {
	return func(c *dtos.ReqContext) {
		if !c.IsSignedIn && options.ReqSignedIn {
			notAuthorized(c)
			return
		}

	}
}
