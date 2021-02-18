package compress

import (
	"bytes"
	"compress/lzw"
	"encoding/json"
)

func compressWithLzw(logs []dummyLog, order lzw.Order) ([]byte, error) {
	b, err := json.Marshal(logs)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	writer := lzw.NewWriter(&buf, order, 8)
	defer writer.Close()
	_, err = writer.Write(b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
