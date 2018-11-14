package sqlite

import (
	"database/sql"
	"encoding/json"

	"github.com/PacketFire/immigrant/lib/core"
	_ "github.com/mattn/go-sqlite3"
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

func (this errCurrentRemoteState) Error() string {
	return "Unable to fetch remote revision state."
}

type errHeadDoesNotExist struct{}

func (this errHeadDoesNotExist) Error() string {
	return "Remote revision HEAD does not exist."
}

// Type Defs

type stateTrackerRevision struct {
	Id           int
	RevisionID   string
	RevisionJSON string
}

type SqliteDriver struct {
	Db        *sql.DB
	Revisions []core.Revision
}

func (this *SqliteDriver) Init(filepath string) error {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return err
	}

	this.Db = db
	return nil
}

func (this *SqliteDriver) Migrate(r core.Revision) {
	this.Revisions = append(this.Revisions, r)
	tx, err := this.Db.Begin()
	if err != nil {
		return
	}

	for _, mig := range r.Migrate {
		if _, err = tx.Exec(mig); err != nil {
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()
}

func (this *SqliteDriver) Rollback(r core.Revision) {
	if len(this.Revisions) == 0 {
		return
	} else {
		this.Revisions = this.Revisions[:len(this.Revisions)-1]
	}

	tx, err := this.Db.Begin()
	if err != nil {
		return
	}

	for _, mig := range r.Rollback {
		if _, err = tx.Exec(mig); err != nil {
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()
}

func (this *SqliteDriver) State() (*core.Revision, error) {
	rHead := new(core.Revision)

	rows, err := this.Db.Query("SELECT * FROM imm_sequence_tracker ORDER BY id DESC LIMIT 0, 1")
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

func (this *SqliteDriver) initStateManager() error {
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

func (this *SqliteDriver) Close() {
	this.Db.Close()
}
