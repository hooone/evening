package middleware

import (
	"net/url"
	"time"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/infra/log"
	"github.com/hooone/evening/pkg/models"
	"github.com/hooone/evening/pkg/services/auth"
	"github.com/hooone/evening/pkg/setting"
	"gopkg.in/macaron.v1"
)

var (
	ReqSignedIn = Auth(&AuthOptions{ReqSignedIn: true})
)

func GetContextHandler(ats *auth.UserTokenService) macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &models.ReqContext{
			Context:      c,
			SignedInUser: &models.SignedInUser{},
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

		// 超时登出
		// if ctx.ShouldUpdateLastSeenAt() {
		// 	ctx.Logger.Debug("Updating last user_seen_at", "user_id", ctx.UserId)
		// 	if err := bus.Dispatch(&models.UpdateUserLastSeenAtCommand{UserId: ctx.UserId}); err != nil {
		// 		ctx.Logger.Error("Failed to update last_seen_at", "error", err)
		// 	}
		// }
	}
}

func initContextWithAnonymousUser(ctx *models.ReqContext) bool {
	ctx.IsSignedIn = false
	ctx.SignedInUser = &models.SignedInUser{IsAnonymous: true}
	ctx.OrgId = 0
	return true
}
func initContextWithToken(
	authTokenService *auth.UserTokenService,
	ctx *models.ReqContext, orgID int64) bool {
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
	query := models.GetSignedInUserQuery{UserId: token.UserId, OrgId: orgID}
	if err := bus.Dispatch(&query); err != nil {
		ctx.Logger.Error("Failed to get user with id", "userId", token.UserId, "error", err)
		return false
	}

	ctx.SignedInUser = query.Result
	ctx.IsSignedIn = true
	ctx.Test = "AABBCC"
	ctx.UserToken = token

	// Rotate the token just before we write response headers to ensure there is no delay between
	// the new token being generated and the client receiving it.
	// ctx.Resp.Before(rotateEndOfRequestFunc(ctx, authTokenService, token))

	return true
}

//token超时后自动生成
// func rotateEndOfRequestFunc(ctx *models.ReqContext, authTokenService models.UserTokenService, token *models.UserToken) macaron.BeforeFunc {
// 	return func(w macaron.ResponseWriter) {
// 		// // if response has already been written, skip.
// 		// if w.Written() {
// 		// 	return
// 		// }

// 		// // if the request is cancelled by the client we should not try
// 		// // to rotate the token since the client would not accept any result.
// 		// if ctx.Context.Req.Context().Err() == context.Canceled {
// 		// 	return
// 		// }

// 		// rotated, err := authTokenService.TryRotateToken(ctx.Req.Context(), token, ctx.RemoteAddr(), ctx.Req.UserAgent())
// 		// if err != nil {
// 		// 	ctx.Logger.Error("Failed to rotate token", "error", err)
// 		// 	return
// 		// }

// 		// if rotated {
// 		// 	WriteSessionCookie(ctx, token.UnhashedToken, setting.LoginMaxLifetimeDays)
// 		// }
// 	}
// }

func WriteSessionCookie(ctx *models.ReqContext, value string, maxLifetimeDays int) {
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
