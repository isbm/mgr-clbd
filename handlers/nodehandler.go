package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
	"net/http"
)

type NodeHandler struct {
	db       *dbx.Dbx
	bknNodes *backend.Nodes
}

func NewNodeHandler() *NodeHandler {
	nh := new(NodeHandler)
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
			Route:   "nodes/list",
			Handle:  nh.OnListNodes,
			Methods: []string{POST, GET},
		},
		&HandlerMeta{
			Route:   "nodes/add",
			Handle:  nh.OnAddNode,
			Methods: []string{POST},
		},
	}
}

// Handle implements the entry point of the handler
func (nh *NodeHandler) OnListNodes(ctx *gin.Context) {
	nh.bknNodes.ListNodes()
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
