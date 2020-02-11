package hdl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
	"net/http"
)

type ZoneHandler struct {
	BaseHandler
	db  *dbx.Dbx
	bnd *backend.Zones
}

func NewZoneHandler(root string) *ZoneHandler {
	zh := new(ZoneHandler)
	zh.bnd = backend.NewZonesBackend()
	zh.PrepareRoot(root)
	return zh
}

// Backend returns an underlying backend interface
func (zh *ZoneHandler) Backend() backend.Backend {
	return zh.bnd
}

// SetDbx sets the Dbx instance pointer
func (sh *ZoneHandler) SetDbx(db *dbx.Dbx) {
	sh.db = db
	sh.bnd.SetDbx(sh.db)
}

// Handlers returns a map of supported handlers and their configuration
func (zh *ZoneHandler) Handlers() []*HandlerMeta {
	return []*HandlerMeta{
		&HandlerMeta{
			Route:   zh.ToRoute("list"),
			Handle:  zh.ListZones,
			Methods: []string{GET},
		},
		&HandlerMeta{
			Route:   zh.ToRoute("add"),
			Handle:  zh.AddZone,
			Methods: []string{POST},
		},
		&HandlerMeta{
			Route:   zh.ToRoute("remove"),
			Handle:  zh.RemoveZone,
			Methods: []string{DELETE}, // XXX: Probably should be DELETE instead
		},
		&HandlerMeta{
			Route:   zh.ToRoute("update"),
			Handle:  zh.UpdateZone,
			Methods: []string{POST},
		},
		&HandlerMeta{
			Route:   zh.ToRoute("stats"),
			Handle:  zh.ZoneStats,
			Methods: []string{GET},
		},
	}
}

// ZoneStats godoc
// @Summary Return Zone stats.
// @Description ZoneStats returns data about zone.
// @ID zone-stats
// @Accept json
// @Produce json
// @Param name query string true "Name of the Zone"
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/stats [get]
func (zh *ZoneHandler) ZoneStats(ctx *gin.Context) {
	ret := zh.InitQuery(ctx, "name")
	if ret == nil {
		return
	}
	nodes, err := zh.bnd.NodesInZone(ret.GetValues().Get("name"))
	if err != nil {
		ret.SetErrorCode(http.StatusBadRequest).SetError(err)
	} else {
		ret.SetPayload(gin.H{"nodes": nodes}).SendJSON()
	}
}

// AddZone godoc
// @Summary Define a cluster Zone.
// @Description AddZone creates a new empty zone in the cluster.
// @ID add-zone
// @Accept json
// @Produce json
// @Param name query string true "Name of the Zone"
// @Param description query string true "Zone description"
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/add [post]
func (zh *ZoneHandler) AddZone(ctx *gin.Context) {
	ret := zh.InitForm(ctx, "name", "description")
	if ret == nil {
		return
	}
	name := ctx.Request.Form.Get("name")
	descr := ctx.Request.Form.Get("description")
	err := zh.bnd.CreateZone(name, descr)
	if err != nil {
		ret.SetError(err).SetErrorCode(http.StatusNotAcceptable)
	} else {
		ret.SetMessage(fmt.Sprintf("Zone '%s' has been added", name))
	}
	ret.SendJSON()
}

// UpdateZone godoc
// @Summary Update a cluster Zone
// @Description UpdateZone updates a zone data,
// @ID update-zone
// @Accept json
// @Produce json
// @Param name query string true "Name of the Zone"
// @Param description query string true "Zone description"
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/update [post]
func (zh *ZoneHandler) UpdateZone(ctx *gin.Context) {
	ret := zh.InitForm(ctx, "name", "description")
	if ret == nil {
		return
	}

	name := ret.GetValues().Get("name")
	descr := ret.GetValues().Get("description")

	err := zh.bnd.UpdateZone(name, descr)
	if err != nil {
		ret.SetError(err).SetErrorCode(http.StatusBadRequest)
	} else {
		ret.SetMessage(fmt.Sprintf("Zone '%s' has been updated", name))
	}

	ret.SendJSON()
}

// RemoveZone godoc
// @Summary Remove an empty cluster Zone
// @Description RemoveZone removes a zone from the cluster, but only if it is empty (no nodes assigned to it).
// @ID remove-zone
// @Accept json
// @Produce json
// @Param name query string true "Name of the Zone"
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/remove [delete]
func (zh *ZoneHandler) RemoveZone(ctx *gin.Context) {
	ret := zh.InitBody(ctx, "name")
	if ret == nil {
		return
	}
	name := ret.GetValues().Get("name")
	err := zh.bnd.RemoveZone(name)
	if err != nil {
		ret.SetError(err).SetErrorCode(http.StatusBadRequest)
	} else {
		ret.SetMessage(fmt.Sprintf("Zone '%s' has been removed", name))
	}

	ret.SendJSON()
}

// ListZones godoc
// @Summary List cluster zones
// @Description List all zones in the Cluster.
// @ID list-zones
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "0"
// @Router /api/v1/zones/list [get]
func (zh *ZoneHandler) ListZones(ctx *gin.Context) {
	pl := make([]map[string]string, 0)
	for _, zone := range zh.bnd.ListZones() {
		pzone := make(map[string]string)
		pzone["name"] = zone.Name
		pzone["description"] = zone.Description
		pl = append(pl, pzone)
	}
	NewReturnType(ctx).SetErrorCode(http.StatusOK).SetPayload(pl).SendJSON()
}
