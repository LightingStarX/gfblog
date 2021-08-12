package test

import (
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"testing"
)

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/12-17:14

func TestDbTransaction(t *testing.T) {
	db := g.DB()
	if tx, err := db.Begin(); err != nil {
		fmt.Println(err.Error())
	} else {
		_, err := tx.Save("test", g.Map{
			"c1": 1,
			"c2": 2,
		})
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				fmt.Println("roll back failed")
				return
			}
		}
		err = tx.Commit()
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}

	}
}

func TestTransactionLinkedDelete(t *testing.T) {
	ok := dao.Role.DeleteRole(model.Role{Uid: 12315})
	if ok {
		name, o := dao.Role.GetRoleNameByUid(12315)
		if o {
			fmt.Println(name)
		} else {
			fmt.Println("deleted")
		}
	}
}
