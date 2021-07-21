package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"src/github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)
}

func TestJsonHandler2(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/students/1", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("aaa", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/students/2", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("bbb", student.Name)
}

func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(`{"Id:0,"Name:"ccc","Age":15,"Score":78}`))

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/students/3", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ccc", student.Name)
}
