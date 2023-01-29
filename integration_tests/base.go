package integration_tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	test_env_path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(test_env_path)
	apppath, _ := filepath.Abs(test_env_path)
	fmt.Println(apppath)
	web.TestBeegoInit(apppath)
	web.BConfig.RunMode = "test"
}

func MakeRequest(method string, url string, headers map[string]string, body []byte, lastResponse *httptest.ResponseRecorder) *httptest.ResponseRecorder {

	bodyReader := bytes.NewReader(body)
	r, _ := http.NewRequest(method, url, bodyReader)
	for header, value := range headers {
		r.Header.Set(header, value)
	}

	if lastResponse != nil {
		respHeader := lastResponse.Header()
		var cookies []string
		for _, raw_cookie := range respHeader.Values("set-cookie") {
			cookies = append(cookies, strings.Split(raw_cookie, ";")[0])
		}
		r.Header.Set("Cookie", strings.Join(append(cookies, r.Header.Values("Cookie")...), "; "))
	}

	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)
	return w
}
