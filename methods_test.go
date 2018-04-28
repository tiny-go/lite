package lite

import (
	"reflect"
	"testing"
)

func Test_Add(t *testing.T) {
	t.Run("Given Methods (list of HTTP methods)", func(t *testing.T) {
		t.Run("method Add should add strings to a slice", func(t *testing.T) {
			list := &Methods{}
			list.Add("GET")
			list.Add("POST")
			list.Add("PUT")
			if !reflect.DeepEqual(list, &Methods{"GET", "POST", "PUT"}) {
				t.Errorf("list of methods is not valid: %v", list)
			}
		})
		t.Run("method Add should ignore duplicates", func(t *testing.T) {
			list := &Methods{}
			list.Add("GET")
			list.Add("GET")
			list.Add("PUT")
			if !reflect.DeepEqual(list, &Methods{"GET", "PUT"}) {
				t.Errorf("list of methods is not valid: %v", list)
			}
		})
	})
}

func Test_Join(t *testing.T) {
	t.Run("Given Methods (list of HTTP methods)", func(t *testing.T) {
		t.Run("method Join should convert slice to a string", func(t *testing.T) {
			list := &Methods{}
			list.Add("GET")
			list.Add("POST")
			list.Add("PUT")
			if list.Join() != "GET,POST,PUT" {
				t.Errorf("the result was expected to be \"GET,POST,PUT\" but was %q", list.Join())
			}
		})
	})
}

func Test_Empty(t *testing.T) {
	t.Run("Given Methods (list of HTTP methods)", func(t *testing.T) {
		t.Run("method Empty should return true if list is empty", func(t *testing.T) {
			list := &Methods{}
			if list.Empty() != true {
				t.Error("the result was expected to be true")
			}
		})
		t.Run("method Empty should return false if list is not empty", func(t *testing.T) {
			list := &Methods{}
			list.Add("GET")
			if list.Empty() != false {
				t.Error("the result was expected to be false")
			}
		})
	})
}
