package main

import (
	"fmt"
)

type DSN struct {
	user     string
	pass     string
	host     string
	proto    string
	database string
	params   string
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
