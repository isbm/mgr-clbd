package backend

type Nodes struct {
	BaseBackend
}

func NewNodesBackend() *Nodes {
	return new(Nodes)
}

// Initialise cluster schema
func (n *Nodes) StartUp() {
	// Stupid, simple
	if !n.db.DB().HasTable(&ClusterNode{}) {
		n.db.DB().CreateTable(&ClusterNode{})
		logger.Println("Created ClusterNode table")
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
//
// All this is achieved by installing a ncd binary and let it
// run a nanostate.
func (n *Nodes) StageNode(node *ClusterNode) {
	logger.Debugln("Stage node")
}

// ConnectNode adds a node to the cluster. This should check if
// a node was successfully staged and the node must have staging status.
// Node connectivity is trusted by Node Controller, which defines
// if a node is ready for the cluster swarm.
func (n *Nodes) ConnectNode(node *ClusterNode) {
	logger.Debugln("Connect node")
}

// PoolNode only adds any candidate node to the pool,
// which still needs to be staged. This API call is indempotent
// and is *periodically* called by heartbit. It should skip adding
// another node nearby if the node with the same FP/machine-id was
// already added.
func (n *Nodes) PoolNode(node *ClusterNode) {
	logger.Println("Add node to the pool", node.Fqdn)
	n.db.DB().Create(node)
}

// WipePooledNodes removes all pooled nodes from the pool, whose
// timeout is expired. Example, when node was pooled but then turned off.
func (n *Nodes) WipePooledNodes() {

}

// ListNodes returns a list of all nodes in the system,
// included staged and new.
func (n *Nodes) ListAllNodes() {
	logger.Debugln("List all nodes")
	clusterNodes := []ClusterNode{}
	n.db.DB().Find(&clusterNodes)

	logger.Debugln("Found ", len(clusterNodes), "nodes in the cluster")
	for _, node := range clusterNodes {
		logger.Debugln("Node:", node.Fqdn)
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
