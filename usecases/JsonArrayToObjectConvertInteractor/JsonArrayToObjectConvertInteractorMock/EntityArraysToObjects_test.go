package JsonArrayToObjectConvertInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"reflect"
	"testing"
)

func TestInteractor_EntityArraysToObjects(t *testing.T) {
	type fields struct {
		jsonEntityIn      JsonEntity.JsonEntity
		jsonEntityOut     JsonEntity.JsonEntity
		simulateErrorFlag bool
	}
	type args struct {
		jsonEntityIn JsonEntity.JsonEntity
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantJsonEntityOut JsonEntity.JsonEntity
		wantErr           bool
		wantInteractor    *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				jsonEntityOut: JsonEntity.NewFromInterface("out"),
			},
			args: args{
				jsonEntityIn: JsonEntity.NewFromInterface("in"),
			},
			wantJsonEntityOut: JsonEntity.NewFromInterface("out"),
			wantErr:           false,
			wantInteractor: New().
				SetJsonEntityIn(JsonEntity.NewFromInterface("in")).
				SetJsonEntityOut(JsonEntity.NewFromInterface("out")),
		},
		{
			name: "Error",
			fields: fields{
				jsonEntityOut:     JsonEntity.NewFromInterface("out"),
				simulateErrorFlag: true,
			},
			args: args{
				jsonEntityIn: JsonEntity.NewFromInterface("in"),
			},
			wantJsonEntityOut: JsonEntity.NewFromInterface("out"),
			wantErr:           true,
			wantInteractor: New().
				SetJsonEntityIn(JsonEntity.NewFromInterface("in")).
				SetJsonEntityOut(JsonEntity.NewFromInterface("out")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				jsonEntityIn:      tt.fields.jsonEntityIn,
				jsonEntityOut:     tt.fields.jsonEntityOut,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			gotJsonEntityOut, err := i.EntityArraysToObjects(tt.args.jsonEntityIn)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntityArraysToObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotJsonEntityOut, tt.wantJsonEntityOut) {
				t.Errorf("EntityArraysToObjects() gotJsonEntityOut = %v, want %v", gotJsonEntityOut, tt.wantJsonEntityOut)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}
