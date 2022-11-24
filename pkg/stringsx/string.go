package stringsx

// StringOr 返回第一个不是空的字符串，都为空，返回空
func StringOr(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}
