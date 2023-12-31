package Blog

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Blog"
	BlogReq "github.com/flipped-aurora/gin-vue-admin/server/model/Blog/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleApi struct {
}

var articleService = service.ServiceGroupApp.BlogServiceGroup.ArticleService

// CreateArticle 创建Article
// @Tags Article
// @Summary 创建Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Blog.Article true "创建Article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /article/createArticle [post]
func (articleApi *ArticleApi) CreateArticle(c *gin.Context) {
	// 定义 article 变量并绑定 JSON 数据到该变量
	var article Blog.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 定义验证规则
	verify := utils.Rules{
		"Title":  {utils.NotEmpty()},
		"UserId":  {utils.NotEmpty()},
		"Content": {utils.NotEmpty()},
		"Date":    {utils.NotEmpty()},
		// "Likes":{utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
	}

	// 验证 article 是否符合规则
	if err := utils.Verify(article, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 调用 articleService 的 CreateArticle 函数创建文章
	if err := articleService.CreateArticle(&article); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArticle 删除Article
// @Tags Article
// @Summary 删除Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Blog.Article true "删除Article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /article/deleteArticle [delete]
func (articleApi *ArticleApi) DeleteArticle(c *gin.Context) {
	var article Blog.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleService.DeleteArticle(article); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArticleByIds 批量删除Article
// @Tags Article
// @Summary 批量删除Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /article/deleteArticleByIds [delete]
func (articleApi *ArticleApi) DeleteArticleByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleService.DeleteArticleByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArticle 更新Article
// @Tags Article
// @Summary 更新Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Blog.Article true "更新Article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /article/updateArticle [put]
func (articleApi *ArticleApi) UpdateArticle(c *gin.Context) {
	var article Blog.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Title":  {utils.NotEmpty()},
		"UserId":  {utils.NotEmpty()},
		"Content": {utils.NotEmpty()},
		"Date":    {utils.NotEmpty()},
		"Likes":   {utils.NotEmpty()},
		"Status":  {utils.NotEmpty()},
	}
	if err := utils.Verify(article, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleService.UpdateArticle(article); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArticle 用id查询Article
// @Tags Article
// @Summary 用id查询Article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Blog.Article true "用id查询Article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /article/findArticle [get]
func (articleApi *ArticleApi) FindArticle(c *gin.Context) {
	var article Blog.Article
	err := c.ShouldBindQuery(&article)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rearticle, err := articleService.GetArticle(article.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rearticle": rearticle}, c)
	}
}

// GetArticleList 分页获取Article列表
// @Tags Article
// @Summary 分页获取Article列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BlogReq.ArticleSearch true "分页获取Article列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /article/getArticleList [get]
func (articleApi *ArticleApi) GetArticleList(c *gin.Context) {
	var pageInfo BlogReq.ArticleSearch
	err := c.ShouldBindQuery(&pageInfo)
	fmt.Println(pageInfo.LoginUserId)
	fmt.Println(pageInfo.RoleId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := articleService.GetArticleInfoList(pageInfo); err != nil {
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

//获取所有文章列表
func (articleApi *ArticleApi) GetArticleListAll(c * gin.Context){
	var pageInfo BlogReq.ArticleSearch
	err := c.ShouldBindQuery(&pageInfo)
	fmt.Println(pageInfo.LoginUserId)
	fmt.Println(pageInfo.RoleId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := articleService.GetArticleInfoList(pageInfo); err != nil {
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