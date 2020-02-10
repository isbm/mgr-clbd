package hdl

import (
	"github.com/isbm/mgr-clbd/utils"
	"github.com/sirupsen/logrus"
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
