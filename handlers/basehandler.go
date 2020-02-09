package hdl

import (
	"path"
	"strings"
)

type BaseHandler struct {
	root string
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
