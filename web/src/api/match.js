import service from '@/utils/request'

// @Tags Match
// @Summary 创建Match
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Match true "创建Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /match/createMatch [post]
export const createMatch = (data) => {
  return service({
    url: '/match/createMatch',
    method: 'post',
    data
  })
}

// @Tags Match
// @Summary 删除Match
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Match true "删除Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /match/deleteMatch [delete]
export const deleteMatch = (data) => {
  return service({
    url: '/match/deleteMatch',
    method: 'delete',
    data
  })
}

// @Tags Match
// @Summary 删除Match
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /match/deleteMatch [delete]
export const deleteMatchByIds = (data) => {
  return service({
    url: '/match/deleteMatchByIds',
    method: 'delete',
    data
  })
}

// @Tags Match
// @Summary 更新Match
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Match true "更新Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /match/updateMatch [put]
export const updateMatch = (data) => {
  return service({
    url: '/match/updateMatch',
    method: 'put',
    data
  })
}

// @Tags Match
// @Summary 用id查询Match
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Match true "用id查询Match"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /match/findMatch [get]
export const findMatch = (params) => {
  return service({
    url: '/match/findMatch',
    method: 'get',
    params
  })
}

// @Tags Match
// @Summary 分页获取Match列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Match列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /match/getMatchList [get]
export const getMatchList = (params) => {
  return service({
    url: '/match/getMatchList',
    method: 'get',
    params
  })
}
