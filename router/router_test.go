package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisteringHandlerForPath(t *testing.T) {
	router := NewRouter()
	handler1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	router.Register("path", handler1)
	handler2 := router.GetHandler("path")

	handler1Ref := fmt.Sprintf("%v", handler1)
	handler2Ref := fmt.Sprintf("%v", handler2)
	if handler1Ref != handler2Ref {
		t.Error("expected same handler")
	}
}

func TestCallCorrectHandler(t *testing.T) {
	router := NewRouter()
	called := false
	handler1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})
	router.Register("path/to/handler", handler1)
	request := httptest.NewRequest("GET", "http://domain.org/path/to/handler", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	if called != true {
		t.Error("handler was not called")
	}
}

func TestReturn404WhenNoHandlerIsFound(t *testing.T) {
	router := NewRouter()
	request := httptest.NewRequest("GET", "http://domain.org/path/to/handler", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	if writer.Code != http.StatusNotFound {
		t.Error("expected response code is http.StatusNotFound")
	}
}

func TestRegisteringHandlerForPathWithVariable(t *testing.T) {
	router := NewRouter()
	handler1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	router.Register("path/{variable}", handler1)
	handler2 := router.GetHandler("path/value")

	handler1Ref := fmt.Sprintf("%v", handler1)
	handler2Ref := fmt.Sprintf("%v", handler2)
	if handler1Ref != handler2Ref {
		t.Error("expected same handler")
	}
}
