package dbx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	d := new(Dbx)
	d._host = "localhost"
	d._port = 3306
	return d
}

// SetUser sets the username for the database connection credentials
func (d *Dbx) SetUser(user string) *Dbx {
	d._user = user
	return d
}

// SetPassword sets the password credentials for the database connection
func (d *Dbx) SetPassword(password string) *Dbx {
	d._password = password
	return d
}

// SetDBHost sets the hostname of the remote database URI
func (d *Dbx) SetDBHost(fqdn string) *Dbx {
	d._host = fqdn
	return d
}

// SetPort sets the connection port of the remote database URI
func (d *Dbx) SetDBPort(port int) *Dbx {
	d._port = port
	return d
}

// SetDBName sets the name of the database
func (d *Dbx) SetDBName(dbname string) *Dbx {
	d._dbname = dbname
	return d
}

// Close function closes the database connection
func (d *Dbx) Close() error {
	var err error
	if d.db != nil {
		err = d.db.Close()
	}
	return err
}

// Open function sets the whole connection to the database opened state
func (d *Dbx) Open() error {
	var err error
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d._user, d._password, d._host, d._port, d._dbname)
	d.db, err = gorm.Open("mysql", uri)
	if err == nil {
		d.db.DB().SetMaxIdleConns(50)
	}
	return err
}

// DB returns the database connection. If database is not opened yet, it will automatically open a new connection.
func (d *Dbx) DB() *gorm.DB {
	if d.db == nil {
		err := d.Open()
		if err != nil {
			panic(err.Error())
		}
	}
	return d.db
}
