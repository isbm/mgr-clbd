package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
	"net/http"
)

type ZoneHandler struct {
	BaseHandler
	db *dbx.Dbx
}

func NewZoneHandler(root string) *ZoneHandler {
	zh := new(ZoneHandler)
	zh.PrepareRoot(root)
	return zh
}

// Backend returns an underlying backend interface
func (ph *ZoneHandler) Backend() backend.Backend {
	return nil
}

// SetDbx sets the Dbx instance pointer
func (sh *ZoneHandler) SetDbx(db *dbx.Dbx) {
	sh.db = db
}

// Handlers returns a map of supported handlers and their configuration
func (zh *ZoneHandler) Handlers() []*HandlerMeta {
	return []*HandlerMeta{
		&HandlerMeta{
			Route:   zh.ToRoute("list"),
			Handle:  zh.ListZones,
			Methods: []string{POST, GET},
		},
	}
}

// Handle implements the entry point of the handler
func (nh *ZoneHandler) ListZones(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"zones": gin.H{
			"name": "foo",
		},
	})
}
