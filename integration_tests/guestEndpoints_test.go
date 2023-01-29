package integration_tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/beego/beego/v2/server/web/session/redis"
	_ "github.com/dionyself/golang-cms/routers"

	log "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
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

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	jsonBody := []byte(`{"client_message": "hello, server!"}`)
	bodyReader := bytes.NewReader(jsonBody)
	r, _ := http.NewRequest("GET", "/", bodyReader)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(w, r)

	log.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
