package lite

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	_ "github.com/tiny-go/codec/driver/json"
	"github.com/tiny-go/middleware"
)

type mockController struct {
	mw.Controller
	ShouldFail bool
}

func newPassController() *mockController {
	return &mockController{mw.NewBaseController(), false}
}

func newFailController() *mockController {
	return &mockController{mw.NewBaseController(), true}
}

func (c *mockController) Init() error { return nil }

func (c *mockController) Get(_ context.Context, pk string) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("single GET error")
	}
	return pk, nil
}

func (c *mockController) GetAll(context.Context) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("plural GET error")
	}
	return map[string]interface{}{"foo": "bar"}, nil
}

func (c *mockController) Post(_ context.Context, pk string, f func(v interface{}) error) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("single POST error")
	}
	data := map[string]interface{}{"pk": pk}
	return data, f(&data)
}

func (c *mockController) PostAll(_ context.Context, f func(v interface{}) error) (interface{}, error) {
	if c.ShouldFail {
		return nil, errors.New("plural POST error")
	}
	data := make(map[string]interface{})
	return data, f(&data)
}

func Test_Handler(t *testing.T) {
	t.Run("Given an HTTP handler with registered module", func(t *testing.T) {
		module := NewBaseModule()
		module.Register("pass", newPassController())
		module.Register("fail", newFailController())
		handler := NewHandler()
		handler.Use("test", module)
		ts := httptest.NewServer(handler)
		defer ts.Close()

		type testCase struct {
			title   string
			request *http.Request
			code    int
			header  http.Header
			body    string
		}
		testCases := []testCase{
			{
				title: "plural OPTIONS request",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodOptions, ts.URL+"/test/pass", nil)
					return r
				}(),
				header: http.Header{"Access-Control-Allow-Methods": []string{"GET,POST"}},
				code:   http.StatusOK,
			},
			{
				title: "single OPTIONS request",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodOptions, ts.URL+"/test/pass/abcd", nil)
					return r
				}(),
				header: http.Header{"Access-Control-Allow-Methods": []string{"GET,POST"}},
				code:   http.StatusOK,
			},
			{
				title: "plural GET with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, ts.URL+"/test/pass", nil)
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\"}\n",
			},
			{
				title: "single GET with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, ts.URL+"/test/pass/abcd", nil)
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "\"abcd\"\n",
			},
			{
				title: "plural POST with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPost, ts.URL+"/test/pass", strings.NewReader("{\"foo\":\"bar\"}"))
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\"}\n",
			},
			{
				title: "single POST with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPost, ts.URL+"/test/pass/abcd", strings.NewReader("{\"foo\":\"bar\"}"))
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\",\"pk\":\"abcd\"}\n",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.title, func(t *testing.T) {
				client := &http.Client{}
				res, err := client.Do(tc.request)
				if err != nil {
					t.Error("should not return an error")
				}
				if res.StatusCode != tc.code {
					t.Errorf("status code %d was expected to be %d", res.StatusCode, tc.code)
				}
				body, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					t.Error("unable to read response body")
				}
				if string(body) != tc.body {
					t.Errorf("the output %q was expected to be %q", body, tc.body)
				}
				for key, val := range tc.header {
					if _, ok := res.Header[key]; !ok {
						t.Errorf("header %q is missing", key)
					} else if !reflect.DeepEqual(res.Header[key], val) {
						t.Errorf("unexpected value %q for header %q", val, key)
					}
				}
			})
		}
	})
}
