package hdl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

type ReturnType struct {
	errmsg  error
	errcode int
	msg     string
	payload interface{}
	ctx     *gin.Context
	values  *url.Values
}

func NewReturnType(ctx *gin.Context) *ReturnType {
	rt := new(ReturnType)
	rt.errcode = http.StatusOK
	rt.ctx = ctx

	return rt
}

func (rt *ReturnType) SetValues(values *url.Values) *ReturnType {
	rt.values = values
	return rt
}

func (rt *ReturnType) GetValues() *url.Values {
	return rt.values
}

func (rt *ReturnType) SetErrorMessage(msg string) *ReturnType {
	rt.errmsg = errors.New(msg)
	return rt
}

func (rt *ReturnType) SetError(err error) *ReturnType {
	rt.errmsg = err
	return rt
}

func (rt *ReturnType) SetErrorCode(code int) *ReturnType {
	rt.errcode = code
	return rt
}

func (rt *ReturnType) SetMessage(msg string) *ReturnType {
	rt.msg = msg
	return rt
}

func (rt *ReturnType) SetPayload(payload interface{}) *ReturnType {
	rt.payload = payload
	return rt
}

func (rt *ReturnType) Serialise() map[string]interface{} {
	out := make(map[string]interface{})
	out["errcode"] = rt.errcode
	if rt.errmsg != nil || rt.errcode != http.StatusOK {
		out["error"] = strings.TrimSpace(rt.errmsg.Error() + " " + rt.msg)
	} else {
		out["msg"] = rt.msg
		out["data"] = rt.payload
	}

	return out
}

func (rt *ReturnType) SendJSON() {
	rt.ctx.JSON(rt.errcode, rt.Serialise())
}
