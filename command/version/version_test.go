package version

import (
	"github.com/PacketFire/immigrant/command/context"
	"github.com/mitchellh/cli"
	"strings"
	"testing"
)

func TestHelpHasTabs(t *testing.T) {
	t.Parallel()
	if strings.ContainsRune(New(context.Context{}, cli.NewMockUi(), "").Help(), '\t') {
		t.Fatal("help has tabs")
	}
}
