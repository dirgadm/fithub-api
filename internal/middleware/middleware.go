package middleware

import (
	"context"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dirgadm/fithub-api/internal/common"
	"github.com/dirgadm/fithub-api/pkg/constants"
	"github.com/dirgadm/fithub-api/pkg/ehttp"
	jwtx "github.com/dirgadm/fithub-api/pkg/jwt"
	"github.com/dirgadm/fithub-api/pkg/log"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Middleware defines object for order api custom middleware
type Middleware struct {
	opt common.Options
}

func NewMiddleware(opt common.Options) *Middleware {
	return &Middleware{
		opt: opt,
	}
}

func (m *Middleware) Authorized() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			authorization := ctx.Request().Header.Get("Authorization")
			var match bool
			match, err = regexp.MatchString("^Bearer .+", authorization)
			if err != nil || !match {
				m.opt.Logger.AddMessage(log.DebugLevel, echo.ErrUnauthorized).Print()
				return echo.ErrUnauthorized
			}

			j := jwtx.NewJWT([]byte(m.opt.Config.Jwt.Key))

			tokenStr := strings.Split(authorization, " ")

			var token *jwt.Token
			token, err = j.Parse(tokenStr[1])
			if err != nil {
				m.opt.Logger.AddMessage(log.DebugLevel, echo.ErrUnauthorized).Print()
				return echo.ErrUnauthorized
			}

			var claims *jwtx.UserClaim
			var ok bool
			claims, ok = token.Claims.(*jwtx.UserClaim)
			if !ok {
				m.opt.Logger.AddMessage(log.DebugLevel, echo.ErrUnauthorized).Print()
				return echo.ErrUnauthorized
			}

			expiresAt := claims.ExpiresAt
			if expiresAt <= time.Now().Unix() {
				m.opt.Logger.AddMessage(log.DebugLevel, echo.ErrBadRequest).Print()
				ctx.JSON(http.StatusUnauthorized, ehttp.FormatResponse{
					Code:    http.StatusUnauthorized,
					Status:  "failure",
					Message: "Your token is expired",
				})
				return echo.ErrUnauthorized
			}

			ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), constants.KeyToken, token)))
			ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), constants.KeyUserID, claims.UserId)))
			return next(ctx)
		}
	}
}
