package request

import (
	"encoding/json"
	"io"
	"log"
)

func Decode[T any](body io.ReadCloser) (*T, error) {
	var payLoad T
	err := json.NewDecoder(body).Decode(&payLoad)
	if err != nil {
		return nil, err
		log.Println("Error decoding payload: %v", err)

	}
	return &payLoad, nil
}
