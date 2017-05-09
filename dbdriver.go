package main

// DBDriver defines the necessary methods to interface a database with
// immigrant. State tracking, Migrations and Rollbacks should mostly be
// defined and implemented by the driver.
type DBDriver interface {
	Init(map[string]string)  // Setup connection and state tracking.
	Migrate(Revision) error  // Execute a migration against the target database.
	Rollback(Revision) error // Execute a rollback against the target database.
	Close()                  // Close a connection to the target database.
}
