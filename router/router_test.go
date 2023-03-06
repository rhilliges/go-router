package router

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {
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
