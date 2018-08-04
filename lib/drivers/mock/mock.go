package mock

import (
	"github.com/ncatelli/immigrant/lib/core"
)

type MockDriver struct {
	Revisions []core.Revision
}

func (this *MockDriver) Init(config map[string]string) error {
	return nil
}

func Migrate(r core.Revision, ec chan error) {
	for _, mig := range r.Migrate {
		if _, err = mig; err != nil {
			c <- err
			return
		}
	}
}

func Rollback(r core.Revision, ec chan error) {
	for _, mig := range r.Rollback {
		if _, err = mig; err != nil {
			c <- err
			return
		}
	}
}

func State(this *MockDriver) *core.Revision {
	fmt.Println(&core.Revision)
	return &core.Revision{}

}

func Close() {
}
