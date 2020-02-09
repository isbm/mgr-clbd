package backend

import (
	"github.com/iancoleman/strcase"
	"github.com/isbm/mgr-clbd/dbx"
	"github.com/isbm/mgr-clbd/utils"
	"github.com/sirupsen/logrus"
	"reflect"
)

var logger *logrus.Logger

func init() {
	logger = utils.GetTextLogger(logrus.DebugLevel, nil)
}

type BaseBackend struct {
	db *dbx.Dbx
}

// SetDbx sets the dbx reference
func (bb *BaseBackend) SetDbx(d *dbx.Dbx) {
	bb.db = d
}

func (bb *BaseBackend) VerifyTables(tables ...interface{}) {
	for _, table := range tables {
		tableName := strcase.ToScreamingSnake(reflect.TypeOf(table).Elem().Name())
		if !bb.db.DB().HasTable(table) {
			bb.db.DB().CreateTable(table)
			logger.Debugln("Creating table", tableName)
		}

		// Automigrate
		bb.db.DB().AutoMigrate(table)
		logger.Debugln("Verified table", tableName)
	}
}
