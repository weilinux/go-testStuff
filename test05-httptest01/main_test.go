package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	//NewRecorder 方法用来创建 http 的响应体。返回的类型是 *httptest.ResponseRecorder ，
	//包含接口返回信息，等价于 http.ResponseWriter。
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Ok response is expected")
	assert.Equal(t, "hello test", response.Body.String(), "Incorrect response body")
}
