package Blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Blog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    BlogReq "github.com/flipped-aurora/gin-vue-admin/server/model/Blog/request"
)

type TestService struct {
}

// CreateTest 创建test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService) CreateTest(test *Blog.Test) (err error) {
	err = global.GVA_DB.Create(test).Error
	return err
}

// DeleteTest 删除test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)DeleteTest(test Blog.Test) (err error) {
	err = global.GVA_DB.Delete(&test).Error
	return err
}

// DeleteTestByIds 批量删除test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)DeleteTestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Blog.Test{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTest 更新test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)UpdateTest(test Blog.Test) (err error) {
	err = global.GVA_DB.Save(&test).Error
	return err
}

// GetTest 根据id获取test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)GetTest(id uint) (test Blog.Test, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&test).Error
	return
}

// GetTestInfoList 分页获取test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)GetTestInfoList(info BlogReq.TestSearch) (list []Blog.Test, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Blog.Test{})
    var tests []Blog.Test
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&tests).Error
	return  tests, total, err
}
