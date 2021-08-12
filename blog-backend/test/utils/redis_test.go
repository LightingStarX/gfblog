package utils

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"testing"
)

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/11-21:02

func TestRedisConnection(t *testing.T) {
	doVar, err := g.Redis("user_cache").DoVar("GET", "test")
	if err != nil {
		fmt.Println("1", err.Error())
	} else {
		fmt.Println("2", doVar.String())
	}

}
