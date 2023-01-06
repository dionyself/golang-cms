// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"bytes"
	context2 "context"
	"errors"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"

	"github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/context/param"
	"github.com/beego/beego/v2/server/web/session"
)

var (
	// ErrAbort custom error when user stop request handler manually.
	ErrAbort = errors.New("user stop run")
	// GlobalControllerRouter store comments with controller. pkgpath+controller:comments
	GlobalControllerRouter = make(map[string][]ControllerComments)
	copyBufferPool         sync.Pool
)

const (
	bytePerKb    = 1024
	copyBufferKb = 32
	filePerm     = 0o666
)

func init() {
	copyBufferPool.New = func() interface{} {
		return make([]byte, bytePerKb*copyBufferKb)
	}
}

// ControllerFilter store the filter for controller
type ControllerFilter struct {
	Pattern        string
	Pos            int
	Filter         web.FilterFunc
	ReturnOnOutput bool
	ResetParams    bool
}

// ControllerFilterComments store the comment for controller level filter
type ControllerFilterComments struct {
	Pattern        string
	Pos            int
	Filter         string // NOQA
	ReturnOnOutput bool
	ResetParams    bool
}

// ControllerImportComments store the import comment for controller needed
type ControllerImportComments struct {
	ImportPath  string
	ImportAlias string
}

// ControllerComments store the comment for the controller method
type ControllerComments struct {
	Method           string
	Router           string
	Filters          []*ControllerFilter
	ImportComments   []*ControllerImportComments
	FilterComments   []*ControllerFilterComments
	AllowHTTPMethods []string
	Params           []map[string]string
	MethodParams     []*param.MethodParam
}

// ControllerCommentsSlice implements the sort interface
type ControllerCommentsSlice []ControllerComments

func (p ControllerCommentsSlice) Len() int { return len(p) }

func (p ControllerCommentsSlice) Less(i, j int) bool { return p[i].Router < p[j].Router }

