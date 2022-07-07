package pkgMatch

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgMatchReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type MatchCategoryApi struct {
}

var matchCategoryService = service.ServiceGroupApp.PkgmatchServiceGroup.MatchCategoryService


// CreateMatchCategory 创建MatchCategory
// @Tags MatchCategory
// @Summary 创建MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgMatch.MatchCategory true "创建MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /matchCategory/createMatchCategory [post]
func (matchCategoryApi *MatchCategoryApi) CreateMatchCategory(c *gin.Context) {
	var matchCategory pkgMatch.MatchCategory
	_ = c.ShouldBindJSON(&matchCategory)
	if err := matchCategoryService.CreateMatchCategory(matchCategory); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMatchCategory 删除MatchCategory
// @Tags MatchCategory
// @Summary 删除MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgMatch.MatchCategory true "删除MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /matchCategory/deleteMatchCategory [delete]
func (matchCategoryApi *MatchCategoryApi) DeleteMatchCategory(c *gin.Context) {
	var matchCategory pkgMatch.MatchCategory
	_ = c.ShouldBindJSON(&matchCategory)
	if err := matchCategoryService.DeleteMatchCategory(matchCategory); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMatchCategoryByIds 批量删除MatchCategory
// @Tags MatchCategory
// @Summary 批量删除MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /matchCategory/deleteMatchCategoryByIds [delete]
func (matchCategoryApi *MatchCategoryApi) DeleteMatchCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := matchCategoryService.DeleteMatchCategoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMatchCategory 更新MatchCategory
// @Tags MatchCategory
// @Summary 更新MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgMatch.MatchCategory true "更新MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /matchCategory/updateMatchCategory [put]
func (matchCategoryApi *MatchCategoryApi) UpdateMatchCategory(c *gin.Context) {
	var matchCategory pkgMatch.MatchCategory
	_ = c.ShouldBindJSON(&matchCategory)
	if err := matchCategoryService.UpdateMatchCategory(matchCategory); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMatchCategory 用id查询MatchCategory
// @Tags MatchCategory
// @Summary 用id查询MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgMatch.MatchCategory true "用id查询MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /matchCategory/findMatchCategory [get]
func (matchCategoryApi *MatchCategoryApi) FindMatchCategory(c *gin.Context) {
	var matchCategory pkgMatch.MatchCategory
	_ = c.ShouldBindQuery(&matchCategory)
	if rematchCategory, err := matchCategoryService.GetMatchCategory(matchCategory.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rematchCategory": rematchCategory}, c)
	}
}

// GetMatchCategoryList 分页获取MatchCategory列表
// @Tags MatchCategory
// @Summary 分页获取MatchCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgMatchReq.MatchCategorySearch true "分页获取MatchCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /matchCategory/getMatchCategoryList [get]
func (matchCategoryApi *MatchCategoryApi) GetMatchCategoryList(c *gin.Context) {
	var pageInfo pkgMatchReq.MatchCategorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := matchCategoryService.GetMatchCategoryInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
