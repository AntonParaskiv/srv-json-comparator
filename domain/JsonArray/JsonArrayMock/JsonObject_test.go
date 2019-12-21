package JsonArrayMock

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		wantJa *JsonArray
	}{
		{
			name:   "Success",
			wantJa: &JsonArray{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotJa := New(); !reflect.DeepEqual(gotJa, tt.wantJa) {
				t.Errorf("New() = %v, want %v", gotJa, tt.wantJa)
			}
		})
	}
}
