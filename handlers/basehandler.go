package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type BaseHandler struct {
	root        string
	_validators *utils.Validators
	_logger     *logrus.Logger
}

// PrepareRoot is sanitising root string, turning into a root URI
func (bh *BaseHandler) PrepareRoot(root string) string {
	bh.root = "/" + strings.Trim(root, "/")
	return bh.root
}

func (bh BaseHandler) ToRoute(route string) string {
	return path.Join(bh.root, route)
}

func (bh *BaseHandler) URI() string {
	return bh.root
}

// GetLogger returns initalised or an instance of working logger
func (bh *BaseHandler) GetLogger() *logrus.Logger {
	if bh._logger == nil {
		bh._logger = utils.GetTextLogger(logrus.DebugLevel, nil)
	}
	return bh._logger
}

// GetValidators returns initialised or an instance of working validators
func (bh *BaseHandler) GetValidators() *utils.Validators {
	if bh._validators == nil {
		bh._validators = utils.NewValidators()
	}
	return bh._validators
}

// InitBody parses query in body (usually ends up on DELETE methods)
func (bh *BaseHandler) InitBody(ctx *gin.Context, names ...string) *ReturnType {
	ret := NewReturnType(ctx)
	data, err := ctx.GetRawData()
	if err != nil {
		ret.SetError(err).SetErrorCode(http.StatusBadRequest).SendJSON()
		return nil
	}

	values, err := url.ParseQuery(string(data))
	if err != nil {
		ret.SetError(err).SetErrorCode(http.StatusBadRequest).SendJSON()
		return nil
	}

	return ret.SetValues(&values)
}

// InitForm initialises the form in the Request object instance
// and returns standard return type.
func (bh *BaseHandler) InitForm(ctx *gin.Context, names ...string) *ReturnType {
	ret := NewReturnType(ctx)
	err := ctx.Request.ParseForm()
	if err != nil {
		ret.SetError(err).SetErrorCode(http.StatusBadRequest).SendJSON()
		return nil
	}

	errcode, msg := bh.GetValidators().VerifyRequired(ctx.Request, names...)
	if errcode != http.StatusOK {
		ret.SetErrorMessage(msg).SetErrorCode(errcode).SendJSON()
		return nil
	}
	return ret
}
