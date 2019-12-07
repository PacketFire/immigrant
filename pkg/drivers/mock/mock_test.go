package mock

import (
	"errors"
	"reflect"
	"testing"

	"github.com/PacketFire/immigrant/pkg/core"
)

var errFmt string = "expected %v got %v"

func TestDriverInitMethodShould(t *testing.T) {
	t.Run("return nil", func(t *testing.T) {
		var dri Driver
		conf := make(map[string]string)

		if rv := dri.Init(conf); rv != nil {
			t.Errorf(errFmt, nil, rv)
		}
	})
}

func TestDriverMigrateMethodShould(t *testing.T) {
	r := core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run("append a revision to the state when invoked", func(t *testing.T) {
		var dri Driver
		dri.Revisions = []*core.Revision{}

		dri.Migrate(r)
		if len(dri.Revisions) != 1 || !reflect.DeepEqual(*dri.Revisions[0], r) {
			t.Errorf(errFmt, r, *dri.Revisions[0])
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
			t.Errorf(errFmt, nil, err)
		}

		rlen := len(dri.Revisions)
		if rlen != 0 {
			t.Errorf(errFmt, 0, rlen)
		}
	})

	t.Run("throw an error if a rollback is run against a driver with no state", func(t *testing.T) {
		var dri Driver
		dri.Revisions = []*core.Revision{}

		if err := dri.Rollback(r); err == nil {
			t.Errorf(errFmt, errors.New("no revisions applied"), err)
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
			t.Errorf(errFmt, *r, *sr)
		}
	})

	t.Run("return nil when revision state doesn't exist", func(t *testing.T) {
		var dri Driver
		dri.Revisions = []*core.Revision{}
		sr := dri.State()
		if sr != nil {
			t.Errorf(errFmt, nil, sr)
		}
	})
}

func TestDriverCloseMethodShould(t *testing.T) {
	t.Run("exist as a method", func(t *testing.T) {
		var dri Driver

		dri.Close()
	})
}
