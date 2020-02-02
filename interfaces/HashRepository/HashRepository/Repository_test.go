package HashRepository

import (
	"errors"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/HasherInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/HasherMock"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/MarshallerInterface"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/MarshallerMock"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantR *Repository
	}{
		{
			name:  "Success",
			wantR: &Repository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := New(); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("New() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestRepository_SetMarshaller(t *testing.T) {
	type fields struct {
		marshaller MarshallerInterface.Marshaller
		hasher     HasherInterface.Hasher
	}
	type args struct {
		marshaller MarshallerInterface.Marshaller
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Repository
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				marshaller: MarshallerMock.New(),
			},
			want: &Repository{
				marshaller: MarshallerMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				marshaller: tt.fields.marshaller,
				hasher:     tt.fields.hasher,
			}
			if got := r.SetMarshaller(tt.args.marshaller); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMarshaller() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SetHasher(t *testing.T) {
	type fields struct{}
	type args struct {
		hasher HasherInterface.Hasher
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Repository
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				hasher: HasherMock.New(),
			},
			want: &Repository{
				hasher: HasherMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{}
			if got := r.SetHasher(tt.args.hasher); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHasher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorMarshalEntityFailed(t *testing.T) {
	type args struct {
		wrapErr error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				wrapErr: errors.New("myError"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errorMarshalEntityFailed(tt.args.wrapErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("errorMarshalEntityFailed() error = %v, wantErr %v", err, tt.wantErr)
			}
			wrapErr := errors.Unwrap(err)
			if !reflect.DeepEqual(wrapErr, tt.args.wrapErr) {
				t.Errorf("wrapErr = %v, want %v", wrapErr, tt.args.wrapErr)
			}
		})
	}
}

func Test_wrapError(t *testing.T) {
	type args struct {
		msg     string
		wrapErr error
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantMessage string
	}{
		{
			name: "No Wrap",
			args: args{
				msg:     "myMessage",
				wrapErr: nil,
			},
			wantErr:     true,
			wantMessage: "myMessage",
		},
		{
			name: "Wrapped",
			args: args{
				msg:     "myMessage",
				wrapErr: errors.New("myError"),
			},
			wantErr:     true,
			wantMessage: "myMessage: myError",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := wrapError(tt.args.msg, tt.args.wrapErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("wrapError() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err.Error() != tt.wantMessage {
				t.Errorf("message = %v, want %v", err.Error(), tt.wantMessage)
			}
			wrapErr := errors.Unwrap(err)
			if !reflect.DeepEqual(wrapErr, tt.args.wrapErr) {
				t.Errorf("wrapErr = %v, want %v", wrapErr, tt.args.wrapErr)
			}
		})
	}
}
