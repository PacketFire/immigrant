package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	stateCreate string = `CREATE DATABASE immStateMan if not exists;
use immStateMan;
CREATE TABLE sequence_tracker (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  revisionID VARCHAR(256),
  revisionJSON TEXT
);
`
)

// MysqlDriver stores configuration information along with the DB connection
// management struct.
type MysqlDriver struct {
	Config DSN
	Db     *sql.DB
}

// Init takes a context with all configuration which it then uses to create the
// DB connection and bootstrap immigrant's state tracker database. on success
// nil is returned. On failure, the corresponding errors are returned.
func (this *MysqlDriver) Init(ctx map[string]string) error {
	this.Config = NewDSN(ctx["username"],
		ctx["password"],
		ctx["proto"],
		ctx["host"],
		"/",
		ctx["params"])

	db, err := sql.Open("mysql", this.Config.String())
	if err != nil {
		return err
	}

	this.Db = db
	if err = db.Ping(); err != nil {
		return err
	}

	if err = this.initStateManager(); err != nil {
		return err
	}

	return nil
}

// Migrate takes a revision and a channel and attempts to execute all migrate
// statements defined in the revision object. The channel is primarily used to
// communicate back to the cli tool whether a migration has completed in case
// the signal handler is invoked. On success, nil is pushed over the channel.
// On failure, an error is pushed over the channel.
func (this *MysqlDriver) Migrate(r Revision, c chan error) {
	tx, err := this.Db.Begin()
	if err != nil {
		c <- err
		return
	}

	for _, mig := range r.Migrate {
		if _, err = tx.Exec(mig); err != nil {
			tx.Rollback()
			c <- err
			return
		}
	}

	err = tx.Commit()
	c <- err
}

// Rollback takes a revision and a channel and attempts to execute all rollback
// statements defined in the revision object. The channel is primarily used to
// communicate back to the cli tool whether a migration has completed in case
// the signal handler is invoked. On success, nil is pushed over the channel.
// On failure, an error is pushed over the channel.
func (this *MysqlDriver) Rollback(r Revision, c chan error) {
	tx, err := this.Db.Begin()
	if err != nil {
		c <- err
		return
	}

	for _, mig := range r.Rollback {
		if _, err = tx.Exec(mig); err != nil {
			tx.Rollback()
			c <- err
			return
		}
	}

	err = tx.Commit()
	c <- err
}

// State returns the current revision that the database is at. On success a
// pointer to a populated Revision is return. On failure, nil is returned.
func (this *MysqlDriver) State() *Revision {

	return new(Revision)
}

// initStateManager attempts to create the state tracker database and tables
// and is only meant to be called by the Init method. On success, nil is
// returned. On failure an error is returned.
func (this *MysqlDriver) initStateManager() error {
	stmt, err := this.Db.Prepare(stateCreate)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

// Closes the DB object associated with the driver.
func (this *MysqlDriver) Close() {
	this.Db.Close()
}
