package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func NewWithRequestContext(method, url string, body io.Reader) (res *httptest.ResponseRecorder, context *gin.Context) {
	r := httptest.NewRequest(method, url, body)
	res = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(res)
	context.Request = r
	return res, context
}

func JSON(v interface{}) io.Reader {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		panic(fmt.Errorf("JSON encode error: %w", err))
	}
	return &buf
}

// DecodeJSON is utility for decoding json into v.
func DecodeJSON(r io.Reader, v interface{}) {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		panic(fmt.Errorf("JSON decode error: %w", err))
	}
}
