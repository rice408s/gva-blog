import service from '@/utils/request'

// @Tags Test
// @Summary 创建test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test true "创建test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /test/createTest [post]
export const createTest = (data) => {
  return service({
    url: '/test/createTest',
    method: 'post',
    data
  })
}

// @Tags Test
// @Summary 删除test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test true "删除test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test/deleteTest [delete]
export const deleteTest = (data) => {
  return service({
    url: '/test/deleteTest',
    method: 'delete',
    data
  })
}

// @Tags Test
// @Summary 批量删除test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /test/deleteTest [delete]
export const deleteTestByIds = (data) => {
  return service({
    url: '/test/deleteTestByIds',
    method: 'delete',
    data
  })
}

// @Tags Test
// @Summary 更新test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Test true "更新test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /test/updateTest [put]
export const updateTest = (data) => {
  return service({
    url: '/test/updateTest',
    method: 'put',
    data
  })
}

// @Tags Test
// @Summary 用id查询test
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Test true "用id查询test"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /test/findTest [get]
export const findTest = (params) => {
  return service({
    url: '/test/findTest',
    method: 'get',
    params
  })
}

// @Tags Test
// @Summary 分页获取test列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取test列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /test/getTestList [get]
export const getTestList = (params) => {
  return service({
    url: '/test/getTestList',
    method: 'get',
    params
  })
}
