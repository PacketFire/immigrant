package mock

import (
	"reflect"
	"testing"

	"github.com/PacketFire/immigrant/lib/core"
)

var ec chan error
var name string
var this *MockDriver

func TestMockDriver_Migrate(t *testing.T) {
	name = "revision 1"
	rs := &core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run(name, func(t *testing.T) {
		this.Migrate(rs, ec)
	})
}

func TestMockDriver_Rollback(t *testing.T) {
	name = "revision 1"
	rs := &core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run(name, func(t *testing.T) {
		this.Rollback(rs, ec)
	})
}

func TestMockDriver_State(t *testing.T) {
	tests := []struct {
		name string
		this *MockDriver
		want *core.Revision
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.State(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDriver.State() = %v, want %v", got, tt.want)
			}
		})
	}
}
