package graphql

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func coerceString(value interface{}) interface{} {
	if v, ok := value.(*string); ok {
		return *v
	}
	if v, ok := value.(string); ok {
		return v
	}
	return fmt.Sprintf("%v", value)
}

func coerceNaturalNumber(value interface{}) interface{} {
	if v, ok := value.(*int); ok {
		if *v >= 0 {
			return *v
		}
	}
	if v, ok := value.(int); ok {
		if v >= 0 {
			return v
		}
	}
	return nil
}

func coerceUUID(value interface{}) interface{} {
	// log.Printf("coerceUUID: %v", value)
	if uuid, err := uuid.Parse(value.(string)); err == nil {
		return uuid
	}
	return uuid.UUID{}
}

// String Scalar
var String = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ValidString",
	Description: "This string cannot be empty or be composed of only spaces. Ex: ' ' and ''  will result in an error",
	Serialize:   coerceString,
	ParseValue:  coerceString,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			if strings.TrimSpace(valueAST.Value) != "" {
				return valueAST.Value
			}
		}
		return nil
	},
})

// NaturalNumber scalar.
var NaturalNumber = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "NaturalNumber",
	Description: "The `NaturalNumber` scalar type represents a natural number value.",
	Serialize:   coerceNaturalNumber,
	ParseValue:  coerceNaturalNumber,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if intValue, err := strconv.Atoi(valueAST.Value); err == nil {
				if intValue >= 0 {
					return intValue
				}
			}
		}
		return nil
	},
})

// ValidUUID Scalar
var ValidUUID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ValidUUID",
	Description: "This string cannot be empty and must be a UUID",
	Serialize:   coerceString,
	ParseValue:  coerceString,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			if _, err := uuid.Parse(valueAST.Value); err == nil {
				if strings.TrimSpace(valueAST.Value) != "" {
					return valueAST.Value
				}
			}
		}
		return nil
	},
})

// Timestamp Scalar
var Timestamp = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Timestamp",
	Description: "A valid timestamp in the form: 2018-07-15 16:46:57.0935786 +0000 UTC",
	Serialize:   coerceString,
	ParseValue:  coerceString,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			form := "2006-01-02T15:04:05Z"
			if _, err := time.Parse(form, valueAST.Value); err == nil {
				return valueAST.Value
			}
		}
		return nil
	},
})

// // Condition Scalar
// var Condition = graphql.NewScalar(graphql.ScalarConfig{
// 	Name:        "Condition",
// 	Description: "The email will be validated for format convention",
// 	Serialize:   coerceString,
// 	ParseValue:  coerceString,
// 	ParseLiteral: func(valueAST ast.Value) interface{} {
// 		// fmt.Printf("Type: %T\n", valueAST)
// 		switch value := valueAST.(type) {
// 		case *ast.ObjectValue:
// 			return value
// 		}
// 		return nil
// 	},
// })
