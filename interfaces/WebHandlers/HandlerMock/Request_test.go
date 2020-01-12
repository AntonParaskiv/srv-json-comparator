package HandlerMock

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

func TestRequest_SetMyRequestVar(t *testing.T) {
	type fields struct {
	}
	type args struct {
		myRequestVar interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				myRequestVar: "myValue",
			},
			want: &Request{
				MyRequestVar: "myValue",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{}
			if got := r.SetMyRequestVar(tt.args.myRequestVar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMyRequestVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
