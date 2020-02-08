package clbd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemsHandler struct {
	dbx *Dbx
}

func NewSystemsHandler() *SystemsHandler {
	return new(SystemsHandler)
}

// SetDbx sets the Dbx instance pointer
func (sh *SystemsHandler) SetDbx(dbx *Dbx) {
	sh.dbx = dbx
}

// Handlers returns a map of supported handlers and their configuration
func (nh *SystemsHandler) Handlers() []*HandlerMeta {
	return []*HandlerMeta{
		&HandlerMeta{
			Route:   "systems",
			Handle:  nh.OnSystems,
			Methods: []string{POST, GET},
		},
	}
}

// Handle implements the entry point of the handler
func (nh *SystemsHandler) OnSystems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
