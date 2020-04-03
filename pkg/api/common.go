package api

import (
	"encoding/json"
	"net/http"

	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/log"
	"github.com/hooone/evening/pkg/setting"
	macaron "gopkg.in/macaron.v1"
)

var (
	//NotFound 404
	NotFound = func() Response {
		return Error(404, "Not found", nil)
	}

	//ServerError 500
	ServerError = func(err error) Response {
		return Error(500, "Server error", err)
	}
)

//CommonResult result for all Json request
type CommonResult struct {
	Success bool
	Message string
	Data    interface{}
}

//Response interface for WriteTo function
type Response interface {
	WriteTo(ctx *dtos.ReqContext)
}

//Wrap convert action to macaron.Handler
func Wrap(action interface{}) macaron.Handler {
	return func(c *dtos.ReqContext) {
		var res Response
		val, err := c.Invoke(action)

		if err == nil && val != nil && len(val) > 0 {
			res = val[0].Interface().(Response)
		} else {
			res = ServerError(err)
		}

		res.WriteTo(c)
	}
}

//NormalResponse normal result
type NormalResponse struct {
	status     int
	body       []byte
	header     http.Header
	errMessage string
	err        error
}

//Header set header
func (r *NormalResponse) Header(key, value string) *NormalResponse {
	r.header.Set(key, value)
	return r
}

//WriteTo use for write return code to http response
func (r *NormalResponse) WriteTo(ctx *dtos.ReqContext) {
	if r.err != nil {
		logger := log.New("context")
		logger.Error(r.errMessage, "error", r.err)
	}

	header := ctx.Resp.Header()
	for k, v := range r.header {
		header[k] = v
	}
	ctx.Resp.WriteHeader(r.status)
	if _, err := ctx.Resp.Write(r.body); err != nil {
		logger := log.New("context")
		logger.Error("Error writing to response", "err", err)
	}
}

// Respond create a response
func Respond(status int, body interface{}) *NormalResponse {
	var b []byte
	var err error
	switch t := body.(type) {
	case []byte:
		b = t
	case string:
		b = []byte(t)
	default:
		if b, err = json.Marshal(body); err != nil {
			return Error(500, "body json marshal", err)
		}
	}
	return &NormalResponse{
		body:   b,
		status: status,
		header: make(http.Header),
	}
}

// JSON create a JSON response
func JSON(status int, body interface{}) *NormalResponse {
	return Respond(status, body).Header("Content-Type", "application/json")
}

// Error create a erroneous response
func Error(status int, message string, err error) *NormalResponse {
	data := make(map[string]interface{})

	switch status {
	case 404:
		data["message"] = "Not Found"
	case 500:
		data["message"] = "Internal Server Error"
	}

	if message != "" {
		data["message"] = message
	}

	if err != nil {
		if setting.Env != setting.PROD {
			data["error"] = err.Error()
		}
	}

	resp := JSON(status, data)

	if err != nil {
		resp.errMessage = message
		resp.err = err
	}

	return resp
}
