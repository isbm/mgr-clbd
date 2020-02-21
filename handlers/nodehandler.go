/*
Handler, which among other things, interfaces the internal
micro-configuration management system. Micro CMS is running
nanostates to call preinstalled Ansible modules (or facilitate
own to be deployed to the client machines).
*/

package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/isbm/mgr-clbd/backend"
	"github.com/isbm/mgr-clbd/dbx"
	"net/http"
)

type NodeHandler struct {
	BaseHandler
	db  *dbx.Dbx
	orm *backend.Nodes
	cms *backend.NanoCms
}

func NewNodeHandler(root string) *NodeHandler {
	nh := new(NodeHandler)
	nh.PrepareRoot(root)
	nh.orm = backend.NewNodesBackend()
	nh.cms = backend.NewNanoCmsBackend()
	return nh
}

func (nh *NodeHandler) Backend() backend.Backend {
	// Backend is called for booting it up.
	// Here is a hook to apply passed-on configuration
	stateroot := nh.config.Find("general:state").String("root", "")
	if stateroot != "" {
		nh.cms.GetStateIndex().AddStateRoot(stateroot).Index()
		nh.cms.SetStaticDataRoot(nh.config.Find("general").String("static-root", ""))
	} else {
		panic("Stateroot and Bootstrap Id must be specified!")
	}

	return nh.orm
}

// SetDbx sets the Dbx instance pointer
func (nh *NodeHandler) SetDbx(db *dbx.Dbx) {
	nh.db = db
	nh.orm.SetDbx(nh.db)
}

// Handlers returns a map of supported handlers and their configuration
func (nh *NodeHandler) Handlers() []*HandlerMeta {
	return []*HandlerMeta{
		&HandlerMeta{
			Route:   nh.ToRoute("list"),
			Handle:  nh.ListNodes,
			Methods: []string{POST, GET},
		},
		&HandlerMeta{
			Route:   nh.ToRoute("stage"),
			Handle:  nh.StageNode,
			Methods: []string{POST},
		},
		&HandlerMeta{
			Route:   nh.ToRoute("add"),
			Handle:  nh.AddNode,
			Methods: []string{POST},
		},
	}
}

// ListNodes godoc
// @Summary List nodes in the cluster.
// @Description List all nodes in the current cluster.
// @ID list-nodes
// @Accept json
// @Produce json
// @Header 200 {string} Token "0"
// @Router /api/v1/node/list [get]
func (nh *NodeHandler) ListNodes(ctx *gin.Context) {
	nh.orm.ListAllNodes()
	ctx.JSON(http.StatusOK, gin.H{
		"nodes": "listed",
	})
}

// AddNode godoc
// @Summary Add a new bootstrapped, ready node to the cluster.
// @Description This will verify if a node meets the requirements and will join to the cluster.
// @ID add-node
// @Accept json
// @Produce json
// @Param fqdn query string true "FQDN of the cluster node for adding it"
// @Header 200 {string} Token "0"
// @Router /api/v1/node/add [post]
func (nh *NodeHandler) AddNode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"node": "added",
	})
}

// StageNode godoc
// @Summary Stage (bootstrap) a new cluster node.
// @Description This will install a client binary over SSH and will run nanostate, required to setup everything in place
// @ID stage-node
// @Accept json
// @Produce json
// @Param fqdn query string true "FQDN of the hostname for staging"
// @Param password query string true "Root password of the node"
// @Param state query string true "State Id for bootstrapping"
// @Header 200 {string} Token "0"
// @Router /api/v1/node/stage [post]
func (nh *NodeHandler) StageNode(ctx *gin.Context) {
	ret := nh.InitForm(ctx, "fqdn", "password", "state")
	fqdn := ret.GetValues().Get("fqdn")
	passwd := ret.GetValues().Get("password")
	state := ret.GetValues().Get("state")

	ret.SetPayload(nh.cms.Bootstrap(fqdn, state, "root", passwd))
	ret.SendJSON()
}
