package backend

import (
	"fmt"
	"github.com/isbm/mgr-clbd/dbx"
	"log"
)

type Nodes struct {
	db *dbx.Dbx
}

func NewNodesBackend() *Nodes {
	return new(Nodes)
}

// SetDbx sets the dbx reference
func (n *Nodes) SetDbx(d *dbx.Dbx) {
	n.db = d
}

// Initialise cluster schema
func (n *Nodes) StartUp() {
	// Stupid, simple
	if !n.db.DB().HasTable(&ClusterNode{}) {
		n.db.DB().CreateTable(&ClusterNode{})
		log.Println("Created ClusterNode table")
	}

	// Automigrate
	n.db.DB().AutoMigrate(&ClusterNode{})
}

// StageNode runs node stanging by calling the Node Controller
// to wipe it out and prepare for being added to the cluster.
//
// To stage any Uyuni Server, one needs to do the following:
//   - Run Uyuni Server in cluster node mode
//   - Configure and start Node Controller on it
//
// After stage node message is received on Node Controller,
// it should:
//   - reset the database to the initial state (remove all
//     systems, channels etc)
//   - Confirm this done and ack the response
func (n *Nodes) StageNode(node *ClusterNode) {
	fmt.Println("Stage node")
}

// ConnectNode adds a node to the cluster. This should check if
// a node was successfully staged and the node must have staging status.
// Node connectivity is trusted by Node Controller, which defines
// if a node is ready for the cluster swarm.
func (n *Nodes) ConnectNode(node *ClusterNode) {
	fmt.Println("Connect node")
}

// PoolNode only adds any candidate node to the pool,
// which still needs to be staged. This API call is indempotent
// and is *periodically* called by heartbit. It should skip adding
// another node nearby if the node with the same FP/machine-id was
// already added.
func (n *Nodes) PoolNode(node *ClusterNode) {
	log.Println("Add node to the pool", node.Fqdn)
	n.db.DB().Create(node)
}

// WipePooledNodes removes all pooled nodes from the pool, whose
// timeout is expired. Example, when node was pooled but then turned off.
func (n *Nodes) WipePooledNodes() {

}

// ListNodes returns a list of all nodes in the system,
// included staged and new.
func (n *Nodes) ListAllNodes() {
	fmt.Println("List all nodes")
	clusterNodes := []ClusterNode{}
	n.db.DB().Find(&clusterNodes)

	fmt.Println("Found ", len(clusterNodes), "nodes in the cluster")
	for _, node := range clusterNodes {
		fmt.Print("Node:", node.Fqdn)
	}
}

// ListPoolNodes returns all nodes that are in the pool and yet to be staged.
func (n *Nodes) ListPoolNodes() {

}

// ListStagedNodes returns all nodes that are staged at the moment.
func (n *Nodes) ListStagedNodes() {

}

// ListClusterNodes returns all the nodes that are added to the current cluster.
func (n *Nodes) ListClusterNodes() {

}
