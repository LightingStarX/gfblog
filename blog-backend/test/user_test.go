package test

import (
	"blog-backend/app/dao"
	"blog-backend/app/dao/utils"
	"blog-backend/app/model"
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"testing"
)

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/11-17:42

func TestGetUserByUID(t *testing.T) {
	user, err := dao.User.GetUserBuUID(1)
	if err == nil {
		fmt.Println(gconv.Map(user))
	} else {
		fmt.Println(err.Error())
	}
}

func TestGetUserByUIDs(t *testing.T) {
	uids := []int64{1, 2}
	ds, err := dao.User.GetUserByUIDs(uids...)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, user := range ds {
			fmt.Println(gconv.Map(user))
		}
	}
}

func TestInsertUser(t *testing.T) {
	user := model.User{
		UserName: "CamelCodeUser",
		Password: "xxxxxxxxxxxx",
		NickName: "CamelCode",
	}

	ok := dao.User.AddUser(user)
	if ok {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加失败")
	}
}

func TestSliceDelete(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	s = append(s[:1], s[2:]...)
	fmt.Println(s)
}

func TestUserUpdate(t *testing.T) {
	user := model.User{
		NickName: "CamelCode",
		Password: "ashjdhgdja3333hs",
	}

	code := dao.User.UpdateUser(user)
	fmt.Println(code)
}

func TestDeleteUserByUID(t *testing.T) {
	_ = dao.User.DeleteUserByUID(31244288)
}

func TestEncrypt(t *testing.T) {
	password := "asdfasfasf"
	encryptPassword := utils.EncryptPassword(password)
	fmt.Println(encryptPassword)
}
