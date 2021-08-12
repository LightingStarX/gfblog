// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// TestDao is the manager for logic model data accessing and custom defined data operations functions management.
type TestDao struct {
	gmvc.M             // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      testColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB      // DB is the raw underlying database management object.
	Table  string      // Table is the underlying table name of the DAO.
}

// TestColumns defines and stores column names for table test.
type testColumns struct {
	C1 string //
	C2 string //
}

// NewTestDao creates and returns a new DAO object for table data access.
func NewTestDao() *TestDao {
	columns := testColumns{
		C1: "c1",
		C2: "c2",
	}
	return &TestDao{
		C:     columns,
		M:     g.DB("default").Model("test").Safe(),
		DB:    g.DB("default"),
		Table: "test",
	}
}
