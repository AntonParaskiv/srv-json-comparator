package JsonEntityMock

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		wantJe *JsonEntity
	}{
		{
			name:   "Success",
			wantJe: &JsonEntity{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotJe := New(); !reflect.DeepEqual(gotJe, tt.wantJe) {
				t.Errorf("New() = %v, want %v", gotJe, tt.wantJe)
			}
		})
	}
}
