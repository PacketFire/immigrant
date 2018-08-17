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
	go func() {
		this.Revisions = append(this.Revisions, *r)
		ec <- nil
	}()
}

func (this *MockDriver) Rollback(r *core.Revision, ec chan error) {
	go func() {
		if len(this.Revisions) == 0 {
			ec <- errors.New("No revisions applied")
			return
		} else {
			this.Revisions = this.Revisions[:len(this.Revisions)-1]
			ec <- nil
		}
	}()
}

func (this *MockDriver) State() *core.Revision {
	return &this.Revisions[len(this.Revisions)-1]
}

func Close() {
}
