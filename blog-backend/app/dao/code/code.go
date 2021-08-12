package code

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/12-14:35

// 操作成功代码
const (
	UserFound         = 1001
	UserUpdateSuccess = 1002
	UserInsertSuccess = 1003
)

// 操作失败代码
const (
	UserNotFound     = 1101
	UserUpdateFailed = 1102
	UserInsertFailed = 1103
	UserAlreadyExist = 1004
)
