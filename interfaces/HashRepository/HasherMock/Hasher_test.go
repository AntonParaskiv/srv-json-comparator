package HasherMock

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

func TestHasher_GetData(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name     string
		fields   fields
		wantData []byte
	}{
		{
			name: "Success",
			fields: fields{
				data: []byte("myData"),
			},
			wantData: []byte("myData"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hasher{
				data: tt.fields.data,
			}
			if gotData := h.GetData(); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetData() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestHasher_GetHash(t *testing.T) {
	type fields struct {
		hash string
	}
	tests := []struct {
		name     string
		fields   fields
		wantHash string
	}{
		{
			name: "Success",
			fields: fields{
				hash: "myHash",
			},
			wantHash: "myHash",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hasher{
				hash: tt.fields.hash,
			}
			if gotHash := h.GetHash(); gotHash != tt.wantHash {
				t.Errorf("GetHash() = %v, want %v", gotHash, tt.wantHash)
			}
		})
	}
}

func TestHasher_SetData(t *testing.T) {
	type fields struct{}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hasher
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				data: []byte("myData"),
			},
			want: &Hasher{
				data: []byte("myData"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hasher{}
			if got := h.SetData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_SetHash(t *testing.T) {
	type fields struct{}
	type args struct {
		hash string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hasher
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				hash: "myHash",
			},
			want: &Hasher{
				hash: "myHash",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hasher{}
			if got := h.SetHash(tt.args.hash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_CreateHash(t *testing.T) {
	type fields struct {
		hash string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantHash   string
		wantHasher *Hasher
	}{
		{
			name: "Success",
			fields: fields{
				hash: "myHash",
			},
			args: args{
				data: []byte("myData"),
			},
			wantHash: "myHash",
			wantHasher: &Hasher{
				data: []byte("myData"),
				hash: "myHash",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hasher{
				hash: tt.fields.hash,
			}
			if gotHash := h.CreateHash(tt.args.data); gotHash != tt.wantHash {
				t.Errorf("CreateHash() = %v, want %v", gotHash, tt.wantHash)
			}
			if !reflect.DeepEqual(h, tt.wantHasher) {
				t.Errorf("Hasher = %v, want %v", h, tt.wantHasher)
			}
		})
	}
}
