package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/12-15:36

func EncryptPassword(pass string) string {
	l := 16
	password, err := bcrypt.GenerateFromPassword([]byte(pass), l)
	if err != nil {
		log.Fatal("加密密码失败")
	}
	return string(password)
}
