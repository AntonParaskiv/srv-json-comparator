package JsonArrayToObjectConvertInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityInterace"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityMock"
	"reflect"
	"testing"
)

func TestInteractor_EntityArraysToObjects(t *testing.T) {
	type fields struct {
		jsonEntityIn      JsonEntityInterace.JsonEntity
		jsonEntityOut     JsonEntityInterace.JsonEntity
		simulateErrorFlag bool
	}
	type args struct {
		jsonEntityIn JsonEntityInterace.JsonEntity
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantJsonEntityOut JsonEntityInterace.JsonEntity
		wantErr           bool
		wantInteractor    *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				jsonEntityOut: JsonEntityMock.New().SetData("out"),
			},
			args: args{
				jsonEntityIn: JsonEntityMock.New().SetData("in"),
			},
			wantJsonEntityOut: JsonEntityMock.New().SetData("out"),
			wantErr:           false,
			wantInteractor: New().
				SetJsonEntityIn(JsonEntityMock.New().SetData("in")).
				SetJsonEntityOut(JsonEntityMock.New().SetData("out")),
		},
		{
			name: "Error",
			fields: fields{
				jsonEntityOut:     JsonEntityMock.New().SetData("out"),
				simulateErrorFlag: true,
			},
			args: args{
				jsonEntityIn: JsonEntityMock.New().SetData("in"),
			},
			wantJsonEntityOut: JsonEntityMock.New().SetData("out"),
			wantErr:           true,
			wantInteractor: New().
				SetJsonEntityIn(JsonEntityMock.New().SetData("in")).
				SetJsonEntityOut(JsonEntityMock.New().SetData("out")),
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
