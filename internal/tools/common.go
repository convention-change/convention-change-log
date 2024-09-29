package tools

import (
	"bytes"
	"encoding/json"
)

func JsonMarshalBeauty(data any) (string, error) {
	dataBytes, errMarshal := json.Marshal(data)
	if errMarshal != nil {
		return "", errMarshal
	}
	var str bytes.Buffer
	errJson := json.Indent(&str, dataBytes, "", "  ")
	if errJson != nil {
		return "", errJson
	}
	return str.String(), nil
}
