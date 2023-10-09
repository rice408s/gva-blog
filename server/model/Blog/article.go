// 自动生成模板Article
package Blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// Article 结构体  Article
type Article struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;"`  //标题 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`  //用户id 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:;"`  //内容 
      Date  *time.Time `json:"date" form:"date" gorm:"column:date;comment:;"`  //发布日期 
      Likes  *int `json:"likes" form:"likes" gorm:"column:likes;comment:;"`  //点赞数 
      Tags  string `json:"tags" form:"tags" gorm:"column:tags;comment:;"`  //标签 
      Category  string `json:"category" form:"category" gorm:"column:category;comment:;"`  //分类 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:;"`  //状态 
}


// TableName Article Article自定义表名 article
func (Article) TableName() string {
  return "article"
}

