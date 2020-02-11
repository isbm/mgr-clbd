package backend

import (
	"fmt"
)

type Zones struct {
	BaseBackend
}

func NewZonesBackend() *Zones {
	return new(Zones)
}

func (z *Zones) StartUp() {
	z.VerifyTables(&ClusterZone{})
}

// CreateZone creates a new zone
func (z *Zones) CreateZone(name string, descr string) error {
	var err error
	zone := &ClusterZone{
		Name: name, Description: descr,
	}
	if z.db.DB().NewRecord(zone) {
		if err = z.db.DB().Create(&zone).Error; err != nil {
			logger.Errorln(err.Error())
		} else {
			logger.Infof("Added a new Zone %s", name)
		}
	} else {
		err = fmt.Errorf("Unable to add a new Zone '%s'", name)
		logger.Errorln(err.Error())
	}

	return err
}

// ListZones lists available zones
func (z *Zones) ListZones() []ClusterZone {
	zones := []ClusterZone{}
	z.db.DB().Find(&zones)
	return zones
}

// UpdateZone updates an existing zone.
func (z *Zones) UpdateZone(name string, descr string) error {
	var err error
	zone := &ClusterZone{}
	if err = z.db.DB().Where("name = ?", name).First(&zone).Error; err != nil {
		logger.Errorln(err.Error())
	} else {
		zone.Description = descr
		z.db.DB().Save(&zone)
		logger.Infof("Zone '%s' has been updated", name)
	}

	return err
}

// RemoveZone removes an empty zone. If zone still contains nodes,
// it won't be removed, but error will be issued.
func (z *Zones) RemoveZone(name string) error {
	nodesCnt, err := z.NodesInZone(name)
	if err != nil {
		return err
	}

	if nodesCnt != 0 {
		err = fmt.Errorf("Zone still contains %d nodes", nodesCnt)
		logger.Errorf("Attempt to delete non-empty zone '%s'", name)
		return err
	}

	if err := z.db.DB().Where("name = ?", name).Delete(&ClusterZone{}).Error; err != nil {
		logger.Errorf("Error deleting node %s: %s", name, err.Error())
	} else {
		logger.Infof("Node %s has been deleted", name)
	}

	return err
}

// NodesInZone returns an amount of attached cluster nodes to the zone
func (z *Zones) NodesInZone(name string) (int, error) {
	var err error
	zone := &ClusterZone{}
	z.db.DB().Where("name = ?", name).First(&zone)
	if zone.ID == 0 && zone.Name == "" {
		err = fmt.Errorf("Zone %s was not found", name)
		logger.Errorln(err.Error())
		return -1, err
	}

	nodes := make([]ClusterNode, 0)
	z.db.DB().Where("name = ?", name).Find(&nodes)
	return len(nodes), err
}
