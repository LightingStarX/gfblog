package valid

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/12-12:54

var (
	UserAddingRules = map[string]string{
		"userName": "required|length:8,18",          // 用户名长度限制在[3,18]之间
		"password": "required|length:8,16|password", // 用户密码长度限制再[8,16]之间
		"nickName": "required|length:5,100",         // 用户昵称长度限制在[5,100]之间
		"birthday": "date",
	}

	UserAddingMsg = map[string]interface{}{
		"userName": map[string]string{
			"required": "必须提供用户名UserName",
			"length":   "用户名长度限制在[3,18]之间",
		},
		"password": map[string]string{
			"required": "必须提供密码Password",
			"length":   "密码长度必须在[8,16]之间",
			"password": "格式必须是密码格式",
		},
		"nickName": map[string]string{
			"required": "必须提供昵称",
			"length":   "昵称长度必须在[5,100]之间",
		},
		"birthday": "生日必须为合法日期",
	}
)

var (
	UserUpdateRules = map[string]string{
		"uid":           "min:1",
		"roleUid":       "min:1",
		"userName":      "length:8,18",
		"password":      "password",
		"nickName":      "length:5,100",
		"birthday":      "date",
		"indexUrl":      "",
		"lastLoginIp":   "",
		"enabled":       "boolean",
		"accountLevel":  "min:0",
		"totalExp":      "min:0",
		"totalArticle":  "min:0",
		"totalStar":     "min:0",
		"totalCollect":  "min:0",
		"totalComment":  "min:0",
		"totalChar":     "min:0",
		"totalFollower": "min:0",
		"gmtRegister":   "date",
		"gmtLastLogin":  "date",
		"gmtUpdate":     "date",
	}
)
