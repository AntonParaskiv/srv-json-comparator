package WebService

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		wantWs *WebService
	}{
		{
			name:   "Success",
			wantWs: &WebService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWs := New(); !reflect.DeepEqual(gotWs, tt.wantWs) {
				t.Errorf("New() = %v, want %v", gotWs, tt.wantWs)
			}
		})
	}
}
