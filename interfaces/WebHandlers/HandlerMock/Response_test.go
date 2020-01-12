package HandlerMock

import (
	"reflect"
	"testing"
)

func TestNewResponse(t *testing.T) {
	tests := []struct {
		name  string
		wantR *Response
	}{
		{
			name:  "Success",
			wantR: &Response{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := NewResponse(); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewResponse() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestResponse_SetMyResponseVar(t *testing.T) {
	type fields struct {
		MyResponseVar interface{}
	}
	type args struct {
		myResponseVar interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				myResponseVar: "myValue",
			},
			want: &Response{
				MyResponseVar: "myValue",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				MyResponseVar: tt.fields.MyResponseVar,
			}
			if got := r.SetMyResponseVar(tt.args.myResponseVar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMyResponseVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
