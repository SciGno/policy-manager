package graphql

import (
	"context"
	"strings"

	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/location"
	"github.com/graphql-go/handler"
)

type gqlContextError struct {
	field       string
	argument    string
	code        string
	description string
	locations   *GQLErrorLocations
}

// GQLErrorLocations struct
type GQLErrorLocations struct {
	locations []location.SourceLocation
}

// Add to the locations array
func (l *GQLErrorLocations) Add(line int, column int) {
	l.locations = append(l.locations, location.SourceLocation{Line: line, Column: column})
}

// Get to the locations array
func (l *GQLErrorLocations) Get() []location.SourceLocation {
	return l.locations
}

// ContextErrors struct
type ContextErrors struct {
	gqlErrors []gqlContextError
}

// Add error to the array
func (c *ContextErrors) Add(field string, argument string, code string, line int, column int) *GQLErrorLocations {
	locations := GQLErrorLocations{}
	locations.Add(line, column)

	e := gqlContextError{
		field:       field,
		argument:    argument,
		code:        code,
		description: messageMap[code],
		locations:   &locations,
	}
	c.gqlErrors = append(c.gqlErrors, e)
	return &locations
}

// Get returns the errors array
func (c *ContextErrors) Get() gqlerrors.FormattedErrors {

	arr := gqlerrors.FormattedErrors{}

	for _, v := range c.gqlErrors {
		tmp := gqlerrors.FormattedError{
			Message:   strings.Join([]string{v.field, v.argument, v.code, v.description}, ":"),
			Locations: v.locations.Get(),
		}
		arr = append(arr, tmp)
	}

	return arr
}

// Publish error to the array
func (c *ContextErrors) Publish(ctx context.Context) {

	arr := gqlerrors.FormattedErrors{}

	for _, v := range c.gqlErrors {
		tmp := gqlerrors.FormattedError{
			Message:   strings.Join([]string{v.field, v.argument, v.code, v.description}, ":"),
			Locations: v.locations.Get(),
		}
		arr = append(arr, tmp)
	}

	if len(arr) > 0 {
		errorsFunc := ctx.Value(handler.ContextKey(handler.ErrorsContextKey)).(func(e gqlerrors.FormattedErrors))
		errorsFunc(arr)
	}
}
