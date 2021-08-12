// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog-backend/app/dao/internal"
)

// commentDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type commentDao struct {
	*internal.CommentDao
}

var (
	// Comment is globally public accessible object for table comment operations.
	Comment commentDao
)

func init() {
	Comment = commentDao{
		internal.NewCommentDao(),
	}
}

// Fill with you ideas below.
