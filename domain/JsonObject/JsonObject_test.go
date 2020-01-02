package JsonObject

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		wantJo *JsonObject
	}{
		{
			name:   "Success",
			wantJo: &JsonObject{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotJo := New(); !reflect.DeepEqual(gotJo, tt.wantJo) {
				t.Errorf("New() = %v, want %v", gotJo, tt.wantJo)
			}
		})
	}
}

func TestJsonObject_Add(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		jo   JsonObject
		args args
		want *JsonObject
	}{
		{
			name: "Success",
			jo:   JsonObject{},
			args: args{
				key:   "myKey",
				value: "myValue",
			},
			want: &JsonObject{
				"myKey": "myValue",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.jo.Add(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonObject_IsKeyExist(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name        string
		jo          JsonObject
		args        args
		wantIsExist bool
	}{
		{
			name: "Exist",
			jo: JsonObject{
				"myKey": "myValue",
			},
			args: args{
				key: "myKey",
			},
			wantIsExist: true,
		},
		{
			name: "Not Exist",
			jo: JsonObject{
				"anotherKey": "myValue",
			},
			args: args{
				key: "myKey",
			},
			wantIsExist: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsExist := tt.jo.IsKeyExist(tt.args.key); gotIsExist != tt.wantIsExist {
				t.Errorf("IsKeyExist() = %v, want %v", gotIsExist, tt.wantIsExist)
			}
		})
	}
}
