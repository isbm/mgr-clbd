// Handler interface

package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
)

const (
	ANY  = "any"
	POST = "post"
	GET  = "get"
)

type HandlerMeta struct {
	Route   string
	Handle  gin.HandlerFunc
	Methods []string
}

type Handler interface {
	// Return underlying backend
	Backend() backend.Backend

	// Route map
	Handlers() []*HandlerMeta

	// Set Dbx object to cross-access the database calls
	SetDbx(db *dbx.Dbx)
}
