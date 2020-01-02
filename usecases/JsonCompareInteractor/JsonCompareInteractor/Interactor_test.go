package JsonCompareInteractor

import (
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/JsonArrayToObjectConvertInteractorInterface"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/JsonArrayToObjectConvertInteractorMock"
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

func TestInteractor_SetArrayToObjectConverter(t *testing.T) {
	type fields struct {
		arrayToObjectConverter JsonArrayToObjectConvertInteractorInterface.Interactor
	}
	type args struct {
		arrayToObjectConverter JsonArrayToObjectConvertInteractorInterface.Interactor
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractorMock.New(),
			},
			want: &Interactor{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractorMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				arrayToObjectConverter: tt.fields.arrayToObjectConverter,
			}
			if got := i.SetArrayToObjectConverter(tt.args.arrayToObjectConverter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetArrayToObjectConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}
