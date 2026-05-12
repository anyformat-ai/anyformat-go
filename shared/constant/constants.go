// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/anyformat-ai/anyformat-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Boolean string     // Always "boolean"
type Date string        // Always "date"
type Datetime string    // Always "datetime"
type Enum string        // Always "enum"
type Float string       // Always "float"
type Integer string     // Always "integer"
type MultiSelect string // Always "multi_select"
type Object string      // Always "object"
type String string      // Always "string"

func (c Boolean) Default() Boolean         { return "boolean" }
func (c Date) Default() Date               { return "date" }
func (c Datetime) Default() Datetime       { return "datetime" }
func (c Enum) Default() Enum               { return "enum" }
func (c Float) Default() Float             { return "float" }
func (c Integer) Default() Integer         { return "integer" }
func (c MultiSelect) Default() MultiSelect { return "multi_select" }
func (c Object) Default() Object           { return "object" }
func (c String) Default() String           { return "string" }

func (c Boolean) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Date) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Datetime) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c Enum) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Float) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Integer) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c MultiSelect) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Object) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c String) MarshalJSON() ([]byte, error)      { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
