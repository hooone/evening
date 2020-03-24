package middleware

import (
	"net/url"

	"github.com/hooone/evening/pkg/models"
	"github.com/hooone/evening/pkg/setting"
	macaron "gopkg.in/macaron.v1"
)

type AuthOptions struct {
	ReqSignedIn bool
}

func notAuthorized(c *models.ReqContext) {
	if c.IsApiRequest() {
		c.JsonApiErr(401, "Unauthorized", nil)
		return
	}

	WriteCookie(c.Resp, "redirect_to", url.QueryEscape(c.Req.RequestURI), 0, newCookieOptions)

	c.Redirect(setting.AppSubURL + "/login")
}

func accessForbidden(c *models.ReqContext) {
	if c.IsApiRequest() {
		c.JsonApiErr(403, "Permission denied", nil)
		return
	}

	c.Redirect(setting.AppSubURL + "/")
}

func Auth(options *AuthOptions) macaron.Handler {
	return func(c *models.ReqContext) {
		if !c.IsSignedIn && options.ReqSignedIn {
			notAuthorized(c)
			return
		}

	}
}
