package bindgraphql

import (
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
)

type child struct {
	StringType string `graph:"string_type"`
}

type dummy struct {
	Skip    string `graph:"-"`
	ID      int64  `graph:"ID"`
	IntType int    `graph:"int_type"`
	Child   child
}

func TestSkip(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"skip1", args{"-"}, true},
		{"skip2", args{"id"}, false},
		{"skip3", args{"payment_id"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skip(tt.args.tag); got != tt.want {
				t.Errorf("skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTag(t *testing.T) {
	type args struct {
		sf reflect.StructField
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"getTagGraph1", args{reflect.StructField{Tag: `graph:"id"`}}, "id"},
		{"getTagGraph2", args{reflect.StructField{Tag: `graph:"payment_id"`}}, "payment_id"},
		{"getTagGraph3", args{reflect.StructField{Tag: `graph:"user_id"`}}, "user_id"},

		{"getTagJSON1", args{reflect.StructField{Tag: `json:"id"`}}, "id"},
		{"getTagJSON2", args{reflect.StructField{Tag: `json:"payment_id"`}}, "payment_id"},
		{"getTagJSON3", args{reflect.StructField{Tag: `json:"user_id"`}}, "user_id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTag(tt.args.sf); got != tt.want {
				t.Errorf("getTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendFields(t *testing.T) {
	type args struct {
		dest   graphql.Fields
		source graphql.Fields
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"AppendFields1", args{graphql.Fields{}, graphql.Fields{}}, false},
		{"AppendFields2", args{graphql.Fields{}, graphql.Fields{"id": &graphql.Field{Type: graphql.ID}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := appendFields(tt.args.dest, tt.args.source); (err != nil) != tt.wantErr {
				t.Errorf("appendFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetGraphType(t *testing.T) {
	type args struct {
		tag       string
		fieldType reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want *graphql.Scalar
	}{
		{"GetGraphTypeID", args{"ID", reflect.Int}, graphql.ID},

		{"GetGraphTypeInt", args{"int", reflect.Int}, graphql.Int},
		{"GetGraphTypeInt8", args{"int8", reflect.Int8}, graphql.Int},
		{"GetGraphTypeInt16", args{"int16", reflect.Int16}, graphql.Int},
		{"GetGraphTypeInt32", args{"int32", reflect.Int32}, graphql.Int},
		{"GetGraphTypeInt64", args{"int64", reflect.Int64}, graphql.Int},

		{"GetGraphTypeFloat32", args{"float32", reflect.Float32}, graphql.Float},
		{"GetGraphTypeFloat64", args{"float64", reflect.Float64}, graphql.Float},

		{"GetGraphTypeBool", args{"bool", reflect.Bool}, graphql.Boolean},

		{"GetGraphTypeString", args{"string", reflect.String}, graphql.String},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGraphType(tt.args.tag, tt.args.fieldType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGraphType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetResolve(t *testing.T) {
	type args struct {
		fieldTag string
		obj      interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"GetResolveSkip", args{fieldTag: "-", obj: dummy{Skip: "skip"}}, nil},

		{"GetResolveID", args{fieldTag: "ID", obj: dummy{ID: int64(1)}}, int64(1)},
		{"GetResolveInt", args{fieldTag: "int_type", obj: dummy{IntType: int(1)}}, int(1)},

		{"GetResolveChildString", args{fieldTag: "string_type", obj: dummy{Child: child{StringType: "child"}}}, "child"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getResolve(tt.args.fieldTag, tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
