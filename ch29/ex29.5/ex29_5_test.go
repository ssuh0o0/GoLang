//ch29/ex29.5/ex29_5_test.go
package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) // 테스트용 "/" 경로 요청 객체 만듦

	mux := MakeWebHandler() // 핸들러 인스턴스의 ServeHTTP() 메서드를 호출하여 요청에 대한 결과를 가져옴
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)     // status가 ok인지 확인
	data, _ := io.ReadAll(res.Body)           // 데이터를 읽어옴
	assert.Equal("Hello World", string(data)) //읽어온 데이터의 타입이 []byte이므로 string으로 타입변환해서 hello world 랑 같은지 확인
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil) // bar 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}
