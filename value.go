package interface_value

// 将interface{}转换为int，如果转换失败返回默认值
func Int(raw interface{}, val int) int {
	if raw != nil {
		switch raw.(type) {
		case int:
			return raw.(int)
		case int32:
			return int(raw.(int32))
		case int16:
			return int(raw.(int16))
		case int8:
			return int(raw.(int8))
		}
	}

	return val
}
