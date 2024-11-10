package request

import (
	"net/http"
	"url-shortner/pkg/response"
)

func HandleBody[T any](writer *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		response.Json(*writer, err.Error(), 402)
		return nil, err
	}
	err = isValid(body)
	if err != nil {
		response.Json(*writer, err.Error(), 402)
	}
	return &body, err
}
