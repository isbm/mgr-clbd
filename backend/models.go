// Backend DB models

package backend

const (
	STATUS_NEW = iota
	STATUS_ORPHAN
	STATUS_ACCEPTED
	STATUS_DELETED
)

// Cluster node object
type ClusterNode struct {
	ID            int64
	Fqdn          string
	MachineId     string // systemd
	ClientSystems int64
	LoadAverage   float64 // 5 min
}

// Client system object
type ClientSystem struct {
	ID        int64
	PubkeyFp  string
	Fqdn      string
	MachineId string
	Status    int
}
