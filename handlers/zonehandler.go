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
		&HandlerMeta{
			Route:   zh.ToRoute("add"),
			Handle:  zh.AddZone,
			Methods: []string{POST},
		},
		&HandlerMeta{
			Route:   zh.ToRoute("remove"),
			Handle:  zh.RemoveZone,
			Methods: []string{POST}, // XXX: Probably should be DELETE instead
		},
		&HandlerMeta{
			Route:   zh.ToRoute("update"),
			Handle:  zh.UpdateZone,
			Methods: []string{POST},
		},
	}
}

// AddZone creates a zone in the cluster
func (zh *ZoneHandler) AddZone(ctx *gin.Context) {
}

// UpdateZone updates a zone data
func (zh *ZoneHandler) UpdateZone(ctx *gin.Context) {
}

// RemoveZone removes a zone from the cluster, but only if it is
// empty (no nodes assigned to it).
func (zh *ZoneHandler) RemoveZone(ctx *gin.Context) {
}

// Handle implements the entry point of the handler
func (nh *ZoneHandler) ListZones(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"zones": gin.H{
			"name": "foo",
		},
	})
}
