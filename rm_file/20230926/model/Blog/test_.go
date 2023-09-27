// 自动生成模板Test
package Blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// test 结构体  Test
type Test struct {
      global.GVA_MODEL
      TestName  string `json:"testName" form:"testName" gorm:"column:test_name;comment:;"`  //测试 
}


// TableName test Test自定义表名 Test
func (Test) TableName() string {
  return "Test"
}

