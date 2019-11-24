package core

// Driver defines the necessary methods to interface a database with
// immigrant. State tracking, Migrations and Rollbacks should mostly be
// defined and implemented by the driver.
type Driver interface {
	Init(map[string]string) error // Setup connection and state tracking.
	Migrate(Revision) error       // Execute a migration against the target database.
	Rollback(Revision) error      // Execute a rollback against the target database.
	State() *Revision             // State returns the current revision of the database.
	Close()                       // Close a connection to the target database.
}
