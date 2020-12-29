package common

func Bool(b bool) *bool {
	var ptr = new(bool)
	*ptr = b
	return ptr
}

func String(s string) *string {
	var ptr = new(string)
	*ptr = s
	return ptr
}

func Byte(b byte) *byte {
	var ptr = new(byte)
	*ptr = b
	return ptr
}

func Int(i int) *int {
	var ptr = new(int)
	*ptr = i
	return ptr
}

func Int32(i int32) *int32 {
	var ptr = new(int32)
	*ptr = i
	return ptr
}

func Int64(i int64) *int64 {
	var ptr = new(int64)
	*ptr = i
	return ptr
}

func Float32(f float32) *float32 {
	var ptr = new(float32)
	*ptr = f
	return ptr
}

func Float64(f float64) *float64 {
	var ptr = new(float64)
	*ptr = f
	return ptr
}
