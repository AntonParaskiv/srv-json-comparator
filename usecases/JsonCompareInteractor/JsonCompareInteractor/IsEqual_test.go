package JsonCompareInteractor

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonArray"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonObject"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/JsonArrayToObjectConvertInteractor"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/JsonArrayToObjectConvertInteractorInterface"
	"testing"
)

func TestInteractor_IsEqual(t *testing.T) {
	type fields struct {
		arrayToObjectConverter JsonArrayToObjectConvertInteractorInterface.Interactor
	}
	type args struct {
		first  JsonEntity.JsonEntity
		second JsonEntity.JsonEntity
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantIsEqual bool
		wantErr     bool
	}{
		{
			name: "Simple field Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first:  "myValue",
				second: "myValue",
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Simple field UnEqual",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first:  "myValue",
				second: "anotherValue",
			},
			wantIsEqual: false,
			wantErr:     false,
		},
		{
			name: "Object with one field Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonObject.New().
					Add("first", "myString"),
				second: JsonObject.New().
					Add("first", "myString"),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Object with one field UnEqual",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonObject.New().
					Add("first", "myString"),
				second: JsonObject.New().
					Add("second", "myString"),
			},
			wantIsEqual: false,
			wantErr:     false,
		},
		{
			name: "Object with multiple fields Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonObject.New().
					Add("first", "myFirstString").
					Add("second", "mySecondString"),
				second: JsonObject.New().
					Add("second", "mySecondString").
					Add("first", "myFirstString"),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Object with multiple fields UnEqual",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonObject.New().
					Add("first", "myFirstString").
					Add("second", "mySecondString"),
				second: JsonObject.New().
					Add("second", "anotherSecondString").
					Add("first", "anotherFirstString"),
			},
			wantIsEqual: false,
			wantErr:     false,
		},
		{
			name: "Object with object Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonObject.New().
					Add("first", JsonObject.New().Add("first", "myString")),
				second: JsonObject.New().
					Add("first", JsonObject.New().Add("first", "myString")),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Object with object UnEqual",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonObject.New().
					Add("first", JsonObject.New().Add("first", "myString")),
				second: JsonObject.New().
					Add("first", JsonObject.New().Add("first", "anotherString")),
			},
			wantIsEqual: false,
			wantErr:     false,
		},
		{
			name: "Array with one element Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonArray.New().
					Add("first"),
				second: JsonArray.New().
					Add("first"),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Array with one element UnEqual",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonArray.New().
					Add("first"),
				second: JsonArray.New().
					Add("second"),
			},
			wantIsEqual: false,
			wantErr:     false,
		},
		{
			name: "Array with multiple elements Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonArray.New().
					Add("first").
					Add("second").
					Add("third"),
				second: JsonArray.New().
					Add("first").
					Add("second").
					Add("third"),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Array with multiple shuffled elements Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonArray.New().
					Add("first").
					Add("second").
					Add("third"),
				second: JsonArray.New().
					Add("third").
					Add("first").
					Add("second"),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Array with multiple elements UnEqual",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonArray.New().
					Add("first").
					Add("second").
					Add("third"),
				second: JsonArray.New().
					Add("first").
					Add("second").
					Add("3"),
			},
			wantIsEqual: false,
			wantErr:     false,
		},
		{
			name: "Array with arrays Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonArray.New().
					Add(JsonArray.New().
						Add(JsonArray.New().
							Add("value111").
							Add("value112"),
						).
						Add(JsonArray.New().
							Add("value121").
							Add("value122"),
						),
					).
					Add(JsonArray.New().
						Add(JsonArray.New().
							Add("value211").
							Add("value212"),
						).
						Add(JsonArray.New().
							Add("value221").
							Add("value222"),
						),
					),
				second: JsonArray.New().
					Add(JsonArray.New().
						Add(JsonArray.New().
							Add("value212").
							Add("value211"),
						).
						Add(JsonArray.New().
							Add("value222").
							Add("value221"),
						),
					).
					Add(JsonArray.New().
						Add(JsonArray.New().
							Add("value112").
							Add("value111"),
						).
						Add(JsonArray.New().
							Add("value122").
							Add("value121"),
						),
					),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Object with array Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonObject.New().
					Add("first", JsonArray.New().Add("second")),
				second: JsonObject.New().
					Add("first", JsonArray.New().Add("second")),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Array of Objects Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonArray.New().
					Add(JsonObject.New().Add("first", "myValue")),
				second: JsonArray.New().
					Add(JsonObject.New().Add("first", "myValue")),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		{
			name: "Object-Array multiple levels Equal",
			fields: fields{
				arrayToObjectConverter: JsonArrayToObjectConvertInteractor.New(),
			},
			args: args{
				first: JsonObject.New().
					Add("level", "first").
					Add("children",
						JsonArray.New().Add(
							JsonObject.New().
								Add("level", "second").
								Add("children",
									JsonArray.New().Add(JsonObject.New().
										Add("level", "third").
										Add("children",
											JsonArray.New().Add("enough"),
										),
									),
								),
						),
					),
				second: JsonObject.New().
					Add("children",
						JsonArray.New().Add(
							JsonObject.New().
								Add("children",
									JsonArray.New().Add(JsonObject.New().
										Add("children",
											JsonArray.New().Add("enough"),
										).
										Add("level", "third"),
									),
								).
								Add("level", "second"),
						),
					).
					Add("level", "first"),
			},
			wantIsEqual: true,
			wantErr:     false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				arrayToObjectConverter: tt.fields.arrayToObjectConverter,
			}
			gotIsEqual, err := i.IsEqual(tt.args.first, tt.args.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsEqual != tt.wantIsEqual {
				t.Errorf("IsEqual() gotIsEqual = %v, want %v", gotIsEqual, tt.wantIsEqual)
			}
		})
	}
}
