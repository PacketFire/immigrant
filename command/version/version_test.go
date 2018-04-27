package version

import (
	"github.com/mitchellh/cli"
	"strings"
	"testing"
)

func TestHelpHasTabs(t *testing.T) {
	t.Parallel()
	if strings.ContainsRune(New(cli.NewMockUi(), "").Help(), '\t') {
		t.Fatal("help has tabs")
	}
}
