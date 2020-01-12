package HandlerMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/WebHandlers/HandlerInterface"
	"reflect"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	type fields struct {
		request  *Request
		response *Response
	}
	type args struct {
		requestInterface HandlerInterface.Request
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse HandlerInterface.Response
		wantHandler  *Handler
	}{
		{
			name: "Success",
			fields: fields{
				response: NewResponse().SetMyResponseVar("response"),
			},
			args: args{
				requestInterface: NewRequest().SetMyRequestVar("request"),
			},
			wantResponse: NewResponse().SetMyResponseVar("response"),
			wantHandler: &Handler{
				request:  NewRequest().SetMyRequestVar("request"),
				response: NewResponse().SetMyResponseVar("response"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				request:  tt.fields.request,
				response: tt.fields.response,
			}
			if gotResponse := h.Handle(tt.args.requestInterface); !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Handle() = %v, want %v", gotResponse, tt.wantResponse)
			}
			if !reflect.DeepEqual(h, tt.wantHandler) {
				t.Errorf("handler = %v, want %v", h, tt.wantHandler)
			}
		})
	}
}
