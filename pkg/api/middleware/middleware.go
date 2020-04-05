package middleware

import (
	"net/url"
	"time"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/log"
	"github.com/hooone/evening/pkg/managers/user"
	"github.com/hooone/evening/pkg/services/auth"
	"github.com/hooone/evening/pkg/setting"
	"gopkg.in/macaron.v1"
)

var (
	ReqSignedIn = Auth(&AuthOptions{ReqSignedIn: true})
)

func GetContextHandler(ats *auth.UserTokenService) macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &dtos.ReqContext{
			Context:      c,
			SignedInUser: &user.SignedInUser{},
			IsSignedIn:   false,
			Logger:       log.New("context"),
		}

		orgId := int64(0)
		// orgIdHeader := ctx.Req.Header.Get("X-Evening-Org-Id")
		// if orgIdHeader != "" {
		// 	orgId, _ = strconv.ParseInt(orgIdHeader, 10, 64)
		// }
		switch {
		case initContextWithToken(ats, ctx, orgId):
		case initContextWithAnonymousUser(ctx):
		}

		ctx.Logger = log.New("context", "userId", ctx.Id, "orgId", ctx.OrgId, "uname", ctx.Login)
		ctx.Data["ctx"] = ctx

		c.Map(ctx)

		// update last seen every 5min
		if ctx.ShouldUpdateLastSeenAt() {
			ctx.Logger.Debug("Updating last user_seen_at", "user_id", ctx.SignedInUser.Id)
			if err := bus.Dispatch(&user.UpdateUserLastSeenAtCommand{UserId: ctx.SignedInUser.Id}); err != nil {
				ctx.Logger.Error("Failed to update last_seen_at", "error", err)
			}
		}
	}
}

func initContextWithAnonymousUser(ctx *dtos.ReqContext) bool {
	ctx.IsSignedIn = false
	ctx.SignedInUser = &user.SignedInUser{IsAnonymous: true}
	ctx.OrgId = 0
	return true
}
func initContextWithToken(
	authTokenService *auth.UserTokenService,
	ctx *dtos.ReqContext, orgID int64) bool {
	if setting.LoginCookieName == "" {
		return false
	}

	rawToken := ctx.GetCookie(setting.LoginCookieName)
	if rawToken == "" {
		return false
	}

	token, err := authTokenService.LookupToken(ctx.Req.Context(), rawToken)
	if err != nil {
		ctx.Logger.Error("Failed to look up user based on cookie", "error", err)
		WriteSessionCookie(ctx, "", -1)
		return false
	}

	//获取用户信息
	query := user.GetSignedInUserQuery{UserId: token.UserId, OrgId: orgID}
	if err := bus.Dispatch(&query); err != nil {
		ctx.Logger.Error("Failed to get user with id", "userId", token.UserId, "error", err)
		return false
	}

	ctx.SignedInUser = query.Result
	ctx.IsSignedIn = true
	ctx.UserToken = token

	return true
}

func WriteSessionCookie(ctx *dtos.ReqContext, value string, maxLifetimeDays int) {
	if setting.Env == setting.DEV {
		ctx.Logger.Info("New token", "unhashed token", value)
	}

	var maxAge int
	if maxLifetimeDays <= 0 {
		maxAge = -1
	} else {
		maxAgeHours := (time.Duration(setting.LoginMaxLifetimeDays) * 24 * time.Hour) + time.Hour
		maxAge = int(maxAgeHours.Seconds())
	}

	WriteCookie(ctx.Resp, setting.LoginCookieName, url.QueryEscape(value), maxAge, newCookieOptions)
}
