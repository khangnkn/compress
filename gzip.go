package compress

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
)

func compressWithGzip(logs []dummyLog) (result []byte, err error) {
	b, err := json.Marshal(logs)
	if err != nil {
		return
	}
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	defer writer.Close()
	_, err = writer.Write(b)
	if err != nil {
		return
	}
	writer.Flush()
	return buf.Bytes(), nil
}
