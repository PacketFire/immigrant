package mock

import (
	"errors"

	"github.com/PacketFire/immigrant/pkg/core"
)

// Driver implements github.com/PacketFire/immigrant/pkg/core.Driver, storing
// revisions to an in memory representation of the Revisions store.
type Driver struct {
	Revisions []*core.Revision
}

// Init mocks the requirements for Init on the
// github.com/PacketFire/immigrant/pkg/core.Driver interface. This method
// always returns a successful call.
func (dri *Driver) Init(config map[string]string) error {
	return nil
}

// Migrate mocks the requirements for Migrate on the
// github.com/PacketFire/immigrant/pkg/core.Driver interface. This method
// appends the past Revision to an in memory store and returns a success.
func (dri *Driver) Migrate(r core.Revision) error {
	dri.Revisions = append(dri.Revisions, &r)
	return nil
}

// Rollback mocks the requirements for Rollback on the
// github.com/PacketFire/immigrant/pkg/core.Driver interface. This method pops
// the last executed method off of the in memory store, returning a success if
// one exists or an error if one doesn't.
func (dri *Driver) Rollback(r core.Revision) error {
	if len(dri.Revisions) == 0 {
		return errors.New("no revisions applied")
	}

	dri.Revisions = dri.Revisions[:len(dri.Revisions)-1]
	return nil
}

// State mocks the requirements for State on the
// github.com/PacketFire/immigrant/pkg/core.Driver interface. This method
// retrieves the previous revision from the internal representation of the
// revision map. If the Revision list is empty, nil is returned.
func (dri *Driver) State() *core.Revision {
	rtl := len(dri.Revisions)
	if rtl == 0 {
		return nil
	}

	return dri.Revisions[rtl-1]
}

// Close mocks the requirements for Close on the
// github.com/PacketFire/immigrant/pkg/core.Driver interface. This method
// simply functions as a noop.
func (dri *Driver) Close() {}
