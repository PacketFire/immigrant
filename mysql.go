package main

import (
	"database/sql"
	"fmt"
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

func (This *MysqlDriver) Close() {

}

func NewDSN(user, pass, proto, host, database, params string) DSN {
	if proto == "" {
		proto = "tcp"
	}

	return DSN{
		user:     user,
		pass:     pass,
		host:     host,
		database: database,
		params:   params,
	}
}

func (this DSN) String() string {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s)",
		this.user,
		this.pass,
		this.proto,
		this.host,
		this.database)
	if this.params != "" {
		dsn = fmt.Sprintf("%s?%s", dsn, this.params)
	}

	return dsn
}
