package JsonCompareInteractorMock

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

func TestInteractor_SetFirst(t *testing.T) {
	type fields struct {
		first JsonEntity.JsonEntity
	}
	type args struct {
		first JsonEntity.JsonEntity
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
				first: JsonEntity.NewFromInterface("first"),
			},
			want: &Interactor{
				first: JsonEntity.NewFromInterface("first"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				first: tt.fields.first,
			}
			if got := i.SetFirst(tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetJsonEntityIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetSecond(t *testing.T) {
	type fields struct {
		second JsonEntity.JsonEntity
	}
	type args struct {
		second JsonEntity.JsonEntity
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
				second: JsonEntity.NewFromInterface("second"),
			},
			want: &Interactor{
				second: JsonEntity.NewFromInterface("second"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				second: tt.fields.second,
			}
			if got := i.SetSecond(tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetJsonEntityOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetIsEqual(t *testing.T) {
	type fields struct {
		isEqual bool
	}
	type args struct {
		isEqual bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name:   "True",
			fields: fields{},
			args: args{
				isEqual: true,
			},
			want: &Interactor{
				isEqual: true,
			},
		},
		{
			name:   "False",
			fields: fields{},
			args: args{
				isEqual: false,
			},
			want: &Interactor{
				isEqual: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				isEqual: tt.fields.isEqual,
			}
			if got := i.SetIsEqual(tt.args.isEqual); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
