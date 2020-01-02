package JsonArray

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		wantJa *JsonArray
	}{
		{
			name:   "Success",
			wantJa: &JsonArray{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotJa := New(); !reflect.DeepEqual(gotJa, tt.wantJa) {
				t.Errorf("New() = %v, want %v", gotJa, tt.wantJa)
			}
		})
	}
}

func TestJsonArray_Add(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		ja   JsonArray
		args args
		want *JsonArray
	}{
		{
			name: "Success",
			ja: JsonArray{
				"first",
			},
			args: args{
				value: "second",
			},
			want: &JsonArray{
				"first",
				"second",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ja.Add(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
