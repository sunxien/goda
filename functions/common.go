package functions

import (
	"strings"
	"time"
)

const EmptyString string = ""

const (
	HexEncoder    = "H" // Hex
	Base64Encoder = "B" // Base64
)

const (
	Year = iota
	Month
	Hour
	Minute
)

// String
/*
 * Avoid to use `nil` as return value.
 */
type String struct {
	Valid bool
	Value *string
}

// NullString
/*
 * Construct a `null` String
 */
func NullString() *String {
	return &String{
		Valid: false,
		Value: nil,
	}
}

// NonNullString
/*
 * Construct a `non-null` String
 */
func NonNullString(value string) *String {
	var temp = value
	return &String{
		Valid: true,
		Value: &temp,
	}
}

// Integer
/*
 * Avoid to use `nil` as return value.
 */
type Integer struct {
	Valid bool
	Value *int
}

// NullInteger
/*
 * Construct a `null` Integer.
 */
func NullInteger() *Integer {
	return &Integer{
		Valid: false,
		Value: nil,
	}
}

// NonNullInteger
/*
 * Construct a `non-null` Integer
 */
func NonNullInteger(value int) *Integer {
	var temp = value
	return &Integer{
		Valid: true,
		Value: &temp,
	}
}

// DateTime
/*
 * Avoid to use `nil` as return value.
 */
type DateTime struct {
	Valid bool
	Value *time.Time
}

// NullDateTime
/*
 * Construct a `null` DateTime.
 */
func NullDateTime() *DateTime {
	return &DateTime{
		Valid: false,
		Value: nil,
	}
}

// NonNullDateTime
/*
 * Construct a `non-null` DateTime
 */
func NonNullDateTime(value time.Time) *DateTime {
	var temp = value
	return &DateTime{
		Valid: true,
		Value: &temp,
	}
}

// ByteLength
/*
 * Usage: ByteLength("我")
 * Output: 3
 */
func ByteLength(s *string) *Integer {
	if s == nil {
		return NullInteger()
	} else if *s == "" {
		return NonNullInteger(0)
	} else {
		return NonNullInteger(len(*s))
	}
}

// CharLength
/*
 * Usage: CharLength("我")
 * Output: 1
 */
func CharLength(s *string) *Integer {
	if s == nil {
		return NullInteger()
	} else if *s == "" {
		return NonNullInteger(0)
	} else {
		return NonNullInteger(strings.Count(*s, "") - 1)
	}
}
