package Blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

// InitTestRouter 初始化 test 路由信息
func (s *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test").Use(middleware.OperationRecord())
	testRouterWithoutRecord := Router.Group("test")
	var testApi = v1.ApiGroupApp.BlogApiGroup.TestApi
	{
		testRouter.POST("createTest", testApi.CreateTest)   // 新建test
		testRouter.DELETE("deleteTest", testApi.DeleteTest) // 删除test
		testRouter.DELETE("deleteTestByIds", testApi.DeleteTestByIds) // 批量删除test
		testRouter.PUT("updateTest", testApi.UpdateTest)    // 更新test
	}
	{
		testRouterWithoutRecord.GET("findTest", testApi.FindTest)        // 根据ID获取test
		testRouterWithoutRecord.GET("getTestList", testApi.GetTestList)  // 获取test列表
	}
}
