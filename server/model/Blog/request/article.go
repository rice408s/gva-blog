package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Blog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ArticleSearch struct{
    Blog.Article
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartDate  *time.Time  `json:"startDate" form:"startDate"`
    EndDate  *time.Time  `json:"endDate" form:"endDate"`
    // request.PageInfo
    ArticlePageInfo
}



type ArticlePageInfo struct {
    request.PageInfo
    LoginUserId uint `json:"userId" form:"userId"`
    RoleId uint `json:"roleId" form:"roleId"`
}