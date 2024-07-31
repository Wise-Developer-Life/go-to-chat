package job

import "encoding/json"

type Common[T any] struct {
	CreatedAt int64 `json:"created_at"`
	Content   T     `json:"content"`
}

func DecodeJobPayload[T any](payloadByte []byte) (*Common[T], error) {
	var payload Common[T]

	if err := json.Unmarshal(payloadByte, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func EncodeJobPayload(payload *Common[any]) ([]byte, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return payloadBytes, nil
}

type FindMatchJobPayload struct {
	User string `json:"user"`
}
