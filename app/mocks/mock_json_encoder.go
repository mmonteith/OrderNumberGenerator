package mocks

type JsonEncoder struct {
	MarshalResponse    []byte
	MarshalResponseErr error

	UnmarshalResponseErr error
}

func (je JsonEncoder) Marshal(v interface{}) ([]byte, error) {
	return je.MarshalResponse, je.MarshalResponseErr
}

func (je JsonEncoder) Unmarshal(data []byte, v interface{}) error {
	return je.UnmarshalResponseErr
}
