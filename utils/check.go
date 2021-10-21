package utils

func If(cond bool, trueVal, falseVal interface{}) interface{} {
	if cond {
		return trueVal
	}
	return falseVal
}
