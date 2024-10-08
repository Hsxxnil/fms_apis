package middleware

import (
	"net/http"

	"fms/config"
	"fms/internal/interactor/pkg/jwx"
	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"

	"github.com/gin-gonic/gin"
)

func Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		j := &jwx.JWE{
			PrivateKey: config.AccessPrivateKey,
			Token:      ctx.GetHeader("Authorization"),
		}

		if len(j.Token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, code.GetCodeMessage(code.JWTRejected, "AccessToken is null."))
			return
		}

		j, err := j.Verify()
		if err != nil {
			log.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, code.GetCodeMessage(code.JWTRejected, "AccessToken is error."))
			return
		}

		ctx.Set("user_id", j.Other["user_id"])
		ctx.Set("role", j.Other["role"])
		ctx.Set("name", j.Other["name"])
		ctx.Set("fleet_id", j.Other["fleet_id"])
		ctx.Next()
	}
}
