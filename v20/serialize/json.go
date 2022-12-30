package serialize

import (
	"encoding/json"
)

type S6Json struct {
}

func F8NewS6Json() *S6Json {
	return &S6Json{}
}

func (this S6Json) F8GetCode() uint8 {
	return 2
}

func (this S6Json) F8Encode(anyInput any) ([]byte, error) {
	return json.Marshal(anyInput)
}

func (this S6Json) F8Decode(s5input []byte, anyOutput any) error {
	return json.Unmarshal(s5input, anyOutput)
}
