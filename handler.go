package clbd

import (
	"github.com/gin-gonic/gin"
)

type HandlerMeta struct {
	Route   string
	Handle  gin.HandlerFunc
	Methods []string
}

type Handler interface {
	// Route map
	Handlers() []*HandlerMeta

	// Set Dbx object to cross-access the database calls
	SetDbx(dbx *Dbx)
}
