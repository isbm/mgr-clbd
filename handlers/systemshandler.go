package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
	"net/http"
)

type SystemsHandler struct {
	BaseHandler
	db *dbx.Dbx
}

func NewSystemsHandler(root string) *SystemsHandler {
	sh := new(SystemsHandler)
	sh.PrepareRoot(root)
	return sh
}

// Backend returns an underlying backend interface
func (ph *SystemsHandler) Backend() backend.Backend {
	return nil
}

// SetDbx sets the Dbx instance pointer
func (sh *SystemsHandler) SetDbx(db *dbx.Dbx) {
	sh.db = db
}

// Handlers returns a map of supported handlers and their configuration
func (nh *SystemsHandler) Handlers() []*HandlerMeta {
	return []*HandlerMeta{
		&HandlerMeta{
			Route:   nh.ToRoute("list"),
			Handle:  nh.ListSystems,
			Methods: []string{POST, GET},
		},
	}
}

// Handle implements the entry point of the handler
func (nh *SystemsHandler) ListSystems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
