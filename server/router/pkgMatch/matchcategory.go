package pkgMatch

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MatchCategoryRouter struct {
}

// InitMatchCategoryRouter 初始化 MatchCategory 路由信息
func (s *MatchCategoryRouter) InitMatchCategoryRouter(Router *gin.RouterGroup) {
	matchCategoryRouter := Router.Group("matchCategory").Use(middleware.OperationRecord())
	matchCategoryRouterWithoutRecord := Router.Group("matchCategory")
	var matchCategoryApi = v1.ApiGroupApp.PkgmatchApiGroup.MatchCategoryApi
	{
		matchCategoryRouter.POST("createMatchCategory", matchCategoryApi.CreateMatchCategory)   // 新建MatchCategory
		matchCategoryRouter.DELETE("deleteMatchCategory", matchCategoryApi.DeleteMatchCategory) // 删除MatchCategory
		matchCategoryRouter.DELETE("deleteMatchCategoryByIds", matchCategoryApi.DeleteMatchCategoryByIds) // 批量删除MatchCategory
		matchCategoryRouter.PUT("updateMatchCategory", matchCategoryApi.UpdateMatchCategory)    // 更新MatchCategory
	}
	{
		matchCategoryRouterWithoutRecord.GET("findMatchCategory", matchCategoryApi.FindMatchCategory)        // 根据ID获取MatchCategory
		matchCategoryRouterWithoutRecord.GET("getMatchCategoryList", matchCategoryApi.GetMatchCategoryList)  // 获取MatchCategory列表
	}
}
