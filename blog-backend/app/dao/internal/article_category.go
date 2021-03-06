// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ArticleCategoryDao is the manager for logic model data accessing and custom defined data operations functions management.
type ArticleCategoryDao struct {
	gmvc.M                        // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      articleCategoryColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB                 // DB is the raw underlying database management object.
	Table  string                 // Table is the underlying table name of the DAO.
}

// ArticleCategoryColumns defines and stores column names for table article_category.
type articleCategoryColumns struct {
	Uid       string //
	Name      string // 分类名称
	Enabled   string //
	GmtCreate string //
	GmtUpdate string //
}

// NewArticleCategoryDao creates and returns a new DAO object for table data access.
func NewArticleCategoryDao() *ArticleCategoryDao {
	columns := articleCategoryColumns{
		Uid:       "uid",
		Name:      "name",
		Enabled:   "enabled",
		GmtCreate: "gmt_create",
		GmtUpdate: "gmt_update",
	}
	return &ArticleCategoryDao{
		C:     columns,
		M:     g.DB("default").Model("article_category").Safe(),
		DB:    g.DB("default"),
		Table: "article_category",
	}
}
