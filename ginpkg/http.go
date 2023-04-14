package ginpkg

import (
	"github.com/gin-gonic/gin"
)

func RouteParamByQuery[T any](f func(ctx *gin.Context, param T)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var param T
		if err := ctx.ShouldBindQuery(&param); err != nil {
			BadRequestResponse(ctx, "请求参数错误", err.Error())
			return
		}
		f(ctx, param)
	}
}

func RouteParamByJson[T any](f func(ctx *gin.Context, param T)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var param T
		if err := ctx.ShouldBindJSON(&param); err != nil {
			BadRequestResponse(ctx, "请求参数错误", err.Error())
			return
		}
		f(ctx, param)
	}
}
