package utils

import "reflect"

// @Author: OxCAFFEE
// @Github: https://github.com/OxCaffee
// @Email: wwh2021@mail.ustc.edu.cn
// @Date: 2021/8/11-19:09

// FilterEmptyFieldInMap 过滤给定map中的空值字段，防止利用空值字段更新数据库
func FilterEmptyFieldInMap(m map[string]interface{}) (map[string]interface{}, bool) {
	if m == nil {
		return nil, false
	}

	filtered := make(map[string]interface{}, 0)
	for k, v := range m {
		val := reflect.ValueOf(v)
		if val.IsValid() && !val.IsZero() {
			filtered[k] = v
		}
	}
	return filtered, true
}
