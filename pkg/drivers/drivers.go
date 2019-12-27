package drivers

import (
	"fmt"
	"github.com/PacketFire/immigrant/pkg/config"
	"github.com/PacketFire/immigrant/pkg/core"
	"github.com/PacketFire/immigrant/pkg/drivers/mock"
)

// ErrUnknownDriverType is returned when an unspecified driver type is
// designated.
type ErrUnknownDriverType struct {
	t string
}

func (e *ErrUnknownDriverType) Error() string {
	return fmt.Sprintf("unknown driver type %v", e.t)
}

// GenerateDriverFromConfig takes a config and attempts to return a
// corresponding driver as derived from the config.
func GenerateDriverFromConfig(c config.Config) (core.Driver, error) {
	dt, e := c.DriverType()
	if e != nil {
		return nil, e
	}

	switch dt {
	case "mock":
		return &mock.Driver{}, nil
	}

	return nil, &ErrUnknownDriverType{dt}
}
