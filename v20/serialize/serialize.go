package serialize

// 序列化
type I9Serialize interface {
	F8GetCode() uint8
	F8Encode(anyInput any) ([]byte, error)
	F8Decode(s5input []byte, anyOutput any) error
}
