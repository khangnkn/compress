package compress

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
)

func compressWithZlib(logs []dummyLog) ([]byte, error) {
	b, err := json.Marshal(logs)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	_, err = writer.Write(b)
	if err != nil {
		return nil, err
	}
	writer.Flush()
	return buf.Bytes(), nil
}
