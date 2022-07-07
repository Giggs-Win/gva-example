import service from '@/utils/request'

// @Tags MatchCategory
// @Summary 创建MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MatchCategory true "创建MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /matchCategory/createMatchCategory [post]
export const createMatchCategory = (data) => {
  return service({
    url: '/matchCategory/createMatchCategory',
    method: 'post',
    data
  })
}

// @Tags MatchCategory
// @Summary 删除MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MatchCategory true "删除MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /matchCategory/deleteMatchCategory [delete]
export const deleteMatchCategory = (data) => {
  return service({
    url: '/matchCategory/deleteMatchCategory',
    method: 'delete',
    data
  })
}

// @Tags MatchCategory
// @Summary 删除MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /matchCategory/deleteMatchCategory [delete]
export const deleteMatchCategoryByIds = (data) => {
  return service({
    url: '/matchCategory/deleteMatchCategoryByIds',
    method: 'delete',
    data
  })
}

// @Tags MatchCategory
// @Summary 更新MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MatchCategory true "更新MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /matchCategory/updateMatchCategory [put]
export const updateMatchCategory = (data) => {
  return service({
    url: '/matchCategory/updateMatchCategory',
    method: 'put',
    data
  })
}

// @Tags MatchCategory
// @Summary 用id查询MatchCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MatchCategory true "用id查询MatchCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /matchCategory/findMatchCategory [get]
export const findMatchCategory = (params) => {
  return service({
    url: '/matchCategory/findMatchCategory',
    method: 'get',
    params
  })
}

// @Tags MatchCategory
// @Summary 分页获取MatchCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MatchCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /matchCategory/getMatchCategoryList [get]
export const getMatchCategoryList = (params) => {
  return service({
    url: '/matchCategory/getMatchCategoryList',
    method: 'get',
    params
  })
}
