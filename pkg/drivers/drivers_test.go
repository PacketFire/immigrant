package drivers

import (
	"reflect"
	"testing"

	"github.com/PacketFire/immigrant/pkg/config"
)

const errFmt string = "expected %v got %v"

func TestPassingAConfigWithExplicitTypeToGenerateDriverFromConfigMethodShould(t *testing.T) {
	t.Run("return a mock driver when type is mock.", func(t *testing.T) {
		c := config.Config{
			"type": "mock",
		}

		d, _ := GenerateDriverFromConfig(c)
		dType := reflect.TypeOf(d).String()
		if dType != "*mock.Driver" {
			t.Errorf(errFmt, "*mock.Driver", dType)
		}
	})
}
