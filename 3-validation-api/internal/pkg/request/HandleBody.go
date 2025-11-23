package request

import (
	"PurpleHW/3-validation-api/internal/pkg/resp"
	"log"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		resp.WriteJSON(w, err, http.StatusBadRequest)
		log.Printf("Error Decode: %v", err)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		log.Printf("Error validating payload: %v", err)
		return nil, err
	}
	return body, nil
}
