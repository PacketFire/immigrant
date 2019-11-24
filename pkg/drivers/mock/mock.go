package mock

import (
	"errors"

	"github.com/PacketFire/immigrant/pkg/core"
)

type Driver struct {
	Revisions []core.Revision
}

func (dri *Driver) Init(config map[string]string) error {
	return nil
}

func (dri *Driver) Migrate(r core.Revision) error {
	dri.Revisions = append(dri.Revisions, r)
	return nil
}

func (dri *Driver) Rollback(r core.Revision) error {
	if len(dri.Revisions) == 0 {
		return errors.New("No revisions applied.")
	} else {
		dri.Revisions = dri.Revisions[:len(dri.Revisions)-1]
	}

	return nil
}

func (dri *Driver) State() *core.Revision {
	return &dri.Revisions[len(dri.Revisions)-1]
}

func Close() {
}
