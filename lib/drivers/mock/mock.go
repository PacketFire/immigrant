package mock

import (
	"github.com/PacketFire/immigrant/lib/core"
)

type MockDriver struct {
	Revisions []core.Revision
}

func (this *MockDriver) Init(config map[string]string) error {
	return nil
}

func (this *MockDriver) Migrate(r core.Revision, ec chan error) {
	for _, mig := range r.Migrate {
		if _, err = mig; err != nil {
			ec <- err
			return
		}
	}
}

func (this *MockDriver) Rollback(r core.Revision, ec chan error) {
	for _, mig := range r.Rollback {
		if _, err = mig; err != nil {
			ec <- err
			return
		}
	}
}

func (this *MockDriver) State() (*core.Revision, error) {
	return nil, &core.Revision{}
}

func Close() {
}
