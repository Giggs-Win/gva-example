package pkgMatch

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch"
	pkgMatchReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgMatch/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MatchApi struct {
}

var matchService = service.ServiceGroupApp.PkgmatchServiceGroup.MatchService

// CreateMatch 创建Match
// @Param data body pkgMatch.Match true "创建Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /match/createMatch [post]
func (matchApi *MatchApi) CreateMatch(c *gin.Context) {
	var match pkgMatch.Match
	_ = c.ShouldBindJSON(&match)
	if err := matchService.CreateMatch(match); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMatch 删除Match
// @Param data body pkgMatch.Match true "删除Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /match/deleteMatch [delete]
func (matchApi *MatchApi) DeleteMatch(c *gin.Context) {
	var match pkgMatch.Match
	_ = c.ShouldBindJSON(&match)
	if err := matchService.DeleteMatch(match); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMatchByIds 批量删除Match
// @Param data body request.IdsReq true "批量删除Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /match/deleteMatchByIds [delete]
func (matchApi *MatchApi) DeleteMatchByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := matchService.DeleteMatchByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMatch 更新Match
// @Param data body pkgMatch.Match true "更新Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /match/updateMatch [put]
func (matchApi *MatchApi) UpdateMatch(c *gin.Context) {
	var match pkgMatch.Match
	_ = c.ShouldBindJSON(&match)
	if err := matchService.UpdateMatch(match); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMatch 用id查询Match
// @Param data query pkgMatch.Match true "用id查询Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /match/findMatch [get]
func (matchApi *MatchApi) FindMatch(c *gin.Context) {
	var match pkgMatch.Match
	const PV int = 1
	_ = c.ShouldBindQuery(&match)
	if rematch, err := matchService.GetMatch(match.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rematch": rematch}, c)
	}
}

// GetMatchList 分页获取Match列表
// @Param data query pkgMatchReq.MatchSearch true "分页获取Match列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /match/getMatchList [get]
func (matchApi *MatchApi) GetMatchList(c *gin.Context) {
	var pageInfo pkgMatchReq.MatchSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := matchService.GetMatchInfoList(pageInfo); err != nil {
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
