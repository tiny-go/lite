package lite

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	_ "github.com/tiny-go/codec/driver/json"
	_ "github.com/tiny-go/codec/driver/xml"
)

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
				header: http.Header{"Access-Control-Allow-Methods": []string{"GET,POST,PATCH,PUT,DELETE"}},
				code:   http.StatusOK,
			},
			{
				title: "single OPTIONS request",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodOptions, ts.URL+"/test/pass/abcd", nil)
					return r
				}(),
				header: http.Header{"Access-Control-Allow-Methods": []string{"GET,POST,PATCH,PUT,DELETE"}},
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
			{
				title: "plural PATCH with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPatch, ts.URL+"/test/pass", strings.NewReader("{\"foo\":\"bar\"}"))
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\"}\n",
			},
			{
				title: "single PATCH with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPatch, ts.URL+"/test/pass/abcd", strings.NewReader("{\"foo\":\"bar\"}"))
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\",\"pk\":\"abcd\"}\n",
			},
			{
				title: "plural PUT with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPut, ts.URL+"/test/pass", strings.NewReader("{\"foo\":\"bar\"}"))
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\"}\n",
			},
			{
				title: "single PUT with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPut, ts.URL+"/test/pass/abcd", strings.NewReader("{\"foo\":\"bar\"}"))
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\",\"pk\":\"abcd\"}\n",
			},
			{
				title: "plural DELETE with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodDelete, ts.URL+"/test/pass", strings.NewReader("{\"foo\":\"bar\"}"))
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\"}\n",
			},
			{
				title: "single DELETE with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodDelete, ts.URL+"/test/pass/abcd", nil)
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Accept", "application/json")
					return r
				}(),
				code: http.StatusOK,
				body: "\"abcd\"\n",
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