func (p ControllerCommentsSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Controller defines some basic http request handler operations, such as
// http context, template and view, session and xsrf.
type CustomController struct {
	// context data
	Ctx  *context.Context
	Data map[interface{}]interface{}

	// route controller info
	controllerName string
	actionName     string
	methodMapping  map[string]func() //method:routertree
	AppController  interface{}

	// template data
	TplName        string
	ViewPath       string
	Layout         string
	LayoutSections map[string]string // the key is the section name and the value is the template name
	TplPrefix      string
	TplExt         string
	EnableRender   bool

	// xsrf data
	EnableXSRF bool
	_xsrfToken string
	XSRFExpire int

	// session
	CruSession session.Store
}

// ControllerInterface is an interface to uniform all controller handler.
type ControllerInterface interface {
	Init(ct *context.Context, controllerName, actionName string, app interface{})
	Prepare()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Trace()
	Finish()
	Render() error
	XSRFToken() string
	CheckXSRFCookie() bool
	HandlerFunc(fn string) bool
	URLMapping()
}

// Init generates default values of controller operations.
func (c *CustomController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Layout = ""
	c.TplName = ""
	c.controllerName = controllerName
	c.actionName = actionName
	c.Ctx = ctx
	c.TplExt = "tpl"
	c.AppController = app
	c.EnableRender = true
	c.EnableXSRF = true
	c.Data = ctx.Input.Data()
	c.methodMapping = make(map[string]func())
}

// Prepare runs after Init before request function execution.
func (c *CustomController) Prepare() {}

// Finish runs after request function execution.
func (c *CustomController) Finish() {}

// Get adds a request function to handle GET request.
func (c *CustomController) Get() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Post adds a request function to handle POST request.
func (c *CustomController) Post() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Delete adds a request function to handle DELETE request.
func (c *CustomController) Delete() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Put adds a request function to handle PUT request.
func (c *CustomController) Put() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Head adds a request function to handle HEAD request.
func (c *CustomController) Head() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Patch adds a request function to handle PATCH request.
func (c *CustomController) Patch() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Options adds a request function to handle OPTIONS request.
func (c *CustomController) Options() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Trace adds a request function to handle Trace request.
// this method SHOULD NOT be overridden.
// https://tools.ietf.org/html/rfc7231#section-4.3.8
// The TRACE method requests a remote, application-level loop-back of
// the request message.  The final recipient of the request SHOULD
// reflect the message received, excluding some fields described below,
// back to the client as the message body of a 200 (OK) response with a
// Content-Type of "message/http" (Section 8.3.1 of [RFC7230]).
func (c *CustomController) Trace() {
	ts := func(h http.Header) (hs string) {
		for k, v := range h {
			hs += fmt.Sprintf("\r\n%s: %s", k, v)
		}
		return
	}
	hs := fmt.Sprintf("\r\nTRACE %s %s%s\r\n", c.Ctx.Request.RequestURI, c.Ctx.Request.Proto, ts(c.Ctx.Request.Header))
	c.Ctx.Output.Header("Content-Type", "message/http")
	c.Ctx.Output.Header("Content-Length", fmt.Sprint(len(hs)))
	c.Ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Ctx.WriteString(hs)
}

// HandlerFunc call function with the name
func (c *CustomController) HandlerFunc(fnname string) bool {
	if v, ok := c.methodMapping[fnname]; ok {
		v()
		return true
	}
	return false
}

// URLMapping register the internal Controller router.
func (c *CustomController) URLMapping() {}

// Bind if the content type is form, we read data from form
// otherwise, read data from request body
func (c *CustomController) Bind(obj interface{}) error {
	return c.Ctx.Bind(obj)
}

// BindYAML only read data from http request body
func (c *CustomController) BindYAML(obj interface{}) error {
	return c.Ctx.BindYAML(obj)
}

// BindForm read data from form
func (c *CustomController) BindForm(obj interface{}) error {
	return c.Ctx.BindForm(obj)
}

// BindJSON only read data from http request body
func (c *CustomController) BindJSON(obj interface{}) error {
	return c.Ctx.BindJSON(obj)
}

// BindProtobuf only read data from http request body
func (c *CustomController) BindProtobuf(obj proto.Message) error {
	return c.Ctx.BindProtobuf(obj)
}

// BindXML only read data from http request body
func (c *CustomController) BindXML(obj interface{}) error {
	return c.Ctx.BindXML(obj)
}

// Mapping the method to function
func (c *CustomController) Mapping(method string, fn func()) {
	c.methodMapping[method] = fn
}

// Render sends the response with rendered template bytes as text/html type.
func (c *CustomController) Render() error {
	if !c.EnableRender {
		return nil
	}
	rb, err := c.RenderBytes()
	if err != nil {
		return err
	}

	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	return c.Ctx.Output.Body(rb)
}

// RenderString returns the rendered template string. Do not send out response.
func (c *CustomController) RenderString() (string, error) {
	b, e := c.RenderBytes()
	if e != nil {
		return "", e
	}
	return string(b), e
}

// RenderBytes returns the bytes of rendered template string. Do not send out response.
func (c *CustomController) RenderBytes() ([]byte, error) {
	buf, err := c.renderTemplate()
	// if the controller has set layout, then first get the tplName's content set the content to the layout
	if err == nil && c.Layout != "" {
		c.Data["LayoutContent"] = template.HTML(buf.String())

		if c.LayoutSections != nil {
			for sectionName, sectionTpl := range c.LayoutSections {
				if sectionTpl == "" {
					c.Data[sectionName] = ""
					continue
				}
				buf.Reset()
				c.Data["SectionData"] = c.Data[sectionName+"_Data"]
				err = web.ExecuteViewPathTemplate(&buf, sectionTpl, c.viewPath(), c.Data)
				if err != nil {
					return nil, err
				}
				c.Data[sectionName] = template.HTML(buf.String())
			}
		}

		buf.Reset()
		err = web.ExecuteViewPathTemplate(&buf, c.Layout, c.viewPath(), c.Data)
	}
	return buf.Bytes(), err
}

func (c *CustomController) renderTemplate() (bytes.Buffer, error) {
	var buf bytes.Buffer
	if c.TplName == "" {
		c.TplName = strings.ToLower(c.controllerName) + "/" + strings.ToLower(c.actionName) + "." + c.TplExt
	}
	if c.TplPrefix != "" {
		c.TplName = c.TplPrefix + c.TplName
	}
	if web.BConfig.RunMode == web.DEV {
		buildFiles := []string{c.TplName}
		if c.Layout != "" {
			buildFiles = append(buildFiles, c.Layout)
			if c.LayoutSections != nil {
				for _, sectionTpl := range c.LayoutSections {
					if sectionTpl == "" {
						continue
					}
					buildFiles = append(buildFiles, sectionTpl)
				}
			}
		}
		web.BuildTemplate(c.viewPath(), buildFiles...)
	}
	return buf, web.ExecuteViewPathTemplate(&buf, c.TplName, c.viewPath(), c.Data)
}

func (c *CustomController) viewPath() string {
	if c.ViewPath == "" {
		return web.BConfig.WebConfig.ViewsPath
	}
	return c.ViewPath
}

// Redirect sends the redirection response to url with status code.
func (c *CustomController) Redirect(url string, code int) {
	web.LogAccess(c.Ctx, nil, code)
	c.Ctx.Redirect(code, url)
}

// SetData set the data depending on the accepted
func (c *CustomController) SetData(data interface{}) {
	accept := c.Ctx.Input.Header("Accept")
	switch accept {
	case context.ApplicationYAML:
		c.Data["yaml"] = data
	case context.ApplicationXML, context.TextXML:
		c.Data["xml"] = data
	default:
		c.Data["json"] = data
	}
}

// Abort stops controller handler and show the error data if code is defined in ErrorMap or code string.
func (c *CustomController) Abort(code string) {
	status, err := strconv.Atoi(code)
	if err != nil {
		status = 200
	}
	c.CustomAbort(status, code)
}

// CustomAbort stops controller handler and show the error data, it's similar Aborts, but support status code and body.
func (c *CustomController) CustomAbort(status int, body string) {
	c.Ctx.Output.Status = status
	// first panic from ErrorMaps, it is user defined error functions.
	if _, ok := web.ErrorMaps[body]; ok {
		panic(body)
	}
	// last panic user string
	c.Ctx.ResponseWriter.WriteHeader(status)
	c.Ctx.ResponseWriter.Write([]byte(body))
	panic(ErrAbort)
}

// StopRun makes panic of USERSTOPRUN error and go to recover function if defined.
func (c *CustomController) StopRun() {
	panic(ErrAbort)
}

// URLFor does another controller handler in this request function.
// it goes to this controller method if endpoint is not clear.
func (c *CustomController) URLFor(endpoint string, values ...interface{}) string {
	if len(endpoint) == 0 {
		return ""
	}
	if endpoint[0] == '.' {
		return web.URLFor(reflect.Indirect(reflect.ValueOf(c.AppController)).Type().Name()+endpoint, values...)
	}
	return web.URLFor(endpoint, values...)
}

func (c *CustomController) JSONResp(data interface{}) error {
	return c.Ctx.JSONResp(data)
}

func (c *CustomController) XMLResp(data interface{}) error {
	return c.Ctx.XMLResp(data)
}

func (c *CustomController) YamlResp(data interface{}) error {
	return c.Ctx.YamlResp(data)
}

// Resp sends response based on the Accept Header
// By default response will be in JSON
// it's different from ServeXXX methods
// because we don't store the data to Data field
func (c *CustomController) Resp(data interface{}) error {
	return c.Ctx.Resp(data)
}

// ServeJSON sends a json response with encoding charset.
func (c *CustomController) ServeJSON(encoding ...bool) error {
	var (
		hasIndent   = web.BConfig.RunMode != web.PROD
		hasEncoding = len(encoding) > 0 && encoding[0]
	)

	return c.Ctx.Output.JSON(c.Data["json"], hasIndent, hasEncoding)
}

// ServeJSONP sends a jsonp response.
func (c *CustomController) ServeJSONP() error {
	hasIndent := web.BConfig.RunMode != web.PROD
	return c.Ctx.Output.JSONP(c.Data["jsonp"], hasIndent)
}

// ServeXML sends xml response.
func (c *CustomController) ServeXML() error {
	hasIndent := web.BConfig.RunMode != web.PROD
	return c.Ctx.Output.XML(c.Data["xml"], hasIndent)
}

// ServeYAML sends yaml response.
func (c *CustomController) ServeYAML() error {
	return c.Ctx.Output.YAML(c.Data["yaml"])
}

// ServeFormatted serve YAML, XML OR JSON, depending on the value of the Accept header
func (c *CustomController) ServeFormatted(encoding ...bool) error {
	hasIndent := web.BConfig.RunMode != web.PROD
	hasEncoding := len(encoding) > 0 && encoding[0]
	return c.Ctx.Output.ServeFormatted(c.Data, hasIndent, hasEncoding)
}

// Input returns the input data map from POST or PUT request body and query string.
func (c *CustomController) Input() (url.Values, error) {
	if c.Ctx.Request.Form == nil {
		err := c.Ctx.Request.ParseForm()
		if err != nil {
			return nil, err
		}
	}
	return c.Ctx.Request.Form, nil
}

// ParseForm maps input data map to obj struct.
func (c *CustomController) ParseForm(obj interface{}) error {
	return c.Ctx.BindForm(obj)
}

// GetString returns the input value by key string or the default value while it's present and input is blank
func (c *CustomController) GetString(key string, def ...string) string {
	if v := c.Ctx.Input.Query(key); v != "" {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetStrings returns the input string slice by key string or the default value while it's present and input is blank
// it's designed for multi-value input field such as checkbox(input[type=checkbox]), multi-selection.
func (c *CustomController) GetStrings(key string, def ...[]string) []string {
	var defv []string
	if len(def) > 0 {
		defv = def[0]
	}

	if f, err := c.Input(); f == nil || err != nil {
		return defv
	} else if vs := f[key]; len(vs) > 0 {
		return vs
	}

	return defv
}

// GetInt returns input as an int or the default value while it's present and input is blank
func (c *CustomController) GetInt(key string, def ...int) (int, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.Atoi(strv)
}

// GetInt8 return input as an int8 or the default value while it's present and input is blank
func (c *CustomController) GetInt8(key string, def ...int8) (int8, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	i64, err := strconv.ParseInt(strv, 10, 8)
	return int8(i64), err
}

// GetUint8 return input as an uint8 or the default value while it's present and input is blank
func (c *CustomController) GetUint8(key string, def ...uint8) (uint8, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	u64, err := strconv.ParseUint(strv, 10, 8)
	return uint8(u64), err
}

// GetInt16 returns input as an int16 or the default value while it's present and input is blank
func (c *CustomController) GetInt16(key string, def ...int16) (int16, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	i64, err := strconv.ParseInt(strv, 10, 16)
	return int16(i64), err
}

// GetUint16 returns input as an uint16 or the default value while it's present and input is blank
func (c *CustomController) GetUint16(key string, def ...uint16) (uint16, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	u64, err := strconv.ParseUint(strv, 10, 16)
	return uint16(u64), err
}

// GetInt32 returns input as an int32 or the default value while it's present and input is blank
func (c *CustomController) GetInt32(key string, def ...int32) (int32, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	i64, err := strconv.ParseInt(strv, 10, 32)
	return int32(i64), err
}

// GetUint32 returns input as an uint32 or the default value while it's present and input is blank
func (c *CustomController) GetUint32(key string, def ...uint32) (uint32, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	u64, err := strconv.ParseUint(strv, 10, 32)
	return uint32(u64), err
}

// GetInt64 returns input value as int64 or the default value while it's present and input is blank.
func (c *CustomController) GetInt64(key string, def ...int64) (int64, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseInt(strv, 10, 64)
}

// GetUint64 returns input value as uint64 or the default value while it's present and input is blank.
func (c *CustomController) GetUint64(key string, def ...uint64) (uint64, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseUint(strv, 10, 64)
}

// GetBool returns input value as bool or the default value while it's present and input is blank.
func (c *CustomController) GetBool(key string, def ...bool) (bool, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseBool(strv)
}

// GetFloat returns input value as float64 or the default value while it's present and input is blank.
func (c *CustomController) GetFloat(key string, def ...float64) (float64, error) {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseFloat(strv, 64)
}

// GetFile returns the file data in file upload field named as key.
// it returns the first one of multi-uploaded files.
func (c *CustomController) GetFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return c.Ctx.Request.FormFile(key)
}

// GetFiles return multi-upload files
// files, err:=c.GetFiles("myfiles")
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusNoContent)
//		return
//	}
//
//	for i, _ := range files {
//		//for each fileheader, get a handle to the actual file
//		file, err := files[i].Open()
//		defer file.Close()
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		//create destination file making sure the path is writeable.
//		dst, err := os.Create("upload/" + files[i].Filename)
//		defer dst.Close()
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		//copy the uploaded file to the destination file
//		if _, err := io.Copy(dst, file); err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}
func (c *CustomController) GetFiles(key string) ([]*multipart.FileHeader, error) {
	if files, ok := c.Ctx.Request.MultipartForm.File[key]; ok {
		return files, nil
	}
	return nil, http.ErrMissingFile
}

// SaveToFile saves uploaded file to new path.
// it only operates the first one of mutil-upload form file field.
func (c *CustomController) SaveToFile(fromFile, toFile string) error {
	buf := copyBufferPool.Get().([]byte)
	defer copyBufferPool.Put(buf)
	return c.SaveToFileWithBuffer(fromFile, toFile, buf)
}

type onlyWriter struct {
	io.Writer
}

func (c *CustomController) SaveToFileWithBuffer(fromFile string, toFile string, buf []byte) error {
	src, _, err := c.Ctx.Request.FormFile(fromFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(toFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filePerm)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.CopyBuffer(onlyWriter{dst}, src, buf)
	return err
}

// StartSession starts session and load old session data info this controller.
func (c *CustomController) StartSession() session.Store {
	if c.CruSession == nil {
		c.CruSession = c.Ctx.Input.CruSession
	}
	return c.CruSession
}

// SetSession puts value into session.
func (c *CustomController) SetSession(name interface{}, value interface{}) error {
	if c.CruSession == nil {
		c.StartSession()
	}
	return c.CruSession.Set(context2.Background(), name, value)
}

// GetSession gets value from session.
func (c *CustomController) GetSession(name interface{}) interface{} {
	if c.CruSession == nil {
		c.StartSession()
	}
	return c.CruSession.Get(context2.Background(), name)
}

// DelSession removes value from session.
func (c *CustomController) DelSession(name interface{}) error {
	if c.CruSession == nil {
		c.StartSession()
	}
	return c.CruSession.Delete(context2.Background(), name)
}

// SessionRegenerateID regenerates session id for this session.
// the session data have no changes.
func (c *CustomController) SessionRegenerateID() error {
	if c.CruSession != nil {
		c.CruSession.SessionRelease(context2.Background(), c.Ctx.ResponseWriter)
	}
	var err error
	c.CruSession, err = web.GlobalSessions.SessionRegenerateID(c.Ctx.ResponseWriter, c.Ctx.Request)
	c.Ctx.Input.CruSession = c.CruSession
	return err
}

// DestroySession cleans session data and session cookie.
func (c *CustomController) DestroySession() error {
	err := c.Ctx.Input.CruSession.Flush(nil)
	if err != nil {
		return err
	}
	c.Ctx.Input.CruSession = nil
	web.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
	return nil
}

// IsAjax returns this request is ajax or not.
func (c *CustomController) IsAjax() bool {
	return c.Ctx.Input.IsAjax()
}

// GetSecureCookie returns decoded cookie value from encoded browser cookie values.
func (c *CustomController) GetSecureCookie(Secret, key string) (string, bool) {
	return c.Ctx.GetSecureCookie(Secret, key)
}

// SetSecureCookie puts value into cookie after encoded the value.
func (c *CustomController) SetSecureCookie(Secret, name, value string, others ...interface{}) {
	c.Ctx.SetSecureCookie(Secret, name, value, others...)
}

// XSRFToken creates a CSRF token string and returns.
func (c *CustomController) XSRFToken() string {
	if c._xsrfToken == "" {
		expire := int64(web.BConfig.WebConfig.XSRFExpire)
		if c.XSRFExpire > 0 {
			expire = int64(c.XSRFExpire)
		}
		c._xsrfToken = c.Ctx.XSRFToken(web.BConfig.WebConfig.XSRFKey, expire)
	}
	return c._xsrfToken
}

// CheckXSRFCookie checks xsrf token in this request is valid or not.
// the token can provided in request header "X-Xsrftoken" and "X-CsrfToken"
// or in form field value named as "_xsrf".
func (c *CustomController) CheckXSRFCookie() bool {
	if !c.EnableXSRF {
		return true
	}
	return c.Ctx.CheckXSRFCookie()
}

// XSRFFormHTML writes an input field contains xsrf token value.
func (c *CustomController) XSRFFormHTML() string {
	return `<input type="hidden" name="_xsrf" value="` +
		c.XSRFToken() + `" />`
}

// GetControllerAndAction gets the executing controller name and action name.
func (c *CustomController) GetControllerAndAction() (string, string) {
	return c.controllerName, c.actionName
}
