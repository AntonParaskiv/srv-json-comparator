package EntityEqualRepositoryMock

import (
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

//func TestRepository_SetRepository(t *testing.T) {
//	type fields struct {
//		Repository RepositoryInterface.Repository
//	}
//	type args struct {
//		Repository RepositoryInterface.Repository
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *Repository
//	}{
//		{
//			name:   "Success",
//			fields: fields{},
//			args: args{
//				Repository: RepositoryMock.New(),
//			},
//			want: &Repository{
//				Repository: RepositoryMock.New(),
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &Repository{
//				Repository: tt.fields.Repository,
//			}
//			if got := r.SetRepository(tt.args.Repository); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("SetRepository() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestRepository_IsEqual(t *testing.T) {
//	type fields struct {
//		Repository RepositoryInterface.Repository
//	}
//	type args struct {
//		first  JsonEntity.JsonEntity
//		second JsonEntity.JsonEntity
//	}
//	tests := []struct {
//		name        string
//		fields      fields
//		args        args
//		wantIsEqual bool
//	}{
//		{
//			name: "True",
//			fields: fields{
//				Repository: RepositoryMock.New().SetIsEqual(true),
//			},
//			args: args{
//				first:  "first",
//				second: "second",
//			},
//			wantIsEqual: true,
//		},
//		{
//			name: "False",
//			fields: fields{
//				Repository: RepositoryMock.New().SetIsEqual(false),
//			},
//			args: args{
//				first:  "first",
//				second: "second",
//			},
//			wantIsEqual: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &Repository{
//				Repository: tt.fields.Repository,
//			}
//			if gotIsEqual := r.IsEqual(tt.args.first, tt.args.second); gotIsEqual != tt.wantIsEqual {
//				t.Errorf("IsEqual() = %v, want %v", gotIsEqual, tt.wantIsEqual)
//			}
//		})
//	}
//}

func TestRepository_GetFirst(t *testing.T) {
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
			e := &Repository{
				first: tt.fields.first,
			}
			if gotFirst := e.GetFirst(); !reflect.DeepEqual(gotFirst, tt.wantFirst) {
				t.Errorf("GetFirst() = %v, want %v", gotFirst, tt.wantFirst)
			}
		})
	}
}

func TestRepository_GetIsEqual(t *testing.T) {
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
			e := &Repository{
				isEqual: tt.fields.isEqual,
			}
			if gotIsEqual := e.GetIsEqual(); gotIsEqual != tt.wantIsEqual {
				t.Errorf("GetIsEqual() = %v, want %v", gotIsEqual, tt.wantIsEqual)
			}
		})
	}
}

func TestRepository_GetSecond(t *testing.T) {
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
			e := &Repository{
				second: tt.fields.second,
			}
			if gotSecond := e.GetSecond(); !reflect.DeepEqual(gotSecond, tt.wantSecond) {
				t.Errorf("GetSecond() = %v, want %v", gotSecond, tt.wantSecond)
			}
		})
	}
}

func TestRepository_SetFirst(t *testing.T) {
	type fields struct{}
	type args struct {
		first interface{}
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
				first: "first",
			},
			want: &Repository{
				first: "first",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Repository{}
			if got := e.SetFirst(tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SetIsEqual(t *testing.T) {
	type fields struct{}
	type args struct {
		isEqual bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Repository
	}{
		{
			name:   "True",
			fields: fields{},
			args: args{
				isEqual: true,
			},
			want: &Repository{
				isEqual: true,
			},
		},
		{
			name:   "False",
			fields: fields{},
			args: args{
				isEqual: false,
			},
			want: &Repository{
				isEqual: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Repository{}
			if got := e.SetIsEqual(tt.args.isEqual); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SetSecond(t *testing.T) {
	type fields struct{}
	type args struct {
		second interface{}
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
				second: "second",
			},
			want: &Repository{
				second: "second",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Repository{}
			if got := e.SetSecond(tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_IsEqual(t *testing.T) {
	type fields struct {
		isEqual bool
	}
	type args struct {
		first  interface{}
		second interface{}
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantIsEqual    bool
		wantRepository *Repository
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
			wantRepository: &Repository{
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
			wantRepository: &Repository{
				first:   "first",
				second:  "second",
				isEqual: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Repository{
				isEqual: tt.fields.isEqual,
			}
			if gotIsEqual := e.IsEqual(tt.args.first, tt.args.second); gotIsEqual != tt.wantIsEqual {
				t.Errorf("IsEqual() = %v, want %v", gotIsEqual, tt.wantIsEqual)
			}
			if !reflect.DeepEqual(e, tt.wantRepository) {
				t.Errorf("Repository = %v, want %v", e, tt.wantRepository)
			}
		})
	}
}
