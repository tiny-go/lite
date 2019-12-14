package lite

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/tiny-go/codec/driver"
	_ "github.com/tiny-go/codec/driver/json"
)

func Test_All(t *testing.T) {
	t.Run("Given an HTTP handler with registered module", func(t *testing.T) {
		driver.Default("application/json")
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
					r, _ := http.NewRequest(http.MethodGet, ts.URL+"/test/pass?foo=bar", nil)
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":[\"bar\"]}\n",
			},
			{
				title: "plural GET with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, ts.URL+"/test/fail", nil)
					return r
				}(),
				code: http.StatusBadRequest,
				body: "plural GET error\n",
			},
			{
				title: "single GET with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, ts.URL+"/test/pass/abcd", nil)
					return r
				}(),
				code: http.StatusOK,
				body: "\"abcd\"\n",
			},
			{
				title: "single GET with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, ts.URL+"/test/fail/abcd", nil)
					return r
				}(),
				code: http.StatusBadRequest,
				body: "single GET error\n",
			},
			{
				title: "plural POST with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPost, ts.URL+"/test/pass", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\"}\n",
			},
			{
				title: "plural POST with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPost, ts.URL+"/test/fail", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusBadRequest,
				body: "plural POST error\n",
			},
			{
				title: "single POST with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPost, ts.URL+"/test/pass/abcd", strings.NewReader("{\"foo\":\"bar\",\"pk\":\"abcd\"}\n"))
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\",\"pk\":\"abcd\"}\n",
			},
			{
				title: "single POST with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPost, ts.URL+"/test/fail/abcd", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusBadRequest,
				body: "single POST error\n",
			},
			{
				title: "plural PATCH with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPatch, ts.URL+"/test/pass", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\"}\n",
			},
			{
				title: "plural PATCH with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPatch, ts.URL+"/test/fail", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusBadRequest,
				body: "plural PATCH error\n",
			},
			{
				title: "single PATCH with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPatch, ts.URL+"/test/pass/abcd", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\",\"pk\":\"abcd\"}\n",
			},
			{
				title: "single PATCH with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPatch, ts.URL+"/test/fail/abcd", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusBadRequest,
				body: "single PATCH error\n",
			},
			{
				title: "plural PUT with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPut, ts.URL+"/test/pass", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\"}\n",
			},
			{
				title: "plural PUT with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPut, ts.URL+"/test/pass", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\"}\n",
			},
			{
				title: "plural PUT with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPut, ts.URL+"/test/fail", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusBadRequest,
				body: "plural PUT error\n",
			},
			{
				title: "single PUT with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPut, ts.URL+"/test/pass/abcd", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":\"bar\",\"pk\":\"abcd\"}\n",
			},
			{
				title: "single PUT with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodPut, ts.URL+"/test/fail/abcd", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusBadRequest,
				body: "single PUT error\n",
			},
			{
				title: "plural DELETE with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodDelete, ts.URL+"/test/pass?foo=bar", nil)
					return r
				}(),
				code: http.StatusOK,
				body: "{\"foo\":[\"bar\"]}\n",
			},
			{
				title: "plural DELETE with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodDelete, ts.URL+"/test/fail", strings.NewReader("{\"foo\":\"bar\"}"))
					return r
				}(),
				code: http.StatusBadRequest,
				body: "plural DELETE error\n",
			},
			{
				title: "single DELETE with success",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodDelete, ts.URL+"/test/pass/abcd", nil)
					return r
				}(),
				code: http.StatusOK,
				body: "\"abcd\"\n",
			},
			{
				title: "single DELETE with failure",
				request: func() *http.Request {
					r, _ := http.NewRequest(http.MethodDelete, ts.URL+"/test/fail/abcd", nil)
					return r
				}(),
				code: http.StatusBadRequest,
				body: "single DELETE error\n",
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
