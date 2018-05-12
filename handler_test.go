package lite

import (
	"errors"
	"reflect"
	"testing"
)

func Test_Handler(t *testing.T) {
	t.Run("Given an HTTP handler", func(t *testing.T) {
		handler := NewHandler()
		t.Run("should register module with provided alias", func(t *testing.T) {
			if handler.Use("one", NewBaseModule()) != nil {
				t.Error("should not return an error")
			}
		})
		t.Run("should return an error trying to use the same alias again", func(t *testing.T) {
			if !reflect.DeepEqual(handler.Use("one", NewBaseModule()), errors.New("alias already in use \"one\"")) {
				t.Error("should return \"not unique\" error")
			}
		})
	})
}
