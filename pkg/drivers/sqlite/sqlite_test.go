package sqlite

import (
	"reflect"
	"testing"

	"github.com/PacketFire/immigrant/pkg/core"
	_ "github.com/mattn/go-sqlite3"
)

var this SqliteDriver

func TestSqliteDriver_Migrate(t *testing.T) {
	rs := core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run("migrate", func(t *testing.T) {
		this.Init("db/immigrant.db")
		this.Migrate(rs)

		if len(this.Revisions) == 1 {
			if !reflect.DeepEqual(this.Revisions[0], rs) {
				t.Fatalf("Migration results did not match, received: %s expected: %s", this.Revisions[0], rs)
			}
		} else {
			t.Fatalf("Revisions is not required length of one. Length received: %d", len(this.Revisions))
		}
	})
}

func TestSqliteDriver_RollBack(t *testing.T) {
	rs := core.Revision{
		Revision: "1-create-test-table",
		Migrate:  []string{"create table test ( `id` int(11) not null, primary key (`id`));"},
		Rollback: []string{"drop table test"},
	}

	t.Run("rollback", func(t *testing.T) {
		this.Init("db/immigrant.db")
		//		this.Migrate(rs)
		this.Rollback(rs)
		t.Log(len(this.Revisions))

		if len(this.Revisions) != 0 {
			t.Fatal("Rollback results do not meet requirements, expected length of zero.")
		}
	})
}

func TestSqliteDriver_State(t *testing.T) {
	rHead := new(core.Revision)

	t.Run("state", func(t *testing.T) {
		this.Init("db/immigrant.db")
		check, err := this.State()
		if err != nil {
			return
		}

		if !reflect.DeepEqual(check, rHead) {
			t.Fatalf("State did not return expected results, received: %s, expected: %s", check, rHead)
		}
	})
}
