package context

import (
	"github.com/PacketFire/immigrant/pkg/config"
	"github.com/PacketFire/immigrant/pkg/core"
)

// Context stores a database, configuration data and arbitrary parameters to be
// shared between the cli.
type Context struct {
	Driver core.Driver
	Config config.Config
}
