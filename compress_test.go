package compress

import (
	"compress/flate"
	"compress/lzw"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
	"unsafe"

	"google.golang.org/protobuf/proto"
)

var testData = make([]dummyLog, 0)
var result []byte
var protoResult *Payload

func TestProtobuf(t *testing.T) {
	r := serializeWithProtobuf(testData)
	t.Logf("protobuf serialized data to %d bytes", proto.Size(r))
}

func TestJSON(t *testing.T) {
	b, err := json.Marshal(testData)
	if err != nil {
		t.Error(err)
	}
	t.Logf("json serialized data to %d bytes", len(b))
}

func TestLzw(t *testing.T) {
	r, err := compressWithLzw(testData, lzw.LSB)
	if err != nil {
		t.Error(err)
	}
	t.Logf("lzw compressed data to %d bytes", len(r))
}

func TestFlate(t *testing.T) {
	r, err := compressWithFlate(testData, flate.DefaultCompression)
	if err != nil {
		t.Error(err)
	}
	t.Logf("flate compressed data to %d bytes", len(r))
}

func TestGzip(t *testing.T) {
	r, err := compressWithGzip(testData)
	if err != nil {
		t.Error(err)
	}
	t.Logf("gzip compressed data to %d bytes", len(r))
}

func TestZlib(t *testing.T) {
	r, err := compressWithZlib(testData)
	if err != nil {
		t.Error(err)
	}
	t.Logf("zlib compressed data to %d bytes", len(r))
}

// Benchmarking to get the average performance of each method.
func BenchmarkProtobuf(b *testing.B) {
	var payload *Payload
	for i := 0; i < b.N; i++ {
		payload = serializeWithProtobuf(testData)
	}
	protoResult = payload
	b.ReportAllocs()
}

func BenchmarkJSON(b *testing.B) {
	var r []byte
	var err error
	for i := 0; i < b.N; i++ {
		r, err = json.Marshal(testData)
		if err != nil {
			b.Error(err)
		}
	}
	result = r
	b.ReportAllocs()
}

func BenchmarkLzw(b *testing.B) {
	var r []byte
	var err error
	for i := 0; i < b.N; i++ {
		r, err = compressWithLzw(testData, lzw.LSB)
		if err != nil {
			b.Error(err)
		}
	}
	result = r
	b.ReportAllocs()
}

func BenchmarkFlate(b *testing.B) {
	var r []byte
	var err error
	for i := 0; i < b.N; i++ {
		r, err = compressWithFlate(testData, flate.BestCompression)
		if err != nil {
			b.Error(err)
		}
	}
	result = r
	b.ReportAllocs()
}

func BenchmarkGzip(b *testing.B) {
	var r []byte
	var err error
	for i := 0; i < b.N; i++ {
		r, err = compressWithGzip(testData)
		if err != nil {
			b.Error(err)
		}
	}
	result = r
	b.ReportAllocs()
}

func BenchmarkZlib(b *testing.B) {
	var r []byte
	var err error
	for i := 0; i < b.N; i++ {
		r, err = compressWithZlib(testData)
		if err != nil {
			b.Error(err)
		}
	}
	result = r
	b.ReportAllocs()
}

func init() {
	randString := func(n int) string {
		const (
			letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
			letterIdxBits = 6                    // 6 bits to represent a letter index
			letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
			letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
		)
		var src = rand.NewSource(time.Now().UnixNano())
		b := make([]byte, n)
		// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
		for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
			if remain == 0 {
				cache, remain = src.Int63(), letterIdxMax
			}
			if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
				b[i] = letterBytes[idx]
				i--
			}
			cache >>= letterIdxBits
			remain--
		}

		return *(*string)(unsafe.Pointer(&b))
	}

	var batch = os.Getenv("TOTAL_RECORDS")
	log.Printf("generating test data of %s records\n", batch)
	total, err := strconv.Atoi(batch)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < total; i++ {
		var a = dummyLog{
			ServiceName:   randString(10),
			APIName:       randString(20),
			CorrelationID: randString(20),
			ReturnCode:    -234,
			ExecutionTime: 23423,
			TemplateID:    1,
			Data:          randString(1000),
		}
		testData = append(testData, a)
	}
}
