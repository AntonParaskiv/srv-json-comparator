package JsonArrayToObjectConvertInteractor

import (
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonArray"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonEntity"
	"github.com/AntonParaskiv/srv-json-comparator/domain/JsonObject"
	"github.com/AntonParaskiv/srv-json-comparator/infrastructure/HasherMd5"
	"github.com/AntonParaskiv/srv-json-comparator/infrastructure/JsonMarshaller"
	"github.com/AntonParaskiv/srv-json-comparator/interfaces/HashRepository/HashRepository"
	"github.com/AntonParaskiv/srv-json-comparator/usecases/JsonArrayToObjectConvertInteractor/HashRepositoryInterface"
	"reflect"
	"testing"
)

func TestInteractor_EntityArraysToObjects(t *testing.T) {
	type fields struct {
		hashRepository HashRepositoryInterface.Repository
	}
	type args struct {
		jsonEntityIn JsonEntity.JsonEntity
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantJsonEntityOut JsonEntity.JsonEntity
		wantErr           bool
	}{
		{
			name: "One field",
			args: args{
				jsonEntityIn: "myString",
			},
			wantJsonEntityOut: "myString",
			wantErr:           false,
		},
		{
			name: "Object with one field",
			args: args{
				jsonEntityIn: JsonObject.New().
					Add("first", "myString"),
			},
			wantJsonEntityOut: JsonObject.New().
				Add("first", "myString"),
			wantErr: false,
		},
		{
			name: "Object with multiple fields",
			args: args{
				jsonEntityIn: JsonObject.New().
					Add("first", "myString").
					Add("second", 2).
					Add("third", true),
			},
			wantJsonEntityOut: JsonObject.New().
				Add("first", "myString").
				Add("second", 2).
				Add("third", true),
			wantErr: false,
		},
		{
			name: "Object with Object",
			args: args{
				jsonEntityIn: JsonObject.New().
					Add("first", JsonObject.New().Add("first", "myString")),
			},
			wantJsonEntityOut: JsonObject.New().
				Add("first", JsonObject.New().Add("first", "myString")),
			wantErr: false,
		},
		{
			name: "Array with one element",
			fields: fields{
				hashRepository: HashRepository.New().
					SetHasher(HasherMd5.New()).
					SetMarshaller(JsonMarshaller.New()),
			},
			args: args{
				jsonEntityIn: JsonArray.New().
					Add("first"),
			},
			wantJsonEntityOut: JsonObject.New().
				Add("1756461953b38cc9e2278a5bcd19ce83_1", "first"),
			wantErr: false,
		},
		{
			name: "Array with multiple elements",
			fields: fields{
				hashRepository: HashRepository.New().
					SetHasher(HasherMd5.New()).
					SetMarshaller(JsonMarshaller.New()),
			},
			args: args{
				jsonEntityIn: JsonArray.New().
					Add("first").
					Add("second").
					Add("third"),
			},
			wantJsonEntityOut: JsonObject.New().
				Add("1756461953b38cc9e2278a5bcd19ce83_1", "first").
				Add("9bf395069fe7b1aaccbc7b2048d149fd_1", "second").
				Add("6da4f03ffe04a6e713a473ac86014243_1", "third"),
			wantErr: false,
		},
		{
			name: "Array with same elements",
			fields: fields{
				hashRepository: HashRepository.New().
					SetHasher(HasherMd5.New()).
					SetMarshaller(JsonMarshaller.New()),
			},
			args: args{
				jsonEntityIn: JsonArray.New().
					Add("first").
					Add("second").
					Add("first"),
			},
			wantJsonEntityOut: JsonObject.New().
				Add("1756461953b38cc9e2278a5bcd19ce83_1", "first").
				Add("9bf395069fe7b1aaccbc7b2048d149fd_1", "second").
				Add("1756461953b38cc9e2278a5bcd19ce83_2", "first"),
			wantErr: false,
		},
		{
			name: "Array with arrays",
			fields: fields{
				hashRepository: HashRepository.New().
					SetHasher(HasherMd5.New()).
					SetMarshaller(JsonMarshaller.New()),
			},
			args: args{
				jsonEntityIn: JsonArray.New().
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
			},
			wantJsonEntityOut: JsonObject.New().
				Add("9335be5f528309249bd1083ba39ccd48_1", JsonObject.New().
					Add("4a68221710541ca5d828abfcf54ba721_1", JsonObject.New().
						Add("7917815b7acca6feb6a6eb776e4092c8_1", "value111").
						Add("1265c7548d7e499aad5b4aa926acc572_1", "value112"),
					).
					Add("aab4b70bf8e73a895a7ad222d047fbde_1", JsonObject.New().
						Add("186669d0273967247213961d80e97fe3_1", "value121").
						Add("13adb9aff68aeef22e3cbbc95e6808d4_1", "value122"),
					),
				).
				Add("eaec0454f8354a4bc74e7816afcf001a_1", JsonObject.New().
					Add("ed929937b7a7697f88cb3710bd0fd2f8_1", JsonObject.New().
						Add("16691d479275125d0b270a0f8376ab04_1", "value211").
						Add("979f1ab7b701f842f144d7a073861669_1", "value212"),
					).
					Add("0eb0ae626ff26d69324fb898f61be4cb_1", JsonObject.New().
						Add("44d63cc4016e55775ff11c9c39818619_1", "value221").
						Add("885b134ed617abd3c9670269d9462d4e_1", "value222"),
					),
				),
			wantErr: false,
		},
		{
			name: "Object with array",
			fields: fields{
				hashRepository: HashRepository.New().
					SetHasher(HasherMd5.New()).
					SetMarshaller(JsonMarshaller.New()),
			},
			args: args{
				jsonEntityIn: JsonObject.New().
					Add("first", JsonArray.New().Add("second")),
			},
			wantJsonEntityOut: JsonObject.New().
				Add("first", JsonObject.New().Add("9bf395069fe7b1aaccbc7b2048d149fd_1", "second")),
			wantErr: false,
		},
		{
			name: "Array of Objects",
			fields: fields{
				hashRepository: HashRepository.New().
					SetHasher(HasherMd5.New()).
					SetMarshaller(JsonMarshaller.New()),
			},
			args: args{
				jsonEntityIn: JsonArray.New().
					Add(JsonObject.New().Add("first", "myValue")),
			},
			wantJsonEntityOut: JsonObject.New().
				Add("03c4e9df22b01e2d2fb1bdcc2b13ac41_1", JsonObject.New().Add("first", "myValue")),
			wantErr: false,
		},
		{
			name: "Object-Array multiple levels",
			fields: fields{
				hashRepository: HashRepository.New().
					SetHasher(HasherMd5.New()).
					SetMarshaller(JsonMarshaller.New()),
			},
			args: args{
				jsonEntityIn: JsonObject.New().
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
			},
			wantJsonEntityOut: JsonObject.New().
				Add("level", "first").
				Add("children",
					JsonObject.New().Add("6692e2eeefbc41be1d9f6107a65fcc8b_1",
						JsonObject.New().
							Add("level", "second").
							Add("children",
								JsonObject.New().Add("9b16867ed3e8cfa23c17bc65968c84fc_1",
									JsonObject.New().
										Add("level", "third").
										Add("children",
											JsonObject.New().
												Add("59d05c4ee0744561dcc962d11898cc9e_1", "enough"),
										),
								),
							),
					),
				),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				hashRepository: tt.fields.hashRepository,
			}
			gotJsonEntityOut, err := i.EntityArraysToObjects(tt.args.jsonEntityIn)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntityArraysToObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotJsonEntityOut, tt.wantJsonEntityOut) {
				t.Errorf("EntityArraysToObjects() gotJsonEntityOut = %v, want %v", gotJsonEntityOut, tt.wantJsonEntityOut)
			}
		})
	}
}
