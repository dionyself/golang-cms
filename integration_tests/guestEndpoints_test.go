package integration_tests

import (
	"testing"

	_ "github.com/beego/beego/v2/server/web/session/redis"
	_ "github.com/dionyself/golang-cms/routers"

	log "github.com/beego/beego/v2/core/logs"
	. "github.com/smartystreets/goconvey/convey"
)

// TestMain tests the main page
func TestMain(t *testing.T) {
	jsonBody := []byte(`{"client_message": "hello, server!"}`)
	w := MakeRequest("GET", "/", map[string]string{"Content-Type": "application/json"}, jsonBody, nil)

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
