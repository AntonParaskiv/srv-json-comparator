package JsonCompareInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonCompareInteractor/JsonCompareInteractorInterface"
	"reflect"
	"testing"
)

func TestInteractor_IsEqual(t *testing.T) {
	type fields struct {
		first             JsonEntity.JsonEntity
		second            JsonEntity.JsonEntity
		isEqual           bool
		simulateErrorFlag bool
	}
	type args struct {
		first  JsonEntity.JsonEntity
		second JsonEntity.JsonEntity
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
				first:  JsonEntity.NewFromInterface("first"),
				second: JsonEntity.NewFromInterface("second"),
			},
			wantIsEqual: true,
			wantErr:     false,
			wantInteractor: New().
				SetFirst(JsonEntity.NewFromInterface("first")).
				SetSecond(JsonEntity.NewFromInterface("second")).
				SetIsEqual(true),
		},
		{
			name: "UnEqual",
			fields: fields{
				isEqual: false,
			},
			args: args{
				first:  JsonEntity.NewFromInterface("first"),
				second: JsonEntity.NewFromInterface("second"),
			},
			wantIsEqual: false,
			wantErr:     false,
			wantInteractor: New().
				SetFirst(JsonEntity.NewFromInterface("first")).
				SetSecond(JsonEntity.NewFromInterface("second")).
				SetIsEqual(false),
		},
		{
			name: "Error",
			fields: fields{
				isEqual:           false,
				simulateErrorFlag: true,
			},
			args: args{
				first:  JsonEntity.NewFromInterface("first"),
				second: JsonEntity.NewFromInterface("second"),
			},
			wantIsEqual: false,
			wantErr:     true,
			wantInteractor: New().
				SetFirst(JsonEntity.NewFromInterface("first")).
				SetSecond(JsonEntity.NewFromInterface("second")).
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
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsEqual != tt.wantIsEqual {
				t.Errorf("Handle() gotIsEqual = %v, want %v", gotIsEqual, tt.wantIsEqual)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}
