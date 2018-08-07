package mock

import (
	"reflect"
	"testing"

	"github.com/PacketFire/immigrant/lib/core"
)

func TestMockDriver_Init(t *testing.T) {
	type args struct {
		config map[string]string
	}
	tests := []struct {
		name    string
		this    *MockDriver
		args    args
		wantErr bool
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.this.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("MockDriver.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestMockDriver_Migrate(t *testing.T) {
	type args struct {
		r  *core.Revision
		ec chan error
	}
	tests := []struct {
		name string
		this *MockDriver
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Migrate(tt.args.r, tt.args.ec)
		})
	}
}

func TestMockDriver_Rollback(t *testing.T) {
	type args struct {
		r  *core.Revision
		ec chan error
	}
	tests := []struct {
		name string
		this *MockDriver
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Rollback(tt.args.r, tt.args.ec)
		})
	}
}

func TestMockDriver_State(t *testing.T) {
	type args struct {
		ec chan error
	}
	tests := []struct {
		name string
		this *MockDriver
		args args
		want *core.Revision
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.State(tt.args.ec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockDriver.State() = %v, want %v", got, tt.want)
			}
		})
	}
}
