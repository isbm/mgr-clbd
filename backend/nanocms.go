package backend

import (
	"github.com/isbm/nano-cms/nanocms"
)

/*
Integration with the nanocms.
*/

type NanoCmsBackend struct {
	BaseBackend
	cms *nanocms_backend.NanoCms
}

func NewNanoCmsBackend() *NanoCmsBackend {
	ncb := new(NanoCmsBackend)
	ncb.cms = nanocms_backend.NewNanoCms()

	return ncb
}

// GetCms returns an embedded CMS instance (configuration management system)
func (ncb *NanoCmsBackend) GetCms() *nanocms_backend.NanoCms {
	return ncb.cms
}
