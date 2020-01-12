package BaseHandler

import (
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name  string
		wantR *Request
	}{
		{
			name:  "Success",
			wantR: new(Request),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := NewRequest(); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewRequest() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
