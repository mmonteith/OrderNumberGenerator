package json

import "encoding/json"

type Encoder struct{}

func (je Encoder) MarshalIndent(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}

func (je Encoder) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (je Encoder) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
