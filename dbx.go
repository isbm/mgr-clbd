package clbd

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Dbx struct {
	_user     string
	_password string
	_host     string
	_dbname   string
	_port     int
	db        *gorm.DB
}

// NewDbx is a constructor for database connection class
func NewDbxConnection() *Dbx {
	dbx := new(Dbx)
	dbx._host = "localhost"
	dbx._port = 3306
	return dbx
}

// SetUser sets the username for the database connection credentials
func (dbx *Dbx) SetUser(user string) *Dbx {
	dbx._user = user
	return dbx
}

// SetPassword sets the password credentials for the database connection
func (dbx *Dbx) SetPassword(password string) *Dbx {
	dbx._password = password
	return dbx
}

// SetDBHost sets the hostname of the remote database URI
func (dbx *Dbx) SetDBHost(fqdn string) *Dbx {
	dbx._host = fqdn
	return dbx
}

// SetPort sets the connection port of the remote database URI
func (dbx *Dbx) SetDBPort(port int) *Dbx {
	dbx._port = port
	return dbx
}

// SetDBName sets the name of the database
func (dbx *Dbx) SetDBName(dbname string) *Dbx {
	dbx._dbname = dbname
	return dbx
}

// Close function closes the database connection
func (dbx *Dbx) Close() error {
	var err error
	if dbx.db != nil {
		err = dbx.db.Close()
	}
	return err
}

// Open function sets the whole connection to the database opened state
func (dbx *Dbx) Open() error {
	var err error
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbx._user, dbx._password, dbx._host, dbx._port, dbx._dbname)
	dbx.db, err = gorm.Open("mysql", uri)
	if err == nil {
		dbx.db.DB().SetMaxIdleConns(50)
	}
	return err
}

// DB returns the database connection. If database is not opened yet, it will automatically open a new connection.
func (dbx *Dbx) DB() *gorm.DB {
	if dbx.db == nil {
		err := dbx.Open()
		if err != nil {
			panic(err.Error())
		}
	}
	return dbx.db
}
