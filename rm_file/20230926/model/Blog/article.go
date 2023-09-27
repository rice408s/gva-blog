// 自动生成模板Article
package Blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// Article 结构体  Article
type Article struct {
      global.GVA_MODEL
      Titile  string `json:"titile" form:"titile" gorm:"column:titile;comment:;"`  //标题 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`  //用户id 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:;"`  //内容 
}


// TableName Article Article自定义表名 article
func (Article) TableName() string {
  return "article"
}

