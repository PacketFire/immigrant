package main

import (
	"fmt"
)

// DSN stores connection information used to create a sql.DB object.
type DSN struct {
	user     string // username for connecting to the database.
	pass     string // password for the connecting user.
	host     string // hostname of database to connect to.
	proto    string // protocol for connection ex: tcp
	database string // database name.
	params   string // additional parameters.
}

// NewDSN takes 6 strings representing the values for generating a DSN string.
// A corresponding DSN object is returned.
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

// Converts the DSN struct to a DSN string.
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
