package main

// Driver defines the necessary methods to interface a database with
// immigrant. State tracking, Migrations and Rollbacks should mostly be
// defined and implemented by the driver.
type Driver interface {
	Init(map[string]string) error  // Setup connection and state tracking.
	Migrate(Revision, chan error)  // Execute a migration against the target database.
	Rollback(Revision, chan error) // Execute a rollback against the target database.
	State() (*Revision, error)     // State returns the current revision of the database.
	Close()                        // Close a connection to the target database.
}
