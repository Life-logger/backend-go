package converter

import (
	"bytes"
	"encoding/json"
	"io"
)

func ConvertPayload(payload map[string]interface{}) io.Reader {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(payload)
	if err != nil {
		panic(err.Error())
	}
	return buffer
}
