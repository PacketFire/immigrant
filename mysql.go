package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDriver struct {
	Config DSN
	Db     *sql.DB
}

type DSN struct {
	user     string
	pass     string
	host     string
	proto    string
	database string
	params   string
}

func (this *MysqlDriver) Init(ctx map[string]string) error {
	this.Config = NewDSN(ctx["username"],
		ctx["password"],
		ctx["proto"],
		ctx["host"],
		ctx["database"],
		ctx["params"])

	db, err := sql.Open("mysql", this.Config.String())
	if err != nil {
		return err
	}

	this.Db = db

	return nil
}

func (this *MysqlDriver) Migrate(r Revision, c chan error) {

}

func (this *MysqlDriver) Rollback(r Revision, c chan error) {

}

func (this *MysqlDriver) State() *Revision {

	return new(Revision)
}

func (this *MysqlDriver) Close() {
	this.Db.Close()
}
