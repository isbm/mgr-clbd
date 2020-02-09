package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
	"net/http"
)

type NodeHandler struct {
	BaseHandler
	db       *dbx.Dbx
	bknNodes *backend.Nodes
}

func NewNodeHandler(root string) *NodeHandler {
	nh := new(NodeHandler)
	nh.PrepareRoot(root)
	nh.bknNodes = backend.NewNodesBackend()
	return nh
}

func (ph *NodeHandler) Backend() backend.Backend {
	return ph.bknNodes
}

// SetDbx sets the Dbx instance pointer
func (nh *NodeHandler) SetDbx(db *dbx.Dbx) {
	nh.db = db
	nh.bknNodes.SetDbx(nh.db)
}

// Handlers returns a map of supported handlers and their configuration
func (nh *NodeHandler) Handlers() []*HandlerMeta {
	return []*HandlerMeta{
		&HandlerMeta{
			Route:   nh.ToRoute("list"),
			Handle:  nh.OnListNodes,
			Methods: []string{POST, GET},
		},
		&HandlerMeta{
			Route:   nh.ToRoute("add"),
			Handle:  nh.OnAddNode,
			Methods: []string{POST},
		},
	}
}

// Handle implements the entry point of the handler
func (nh *NodeHandler) OnListNodes(ctx *gin.Context) {
	nh.bknNodes.ListAllNodes()
	ctx.JSON(http.StatusOK, gin.H{
		"nodes": "listed",
	})
}

// Handle implements the entry point of the handler
func (nh *NodeHandler) OnAddNode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"node": "added",
	})
}
