package pkgMatch

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MatchRouter struct {
}

// InitMatchRouter 初始化 Match 路由信息
func (s *MatchRouter) InitMatchRouter(Router *gin.RouterGroup) {
	matchRouter := Router.Group("match").Use(middleware.OperationRecord())
	matchRouterWithoutRecord := Router.Group("match")
	var matchApi = v1.ApiGroupApp.PkgmatchApiGroup.MatchApi
	{
		matchRouter.POST("createMatch", matchApi.CreateMatch)   // 新建Match
		matchRouter.DELETE("deleteMatch", matchApi.DeleteMatch) // 删除Match
		matchRouter.DELETE("deleteMatchByIds", matchApi.DeleteMatchByIds) // 批量删除Match
		matchRouter.PUT("updateMatch", matchApi.UpdateMatch)    // 更新Match
	}
	{
		matchRouterWithoutRecord.GET("findMatch", matchApi.FindMatch)        // 根据ID获取Match
		matchRouterWithoutRecord.GET("getMatchList", matchApi.GetMatchList)  // 获取Match列表
	}
}
