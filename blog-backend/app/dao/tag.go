// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog-backend/app/dao/internal"
)

// tagDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type tagDao struct {
	*internal.TagDao
}

var (
	// Tag is globally public accessible object for table tag operations.
	Tag tagDao
)

func init() {
	Tag = tagDao{
		internal.NewTagDao(),
	}
}

// Fill with you ideas below.