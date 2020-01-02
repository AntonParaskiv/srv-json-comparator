package JsonArrayToObjectConvertInteractorMock

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
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
		jsonEntityIn JsonEntity.JsonEntity
	}
	type args struct {
		jsonEntityIn JsonEntity.JsonEntity
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
				jsonEntityIn: JsonEntity.NewFromInterface("in"),
			},
			want: &Interactor{
				jsonEntityIn: JsonEntity.NewFromInterface("in"),
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
		jsonEntityOut JsonEntity.JsonEntity
	}
	type args struct {
		jsonEntityOut JsonEntity.JsonEntity
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
				jsonEntityOut: JsonEntity.NewFromInterface("out"),
			},
			want: &Interactor{
				jsonEntityOut: JsonEntity.NewFromInterface("out"),
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
