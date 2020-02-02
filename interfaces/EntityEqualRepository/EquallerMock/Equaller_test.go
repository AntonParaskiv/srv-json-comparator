package EquallerMock

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantE *Equaller
	}{
		{
			name:  "Success",
			wantE: &Equaller{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotE := New(); !reflect.DeepEqual(gotE, tt.wantE) {
				t.Errorf("New() = %v, want %v", gotE, tt.wantE)
			}
		})
	}
}

func TestEqualler_GetFirst(t *testing.T) {
	type fields struct {
		first interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		wantFirst interface{}
	}{
		{
			name: "Success",
			fields: fields{
				first: "first",
			},
			wantFirst: "first",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Equaller{
				first: tt.fields.first,
			}
			if gotFirst := e.GetFirst(); !reflect.DeepEqual(gotFirst, tt.wantFirst) {
				t.Errorf("GetFirst() = %v, want %v", gotFirst, tt.wantFirst)
			}
		})
	}
}

func TestEqualler_GetIsEqual(t *testing.T) {
	type fields struct {
		isEqual bool
	}
	tests := []struct {
		name        string
		fields      fields
		wantIsEqual bool
	}{
		{
			name: "True",
			fields: fields{
				isEqual: true,
			},
			wantIsEqual: true,
		},
		{
			name: "False",
			fields: fields{
				isEqual: false,
			},
			wantIsEqual: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Equaller{
				isEqual: tt.fields.isEqual,
			}
			if gotIsEqual := e.GetIsEqual(); gotIsEqual != tt.wantIsEqual {
				t.Errorf("GetIsEqual() = %v, want %v", gotIsEqual, tt.wantIsEqual)
			}
		})
	}
}

func TestEqualler_GetSecond(t *testing.T) {
	type fields struct {
		second interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		wantSecond interface{}
	}{
		{
			name: "Success",
			fields: fields{
				second: "second",
			},
			wantSecond: "second",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Equaller{
				second: tt.fields.second,
			}
			if gotSecond := e.GetSecond(); !reflect.DeepEqual(gotSecond, tt.wantSecond) {
				t.Errorf("GetSecond() = %v, want %v", gotSecond, tt.wantSecond)
			}
		})
	}
}

func TestEqualler_SetFirst(t *testing.T) {
	type fields struct{}
	type args struct {
		first interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Equaller
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				first: "first",
			},
			want: &Equaller{
				first: "first",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Equaller{}
			if got := e.SetFirst(tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualler_SetIsEqual(t *testing.T) {
	type fields struct{}
	type args struct {
		isEqual bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Equaller
	}{
		{
			name:   "True",
			fields: fields{},
			args: args{
				isEqual: true,
			},
			want: &Equaller{
				isEqual: true,
			},
		},
		{
			name:   "False",
			fields: fields{},
			args: args{
				isEqual: false,
			},
			want: &Equaller{
				isEqual: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Equaller{}
			if got := e.SetIsEqual(tt.args.isEqual); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualler_SetSecond(t *testing.T) {
	type fields struct{}
	type args struct {
		second interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Equaller
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				second: "second",
			},
			want: &Equaller{
				second: "second",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Equaller{}
			if got := e.SetSecond(tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualler_IsEqual(t *testing.T) {
	type fields struct {
		isEqual bool
	}
	type args struct {
		first  interface{}
		second interface{}
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantIsEqual  bool
		wantEqualler *Equaller
	}{
		{
			name: "True",
			fields: fields{
				isEqual: true,
			},
			args: args{
				first:  "first",
				second: "second",
			},
			wantIsEqual: true,
			wantEqualler: &Equaller{
				first:   "first",
				second:  "second",
				isEqual: true,
			},
		},
		{
			name: "False",
			fields: fields{
				isEqual: false,
			},
			args: args{
				first:  "first",
				second: "second",
			},
			wantIsEqual: false,
			wantEqualler: &Equaller{
				first:   "first",
				second:  "second",
				isEqual: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Equaller{
				isEqual: tt.fields.isEqual,
			}
			if gotIsEqual := e.IsEqual(tt.args.first, tt.args.second); gotIsEqual != tt.wantIsEqual {
				t.Errorf("IsEqual() = %v, want %v", gotIsEqual, tt.wantIsEqual)
			}
			if !reflect.DeepEqual(e, tt.wantEqualler) {
				t.Errorf("equaller = %v, want %v", e, tt.wantEqualler)
			}
		})
	}
}
