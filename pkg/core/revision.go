package core

// Revision represents a versioned change to the database and includes fields
// for a revision name, and a list of both migration and rollback steps.
type Revision struct {
	Revision string   `yaml:"revision"` // A unique name for the revision.
	Migrate  []string `yaml:"migrate"`  // A list of migration steps.
	Rollback []string `yaml:"rollback"` // A list of rollback steps.
}
