package mock

import (
	"errors"

	"github.com/PacketFire/immigrant/lib/core"
)

type MockDriver struct {
	Revisions []core.Revision
}

func (this *MockDriver) Init(config map[string]string) error {
	return nil
}

func (this *MockDriver) Migrate(r *core.Revision, ec chan error) {
	this.Revisions = append(this.Revisions, *r)
	ec <- nil
}
func (this *MockDriver) Rollback(r *core.Revision, ec chan error) {
	if len(this.Revisions) <= 0 {
		ec <- errors.New("No revisions applied.")
		return
	}
	this.Revisions = this.Revisions[:len(this.Revisions)-1]
	ec <- nil
}

func (this *MockDriver) State(r *core.Revision, ec chan error) {
	if len(this.Revisions) <= 0 {
		ec <- errors.New("No revisions found.")
	}
	this.Revisions = this.Revisions
	ec <- nil
}

func Close() {
}
