package mock

import (
	"github.com/PacketFire/immigrant/lib/core"
)

type MockDriver struct {
	Revisions []core.Revision
	State     core.Revision
}

func (this *MockDriver) Init(config map[string]string) error {
	return nil
}

func Migrate(r core.Revision, ec chan error) {
}

func Rollback(r core.Revision, ec chan error) {
}

func State() *core.Revision {
	return &core.Revision{}
}

func Close() {
}
