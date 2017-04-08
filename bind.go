package bindgraphql

import (
	"errors"
	"reflect"

	"github.com/graphql-go/graphql"
)

// NewObject returns new *graphql.Object instance.
// exampleObject2, err := bind.NewObject("Example", example)
// if err != nil {
// 	log.Println(err)
// }
// log.Printf("%+v\n", exampleObject2)
func NewObject(name string, obj interface{}) (*graphql.Object, error) {
	fields, err := NewFields(obj)
	if err != nil {
		return &graphql.Object{}, err
	}

	graphObj := graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	})

	return graphObj, nil
}

// NewFields returns new graphql.Fields that ready
// to be used by your graphql.Object.
func NewFields(obj interface{}) (graphql.Fields, error) {
	graphFields := graphql.Fields{}

	val := reflect.ValueOf(obj)

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := getTag(field)

		if skip(tag) {
			continue
		}

		if _, ok := graphFields[tag]; ok && tag != "ID" {
			return graphql.Fields{}, errors.New("duplicate tag of " + tag)
		}

		if tag == "" {
			if field.Type.Kind() == reflect.Struct {
				structFields, err := NewFields(val.Field(i).Interface())
				if err != nil {
					return graphql.Fields{}, err
				}

				err = appendFields(graphFields, structFields)
				if err != nil {
					return graphql.Fields{}, err
				}
			}

			continue
		}

		graphFields[tag] = &graphql.Field{
			Type: getGraphType(tag, field.Type.Kind()),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return getResolve(tag, p.Source), nil
			},
		}

	}

	return graphFields, nil
}

func skip(tag string) bool {
	return tag == "-"
}

func getTag(sf reflect.StructField) string {
	tag := sf.Tag.Get("graph")

	if tag == "" {
		tag = sf.Tag.Get("json")
	}

	return tag
}

func appendFields(dest, source graphql.Fields) error {
	for k, v := range source {
		dest[k] = v
	}

	return nil
}

func getGraphType(tag string, fieldKind reflect.Kind) *graphql.Scalar {
	if tag == "ID" {
		return graphql.ID
	}

	switch fieldKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return graphql.Int
	case reflect.Float32, reflect.Float64:
		return graphql.Float
	case reflect.Bool:
		return graphql.Boolean
	}

	return graphql.String
}

func getResolve(fieldTag string, obj interface{}) interface{} {
	val := reflect.ValueOf(obj)

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := getTag(field)

		if skip(tag) {
			continue
		}

		if tag == fieldTag {
			return val.Field(i).Interface()
		}

		if field.Type.Kind() == reflect.Struct {
			if res := getResolve(fieldTag, val.Field(i).Interface()); res != nil {
				return res
			}
		}
	}

	return nil
}
