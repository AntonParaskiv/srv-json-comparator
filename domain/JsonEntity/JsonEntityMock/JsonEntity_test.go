package JsonEntityMock

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		wantJe *JsonEntity
	}{
		{
			name:   "Success",
			wantJe: &JsonEntity{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotJe := New(); !reflect.DeepEqual(gotJe, tt.wantJe) {
				t.Errorf("New() = %v, want %v", gotJe, tt.wantJe)
			}
		})
	}
}

func TestJsonEntity_SetData(t *testing.T) {
	type fields struct {
		data interface{}
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *JsonEntity
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				data: "myData",
			},
			want: &JsonEntity{
				data: "myData",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			je := &JsonEntity{
				data: tt.fields.data,
			}
			if got := je.SetData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
