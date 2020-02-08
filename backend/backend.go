// Backend interface

package backend

import (
	"github.com/isbm/mgr-clbd/dbx"
)

type Backend interface {
	SetDbx(d *dbx.Dbx)
	StartUp()
}
