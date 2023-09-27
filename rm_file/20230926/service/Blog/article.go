package Blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Blog"
	BlogReq "github.com/flipped-aurora/gin-vue-admin/server/model/Blog/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ArticleService struct {
}

// CreateArticle 创建Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) CreateArticle(article *Blog.Article) (err error) {
	err = global.GVA_DB.Create(article).Error
	return err
}

// DeleteArticle 删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) DeleteArticle(article Blog.Article) (err error) {
	err = global.GVA_DB.Delete(&article).Error
	return err
}

// DeleteArticleByIds 批量删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) DeleteArticleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Blog.Article{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateArticle 更新Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) UpdateArticle(article Blog.Article) (err error) {
	err = global.GVA_DB.Save(&article).Error
	return err
}

// GetArticle 根据id获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticle(id uint) (article Blog.Article, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&article).Error
	return
}

// GetArticleInfoList 分页获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticleInfoList(info BlogReq.ArticleSearch) (list []Blog.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Blog.Article{})
	var articles []Blog.Article
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&articles).Error
	return articles, total, err
}
