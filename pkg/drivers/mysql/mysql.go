package mysql

import (
	"database/sql"
	"encoding/json"

	"github.com/PacketFire/immigrant/pkg/core"
	"github.com/PacketFire/immigrant/pkg/drivers"
)

const (
	stateCreate string = `CREATE TABLE imm_sequence_tracker (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  revisionID VARCHAR(256),
  revisionJSON TEXT
);
`
)

// ERRORS

// errCurrentRemoteState is returned when immigrant is unable to fetch the
// remote state's HEAD.
type errCurrentRemoteState struct{}

func (e errCurrentRemoteState) Error() string {
	return "Unable to fetch remote revision state."
}

type errHeadDoesNotExist struct{}

func (e errHeadDoesNotExist) Error() string {
	return "Remote revision HEAD does not exist."
}

// Type Defs

type stateTrackerRevision struct {
	ID           int
	RevisionID   string
	RevisionJSON string
}

// Driver stores configuration information along with the DB connection
// management struct.
type Driver struct {
	Config drivers.DSN
	Db     *sql.DB
}

// Init takes a context with all configuration which it then uses to create the
// DB connection and bootstrap immigrant's state tracker table. on success nil
// is returned. On failure, the corresponding errors are returned.
func (dri *Driver) Init(ctx map[string]string) error {
	dri.Config = drivers.NewDSN(ctx["username"],
		ctx["password"],
		ctx["proto"],
		ctx["host"],
		ctx["database"],
		ctx["params"])

	db, err := sql.Open("mysql", dri.Config.String())
	if err != nil {
		return err
	}

	dri.Db = db
	if err = db.Ping(); err != nil {
		return err
	}

	if err = dri.initStateManager(); err != nil {
		return err
	}

	return nil
}

// Migrate takes a revision and a channel and attempts to execute all migrate
// statements defined in the revision object. The channel is primarily used to
// communicate back to the cli tool whether a migration has completed in case
// the signal handler is invoked. On success, nil is pushed over the channel.
// On failure, an error is pushed over the channel.
func (dri *Driver) Migrate(r core.Revision, c chan error) {
	tx, err := dri.Db.Begin()
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
func (dri *Driver) Rollback(r core.Revision, c chan error) {
	tx, err := dri.Db.Begin()
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
// pointer to a populated Revision is return and nil is returned. On failure,
// nil and an error are returned.
func (dri *Driver) State() (*core.Revision, error) {
	rHead := new(core.Revision)

	rows, err := dri.Db.Query("SELECT * FROM imm_sequence_tracker ORDER BY id DESC LIMIT 0, 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := new(stateTrackerRevision)
		if err = rows.Scan(row); err != nil {
			return nil, errCurrentRemoteState{}
		}

		if err = json.Unmarshal([]byte(row.RevisionJSON), rHead); err != nil {
			return nil, err
		}

		return rHead, nil
	}

	return nil, errHeadDoesNotExist{}
}

// initStateManager attempts to create the state tracker table and is only meant
// to be called by the Init method. On success, nil is returned. On failure an
// error is returned.
func (dri *Driver) initStateManager() error {
	stmt, err := dri.Db.Prepare(stateCreate)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

// Close closes the DB object associated with the driver.
func (dri *Driver) Close() {
	dri.Db.Close()
}
