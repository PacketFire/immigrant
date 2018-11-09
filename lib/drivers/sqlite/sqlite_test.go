package sqlite

import (
	"reflect"
	"testing"

	"github.com/PacketFire/immigrant/lib/core"
)

var ec chan error
var name string
var this SqliteDriver

func TestSqliteDriver_Migrate(t *testing.T) {
	name = "migrate"
	rs := &core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run(name, func(t *testing.T) {
		this.Migrate(rs, ec)

		if len(this.Revisions) == 1 {
			if !reflect.DeepEqual(this.Revisions[0], *rs) {
				t.Log("failed")
			}
		}
	})
}

func TestSqliteDriver_RollBack(t *testing.T) {
	name = "rollback"
	rs := &core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run(name, func(t *testing.T) {
		this.Rollback(rs, ec)

		if len(this.Revisions) != 0 {
			t.Log("failed")
		}
	})
}

func TestSqliteDriver_State(t *testing.T) {
	name = "state"
	rs := &core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	rs2 := &core.Revision{
		Revision: "2-create-test2-table",
		Migrate:  []string{"create table test2 ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test2"},
	}

	t.Run(name, func(t *testing.T) {
		this.Migrate(rs, ec)
		this.Migrate(rs2, ec)
		if !reflect.DeepEqual(*this.State(), this.Revisions[len(this.Revisions)-1]) {
			t.Log("failed")
		}
	})
}