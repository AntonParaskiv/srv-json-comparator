package JsonArrayToObjectConvertInteractor

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantI *Interactor
	}{
		{
			name:  "Success",
			wantI: &Interactor{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := New(); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("New() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
