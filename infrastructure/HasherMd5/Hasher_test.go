package HasherMd5

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantH *Hasher
	}{
		{
			name:  "Success",
			wantH: &Hasher{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotH := New(); !reflect.DeepEqual(gotH, tt.wantH) {
				t.Errorf("New() = %v, want %v", gotH, tt.wantH)
			}
		})
	}
}

func TestHasher_CreateHash(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		args     args
		wantHash string
	}{
		{
			name: "Success",
			args: args{
				data: []byte("example"),
			},
			wantHash: "1a79a4d60de6718e8e5b326e338ae533",
		},
		{
			name: "Success",
			args: args{
				data: []byte("anotherExample"),
			},
			wantHash: "5ba2e788f418208462091d80473998e9",
		},
		{
			name: "Success",
			args: args{
				data: []byte("yetAnotherExample"),
			},
			wantHash: "d2a60c3079c969cf9beb78e7c83fee9b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hasher{}
			if gotHash := h.CreateHash(tt.args.data); gotHash != tt.wantHash {
				t.Errorf("CreateHash() = %v, want %v", gotHash, tt.wantHash)
			}
		})
	}
}
