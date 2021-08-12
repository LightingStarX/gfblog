package utils

import (
	"blog-backend/app/dao/utils"
	"blog-backend/app/model"
	"fmt"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"testing"
)

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/11-19:14

func TestFilterEmptyFieldInMap(t *testing.T) {
	user := model.User{
		Uid:      111111111111,
		Birthday: gtime.Now(),
	}

	m := gconv.Map(user)
	fm, b := utils.FilterEmptyFieldInMap(m)
	if b {
		fmt.Println(fm)
	} else {
		fmt.Println("failed")
	}
}
