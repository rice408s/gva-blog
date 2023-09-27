package Blog

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Comment struct{
	global.GVA_MODEL
	ArticleId uint `json:"articleId" form:"articleId" gorm:"column:article_id;comment:;"`
	UserId uint `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:;"`
}

func (Comment) TableName() string {
	return "comment"
}