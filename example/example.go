package main

import (
	"encoding/json"
	"log"

	bind "github.com/ariefrahmansyah/bindgraphql"
	"github.com/graphql-go/graphql"
)

type exampleStruct struct {
	ID          int64   `graph:"ID"`
	IntType     int     `graph:"int_type"`
	Int8Type    int8    `graph:"int8_type"`
	Int16Type   int16   `graph:"int16_type"`
	Int32Type   int32   `graph:"int32_type"`
	Int64Type   int64   `graph:"int64_type"`
	Float32Type float32 `graph:"float32_type"`
	Float64Type float64 `graph:"float64_type"`
	BoolType    bool    `graph:"bool_type"`
	StringType  string  `graph:"string_type"`
}

var example = exampleStruct{
	ID:          1,
	IntType:     1,
	Int8Type:    8,
	Int16Type:   16,
	Int32Type:   32,
	Int64Type:   64,
	Float32Type: 32,
	Float64Type: 64,
	BoolType:    true,
	StringType:  "string",
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	exampleGraphFields, err := bind.NewFields(example)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", exampleGraphFields)

	exampleObject := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Example",
		Fields: exampleGraphFields,
	})
	log.Printf("%+v\n", exampleObject)

	// You also can create new graphql.Object by using bind.NewObject()
	// exampleObject2, err := bind.NewObject("Example", example)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("%+v\n", exampleObject2)

	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"example": &graphql.Field{
				Type: exampleObject,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return example, nil
				},
			},
		},
	})
	log.Printf("%+v\n", query)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", schema)

	request := `
		{
			example {
				ID,
				int_type,
				int8_type,
				int16_type,
				int32_type,
				int64_type,
				float32_type,
				float64_type,
				bool_type,
				string_type
			}
		}
	`

	params := graphql.Params{Schema: schema, RequestString: request}
	resp := graphql.Do(params)
	if len(resp.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v\n", resp.Errors)
	}

	respJSON, _ := json.Marshal(resp)
	log.Println(string(respJSON))
}
