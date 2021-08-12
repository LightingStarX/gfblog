package test

import (
	"blog-backend/app/dao"
	"fmt"
	"testing"
)

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/12-18:00

func TestUserRoleMapping(t *testing.T) {
	ok := dao.UserRoleMapping.AddUserRoleMapping(1, 2)
	if ok {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}
