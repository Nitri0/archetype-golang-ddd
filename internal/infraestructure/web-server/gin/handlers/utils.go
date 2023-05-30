package gin_handlers

import "encoding/json"

type EmptyObject struct{}

func (eo EmptyObject) MarshalJSON() ([]byte, error) {
	return []byte("{}"), nil
}

type DataWrapper struct {
	Data interface{}
}

func (dw DataWrapper) MarshalJSON() ([]byte, error) {
	if dw.Data == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(dw.Data)
}

type GenericResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    DataWrapper `json:"data"`
	Message string      `json:"message"`
}
