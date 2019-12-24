package context

import (
	"testing"
	"github.com/PacketFire/immigrant/pkg/config"
)

var errFmt string = "expected %v got %v"

func TestContextShouldInit(t *testing.T) {
	t.Run("without a defined driver", func(t *testing.T) {
		conf := make(config.Config)
		ctx := Context{Config: conf}
		ctx.Config["type"] = "mock"

		if ctx.Config == nil {
			t.Errorf(errFmt, make(config.Config), conf)
		}
	})
}
