package HandlerMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebHandlers/HandlerInterface"
	"reflect"
	"testing"
)

func TestHandler_NewRequest(t *testing.T) {
	type fields struct {
	}
	tests := []struct {
		name        string
		fields      fields
		wantRequest HandlerInterface.Request
	}{
		{
			name:        "Success",
			fields:      fields{},
			wantRequest: NewRequest(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{}
			if gotRequest := h.NewRequest(); !reflect.DeepEqual(gotRequest, tt.wantRequest) {
				t.Errorf("NewRequest() = %v, want %v", gotRequest, tt.wantRequest)
			}
		})
	}
}

func TestHandler_SetRequest(t *testing.T) {
	type fields struct {
		request *Request
	}
	type args struct {
		request *Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Handler
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				request: NewRequest().SetMyRequestVar("myValue"),
			},
			want: &Handler{
				request: NewRequest().SetMyRequestVar("myValue"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				request: tt.fields.request,
			}
			if got := h.SetRequest(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_SetResponse(t *testing.T) {
	type fields struct {
		response *Response
	}
	type args struct {
		response *Response
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Handler
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				response: NewResponse().SetMyResponseVar("myValue"),
			},
			want: &Handler{
				response: NewResponse().SetMyResponseVar("myValue"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				response: tt.fields.response,
			}
			if got := h.SetResponse(tt.args.response); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_GetRequest(t *testing.T) {
	type fields struct {
		request *Request
	}
	tests := []struct {
		name        string
		fields      fields
		wantRequest *Request
	}{
		{
			name: "Success",
			fields: fields{
				request: NewRequest().SetMyRequestVar("myValue"),
			},
			wantRequest: NewRequest().SetMyRequestVar("myValue"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				request: tt.fields.request,
			}
			if gotRequest := h.GetRequest(); !reflect.DeepEqual(gotRequest, tt.wantRequest) {
				t.Errorf("GetRequest() = %v, want %v", gotRequest, tt.wantRequest)
			}
		})
	}
}

func TestHandler_GetResponse(t *testing.T) {
	type fields struct {
		response *Response
	}
	tests := []struct {
		name         string
		fields       fields
		wantResponse *Response
	}{
		{
			name: "Success",
			fields: fields{
				response: NewResponse().SetMyResponseVar("myValue"),
			},
			wantResponse: NewResponse().SetMyResponseVar("myValue"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				response: tt.fields.response,
			}
			if gotResponse := h.GetResponse(); !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("GetResponse() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
