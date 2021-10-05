// nullable 패키지는 어떤 nil이 아닌 값을 nil이 될 수도 있는 타입으로 표현하기 위해
// 어떤 nil이 아닌 값을 포인터 변수로 바꾸어주는 라이브러리입니다.
// 즉, Int(1)을 호출하면 1이라는 값이 담긴 int 포인터를 반환합니다.

package nullable

// Int 함수는 전달된 int 값을 가지는 int 포인터를 반환합니다
func Int(num int) *int {
	return &num
}

// Float64 함수는 전달된 float64 값을 가지는 float64 포인터를 반환합니다
func Float64(num float64) *float64 {
	return &num
}

// Float32 함수는 전달된 float32 값을 가지는 float32 포인터를 반환합니다
func Float32(num float32) *float32 {
	return &num
}

// String 함수는 전달된 string 값을 가지는 string 포인터를 반환합니다
func String(str string) *string {
	return &str
}

// Bool 함수는 전달된 bool 값을 가지는 bool 포인터를 반환합니다
func Bool(b bool) *bool {
	return &b
}
