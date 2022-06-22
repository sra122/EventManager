package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app App

func TestApp_main(t *testing.T) {
	main()
}

func TestApp_GetEvent(t *testing.T) {
	main()
	req := httptest.NewRequest(http.MethodGet, "/event/{event_id}", nil)
	w := httptest.NewRecorder()
	app.GetEvent(w, req)
	res := w.Result()
	defer res.Body.Close()
	ioutil.ReadAll(res.Body)
	fmt.Println(res.StatusCode)
}
