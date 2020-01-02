package WebService

import (
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/JsonCompareInteractor"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/JsonCompareInteractorInterface"
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

func TestWebService_SetJsonCompareInteractor(t *testing.T) {
	type fields struct {
		jsonCompareInteractor JsonCompareInteractorInterface.Interactor
	}
	type args struct {
		jsonCompareInteractor JsonCompareInteractorInterface.Interactor
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *WebService
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				jsonCompareInteractor: JsonCompareInteractor.New(),
			},
			want: &WebService{
				jsonCompareInteractor: JsonCompareInteractor.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &WebService{
				jsonCompareInteractor: tt.fields.jsonCompareInteractor,
			}
			if got := ws.SetJsonCompareInteractor(tt.args.jsonCompareInteractor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetJsonCompareInteractor() = %v, want %v", got, tt.want)
			}
		})
	}
}
