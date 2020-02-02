package MarshallerMock

import (
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantM *Marshaller
	}{
		{
			name:  "Success",
			wantM: &Marshaller{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := New(); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("New() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestMarshaller_GetData(t *testing.T) {
	type fields struct {
		data interface{}
	}
	tests := []struct {
		name     string
		fields   fields
		wantData interface{}
	}{
		{
			name: "Success",
			fields: fields{
				data: "myData",
			},
			wantData: "myData",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{
				data: tt.fields.data,
			}
			if gotData := m.GetData(); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetData() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestMarshaller_GetMarshalledData(t *testing.T) {
	type fields struct {
		marshalledData []byte
	}
	tests := []struct {
		name               string
		fields             fields
		wantMarshalledData []byte
	}{
		{
			name: "Success",
			fields: fields{
				marshalledData: []byte("myData"),
			},
			wantMarshalledData: []byte("myData"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{
				marshalledData: tt.fields.marshalledData,
			}
			if gotMarshalledData := m.GetMarshalledData(); !reflect.DeepEqual(gotMarshalledData, tt.wantMarshalledData) {
				t.Errorf("GetMarshalledData() = %v, want %v", gotMarshalledData, tt.wantMarshalledData)
			}
		})
	}
}

func TestMarshaller_GetReader(t *testing.T) {
	type fields struct {
		reader io.ReadCloser
	}
	tests := []struct {
		name       string
		fields     fields
		wantReader io.ReadCloser
	}{
		{
			name: "Success",
			fields: fields{
				reader: ioutil.NopCloser(strings.NewReader("myData")),
			},
			wantReader: ioutil.NopCloser(strings.NewReader("myData")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{
				reader: tt.fields.reader,
			}
			if gotReader := m.GetReader(); !reflect.DeepEqual(gotReader, tt.wantReader) {
				t.Errorf("GetReader() = %v, want %v", gotReader, tt.wantReader)
			}
		})
	}
}

func TestMarshaller_GetWriter(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	tests := []struct {
		name       string
		fields     fields
		wantWriter io.Writer
	}{
		{
			name: "Success",
			fields: fields{
				writer: bytes.NewBufferString("myData"),
			},
			wantWriter: bytes.NewBufferString("myData"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{
				writer: tt.fields.writer,
			}
			if gotWriter := m.GetWriter(); !reflect.DeepEqual(gotWriter, tt.wantWriter) {
				t.Errorf("GetWriter() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func TestMarshaller_SetData(t *testing.T) {
	type fields struct{}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Marshaller
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				data: "myData",
			},
			want: &Marshaller{
				data: "myData",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{}
			if got := m.SetData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshaller_SetMarshalledData(t *testing.T) {
	type fields struct{}
	type args struct {
		marshalledData []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Marshaller
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				marshalledData: []byte("myData"),
			},
			want: &Marshaller{
				marshalledData: []byte("myData"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{}
			if got := m.SetMarshalledData(tt.args.marshalledData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMarshalledData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshaller_SetReader(t *testing.T) {
	type fields struct {
	}
	type args struct {
		reader io.ReadCloser
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Marshaller
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				reader: ioutil.NopCloser(strings.NewReader("myData")),
			},
			want: &Marshaller{
				reader: ioutil.NopCloser(strings.NewReader("myData")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{}
			if got := m.SetReader(tt.args.reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshaller_SetWriter(t *testing.T) {
	type fields struct {
		writer                 io.Writer
		reader                 io.ReadCloser
		data                   interface{}
		marshalledData         []byte
		simulateErrorStepMatch int
		simulateErrorFlag      bool
	}
	type args struct {
		writer io.Writer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Marshaller
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				writer: bytes.NewBufferString("myData"),
			},
			want: &Marshaller{
				writer: bytes.NewBufferString("myData"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{}
			if got := m.SetWriter(tt.args.writer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshaller_MarshalWriter(t *testing.T) {
	type fields struct {
		simulateErrorFlag bool
	}
	type args struct {
		data   interface{}
		writer io.Writer
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantMarshaller *Marshaller
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				data:   "myData",
				writer: bytes.NewBufferString("myData"),
			},
			wantErr: false,
			wantMarshaller: &Marshaller{
				data:   "myData",
				writer: bytes.NewBufferString("myData"),
			},
		},
		{
			name: "Error",
			fields: fields{
				simulateErrorFlag: true,
			},
			args: args{
				data:   "myData",
				writer: bytes.NewBufferString("myData"),
			},
			wantErr: true,
			wantMarshaller: &Marshaller{
				data:   "myData",
				writer: bytes.NewBufferString("myData"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := m.MarshalWriter(tt.args.writer, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("MarshalWriter() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.wantMarshaller) {
				t.Errorf("Marshaller = %v, want %v", m, tt.wantMarshaller)
			}
		})
	}
}

func TestMarshaller_UnMarshalReader(t *testing.T) {
	type fields struct {
		simulateErrorFlag bool
	}
	type args struct {
		reader io.ReadCloser
		data   interface{}
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantMarshaller *Marshaller
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				reader: ioutil.NopCloser(strings.NewReader("myData")),
				data:   "myData",
			},
			wantErr: false,
			wantMarshaller: &Marshaller{
				reader: ioutil.NopCloser(strings.NewReader("myData")),
				data:   "myData",
			},
		},
		{
			name: "Error",
			fields: fields{
				simulateErrorFlag: true,
			},
			args: args{
				reader: ioutil.NopCloser(strings.NewReader("myData")),
				data:   "myData",
			},
			wantErr: true,
			wantMarshaller: &Marshaller{
				reader: ioutil.NopCloser(strings.NewReader("myData")),
				data:   "myData",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := m.UnMarshalReader(tt.args.reader, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnMarshalReader() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.wantMarshaller) {
				t.Errorf("Marshaller = %v, want %v", m, tt.wantMarshaller)
			}
		})
	}
}

func TestMarshaller_MarshalBytes(t *testing.T) {
	type fields struct {
		marshalledData    []byte
		simulateErrorFlag bool
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantMarshalledData []byte
		wantErr            bool
		wantMarshaller     *Marshaller
	}{
		{
			name: "Success",
			fields: fields{
				marshalledData: []byte("myMarshalledData"),
			},
			args: args{
				data: "myData",
			},
			wantMarshalledData: []byte("myMarshalledData"),
			wantErr:            false,
			wantMarshaller: &Marshaller{
				marshalledData: []byte("myMarshalledData"),
				data:           "myData",
			},
		},
		{
			name: "Error",
			fields: fields{
				marshalledData:    []byte("myMarshalledData"),
				simulateErrorFlag: true,
			},
			args: args{
				data: "myData",
			},
			wantMarshalledData: []byte("myMarshalledData"),
			wantErr:            true,
			wantMarshaller: &Marshaller{
				marshalledData: []byte("myMarshalledData"),
				data:           "myData",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{
				marshalledData:    tt.fields.marshalledData,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			gotMarshalledData, err := m.MarshalBytes(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMarshalledData, tt.wantMarshalledData) {
				t.Errorf("MarshalBytes() gotMarshalledData = %v, want %v", gotMarshalledData, tt.wantMarshalledData)
			}
			if !reflect.DeepEqual(m, tt.wantMarshaller) {
				t.Errorf("Marshaller = %v, want %v", m, tt.wantMarshaller)
			}
		})
	}
}

func TestMarshaller_UnMarshalBytes(t *testing.T) {
	type fields struct {
		simulateErrorFlag bool
	}
	type args struct {
		marshalledData []byte
		data           interface{}
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantMarshaller *Marshaller
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				marshalledData: []byte("myMarshalledData"),
				data:           "myData",
			},
			wantErr: false,
			wantMarshaller: &Marshaller{
				marshalledData: []byte("myMarshalledData"),
				data:           "myData",
			},
		},
		{
			name: "Error",
			fields: fields{
				simulateErrorFlag: true,
			},
			args: args{
				marshalledData: []byte("myMarshalledData"),
				data:           "myData",
			},
			wantErr: true,
			wantMarshaller: &Marshaller{
				marshalledData: []byte("myMarshalledData"),
				data:           "myData",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := m.UnMarshalBytes(tt.args.marshalledData, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnMarshalBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.wantMarshaller) {
				t.Errorf("Marshaller = %v, want %v", m, tt.wantMarshaller)
			}
		})
	}
}
