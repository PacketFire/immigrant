package version

import (
	"github.com/PacketFire/immigrant/pkg/config"
	"github.com/mitchellh/cli"
	"strings"
	"testing"
)

func TestHelpHasTabs(t *testing.T) {
	t.Parallel()
	if strings.ContainsRune(New(config.Config{}, cli.NewMockUi(), "").Help(), '\t') {
		t.Fatal("help has tabs")
	}
}
