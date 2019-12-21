package JsonArrayToObjectConvertInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityInterace"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity/JsonEntityMock"
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

func TestInteractor_SetJsonEntityIn(t *testing.T) {
	type fields struct {
		jsonEntityIn JsonEntityInterace.JsonEntity
	}
	type args struct {
		jsonEntityIn JsonEntityInterace.JsonEntity
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
				jsonEntityIn: JsonEntityMock.New().SetData("in"),
			},
			want: &Interactor{
				jsonEntityIn: JsonEntityMock.New().SetData("in"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				jsonEntityIn: tt.fields.jsonEntityIn,
			}
			if got := i.SetJsonEntityIn(tt.args.jsonEntityIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetJsonEntityIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetJsonEntityOut(t *testing.T) {
	type fields struct {
		jsonEntityOut JsonEntityInterace.JsonEntity
	}
	type args struct {
		jsonEntityOut JsonEntityInterace.JsonEntity
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
				jsonEntityOut: JsonEntityMock.New().SetData("out"),
			},
			want: &Interactor{
				jsonEntityOut: JsonEntityMock.New().SetData("out"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				jsonEntityOut: tt.fields.jsonEntityOut,
			}
			if got := i.SetJsonEntityOut(tt.args.jsonEntityOut); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetJsonEntityOut() = %v, want %v", got, tt.want)
			}
		})
	}
}
