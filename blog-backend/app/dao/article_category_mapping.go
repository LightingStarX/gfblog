// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog-backend/app/dao/internal"
)

// articleCategoryMappingDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type articleCategoryMappingDao struct {
	*internal.ArticleCategoryMappingDao
}

var (
	// ArticleCategoryMapping is globally public accessible object for table article_category_mapping operations.
	ArticleCategoryMapping articleCategoryMappingDao
)

func init() {
	ArticleCategoryMapping = articleCategoryMappingDao{
		internal.NewArticleCategoryMappingDao(),
	}
}

// Fill with you ideas below.
