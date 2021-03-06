// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ArticleExtDao is the manager for logic model data accessing and custom defined data operations functions management.
type ArticleExtDao struct {
	gmvc.M                   // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      articleExtColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB            // DB is the raw underlying database management object.
	Table  string            // Table is the underlying table name of the DAO.
}

// ArticleExtColumns defines and stores column names for table article_ext.
type articleExtColumns struct {
	Uid        string //
	ArticleUid string // 文章uid
	RawHtml    string // 文章原始html
	RawEncrypt string // 文章加密后内容
	GmtCreate  string //
	GmtUpdate  string //
}

// NewArticleExtDao creates and returns a new DAO object for table data access.
func NewArticleExtDao() *ArticleExtDao {
	columns := articleExtColumns{
		Uid:        "uid",
		ArticleUid: "article_uid",
		RawHtml:    "raw_html",
		RawEncrypt: "raw_encrypt",
		GmtCreate:  "gmt_create",
		GmtUpdate:  "gmt_update",
	}
	return &ArticleExtDao{
		C:     columns,
		M:     g.DB("default").Model("article_ext").Safe(),
		DB:    g.DB("default"),
		Table: "article_ext",
	}
}
