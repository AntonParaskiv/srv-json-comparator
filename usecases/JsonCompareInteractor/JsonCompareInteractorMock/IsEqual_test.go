package JsonCompareInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityInterace"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityMock"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/JsonCompareInteractorInterface"
	"reflect"
	"testing"
)

func TestInteractor_IsEqual(t *testing.T) {
	type fields struct {
		first             JsonEntityInterace.JsonEntity
		second            JsonEntityInterace.JsonEntity
		isEqual           bool
		simulateErrorFlag bool
	}
	type args struct {
		first  JsonEntityInterace.JsonEntity
		second JsonEntityInterace.JsonEntity
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantIsEqual    bool
		wantErr        bool
		wantInteractor JsonCompareInteractorInterface.Interactor
	}{
		{
			name: "Equal",
			fields: fields{
				isEqual: true,
			},
			args: args{
				first:  JsonEntityMock.New(),
				second: JsonEntityMock.New(),
			},
			wantIsEqual: true,
			wantErr:     false,
			wantInteractor: New().
				SetFirst(JsonEntityMock.New()).
				SetSecond(JsonEntityMock.New()).
				SetIsEqual(true),
		},
		{
			name: "UnEqual",
			fields: fields{
				isEqual: false,
			},
			args: args{
				first:  JsonEntityMock.New(),
				second: JsonEntityMock.New(),
			},
			wantIsEqual: false,
			wantErr:     false,
			wantInteractor: New().
				SetFirst(JsonEntityMock.New()).
				SetSecond(JsonEntityMock.New()).
				SetIsEqual(false),
		},
		{
			name: "Error",
			fields: fields{
				isEqual:           false,
				simulateErrorFlag: true,
			},
			args: args{
				first:  JsonEntityMock.New(),
				second: JsonEntityMock.New(),
			},
			wantIsEqual: false,
			wantErr:     true,
			wantInteractor: New().
				SetFirst(JsonEntityMock.New()).
				SetSecond(JsonEntityMock.New()).
				SetIsEqual(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				first:             tt.fields.first,
				second:            tt.fields.second,
				isEqual:           tt.fields.isEqual,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			gotIsEqual, err := i.IsEqual(tt.args.first, tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsEqual != tt.wantIsEqual {
				t.Errorf("IsEqual() gotIsEqual = %v, want %v", gotIsEqual, tt.wantIsEqual)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}
