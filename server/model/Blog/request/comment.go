package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Blog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CommentSearch struct{
	Blog.Comment
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}