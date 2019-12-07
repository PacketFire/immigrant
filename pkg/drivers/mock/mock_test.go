package mock

import (
	"reflect"
	"testing"

	"github.com/PacketFire/immigrant/pkg/core"
)

var ec chan error
var name string
var this Driver

func TestDriver_Migrate(t *testing.T) {
	name = "migrate"
	rs := core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run(name, func(t *testing.T) {
		this.Migrate(rs)

		if len(this.Revisions) == 1 {
			if !reflect.DeepEqual(this.Revisions[0], rs) {
				t.Log("failed")
			}
		}
	})
}

func TestDriverRollbackMethodShould(t *testing.T) {
	r := core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run("pop the most recent revision off of the in memory state if state exists", func(t *testing.T) {
		var dri Driver
		dri.Revisions = []*core.Revision{&r}

		if err := dri.Rollback(r); err != nil {
			t.Errorf("expected %v got %v", nil, err)
		}

		rlen := len(dri.Revisions)
		if rlen != 0 {
			t.Errorf("expected %v got %v", 0, rlen)
		}
	})
}

func TestDriverStateMethodShould(t *testing.T) {
	r := &core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run("return latest revision when revision state exists", func(t *testing.T) {
		var dri Driver
		dri.Revisions = []*core.Revision{r}
		sr := dri.State()
		if !reflect.DeepEqual(*r, *sr) {
			t.Errorf("expected %v got %v", *r, *sr)
		}
	})

	t.Run("return nil when revision state doesn't exist", func(t *testing.T) {
		var dri Driver
		dri.Revisions = []*core.Revision{}
		sr := dri.State()
		if sr != nil {
			t.Errorf("expected %v got %v", nil, sr)
		}
	})
}
