package JsonObjectMock

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		wantJo *JsonObject
	}{
		{
			name:   "Success",
			wantJo: &JsonObject{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotJo := New(); !reflect.DeepEqual(gotJo, tt.wantJo) {
				t.Errorf("New() = %v, want %v", gotJo, tt.wantJo)
			}
		})
	}
}
