package compress

import (
	"bytes"
	"compress/flate"
	"encoding/json"
)

func compressWithFlate(logs []dummyLog, level int) (result []byte, err error) {
	b, err := json.Marshal(logs)
	if err != nil {
		return
	}
	var buf bytes.Buffer
	writer, err := flate.NewWriter(&buf, level)
	if err != nil {
		return
	}
	defer writer.Close()
	if _, err = writer.Write(b); err != nil {
		return
	}
	writer.Flush()
	return buf.Bytes(), nil
}
