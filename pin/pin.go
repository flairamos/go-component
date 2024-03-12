package pin

// 类型转换
// 将基本数据类型转换为指针类型

func I64(i64 int64) *int64 {
	return &i64
}

func I32(i32 int32) *int32 {
	return &i32
}

func I4(i int) *int {
	return &i
}

func Str(str string) *string {
	return &str
}

func II64(i int) *int64 {
	var tmp int64 = int64(i)
	return &tmp
}
func II32(i int) *int32 {
	tmp := int32(i)
	return &tmp
}

func Bool(b bool) *bool {
	return &b
}

func Any[T any](t T) *T {
	return &t
}
