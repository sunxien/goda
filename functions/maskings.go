package functions

import (
	"github.com/araddon/dateparse"
	rand2 "math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Float32Floor
/*
 * Usage: Float32Floor(123456.789, 3)
 * Output: 123000
 */
func Float32Floor(f float32, pos uint) *Integer {
	iv := int(f)
	if iv == 0 {
		return NonNullInteger(0)
	}
	seed := 1
	for i := 0; i < int(pos); i++ {
		iv = iv / 10
		seed = seed * 10
	}
	return NonNullInteger(iv * seed)
}

// Float64Floor
/*
 * Usage: Float64Floor(123456.789, 3)
 * Output: 123000
 */
func Float64Floor(f float64, pos uint) *Integer {
	iv := int(f)
	if iv == 0 {
		return NonNullInteger(0)
	}
	seed := 1
	for i := 0; i < int(pos); i++ {
		iv = iv / 10
		seed = seed * 10
	}
	return NonNullInteger(iv * seed)
}

// DecimalFloor
/*
 * Usage: DecimalFloor("1234567.89", 3)
 * Output: 1234000
 */
func DecimalFloor(f *string, pos uint) *Integer {
	iv, err := strconv.ParseFloat(*f, 64)
	if err != nil {
		panic("strconv.ParseFloat(" + (*f) + ") failed. Error: " + err.Error())
	}
	if iv == 0 {
		return NonNullInteger(0)
	}
	seed := 1
	for i := 0; i < int(pos); i++ {
		iv = iv / 10
		seed = seed * 10
	}
	return NonNullInteger(int(iv) * seed)
}

// DateTimeFloor
/*
 * Usage: DateTimeFloor("2025-01-23 09:39:31", Hour)
 * Output: "2025-01-23 09:00:00"
 */
func DateTimeFloor(t *time.Time, field uint) *DateTime {
	if t == nil {
		return NullDateTime()
	}
	if field < 0 || field > 4 {
		panic("invalid datetime field. e.g: Year(0), Month(1), Hour(3), Minute(4)")
	}
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	switch field {
	case Year:
		year, month, day = 0, -month+1, -day+1
		break
	case Month:
		year, month, day = 0, 0, -day+1
		break
	case Hour:
		year, month, day, hour = 0, 0, 0, 0
		break
	case Minute:
		year, month, day, hour, minute = 0, 0, 0, 0, 0
		break
	default:
		year, month, day, hour, minute, second = 0, 0, 0, 0, 0, 0
	}
	// s string："ns", "us" (or "µs"), "ms", "s", "m", "h"
	pattern := "-" + strconv.Itoa(hour) + "h" + strconv.Itoa(minute) + "m" + strconv.Itoa(second) + "s"
	clock, err := time.ParseDuration(pattern)
	if err != nil {
		panic("time.ParseDuration" + pattern + " failed. Error: " + err.Error())
	}
	return NonNullDateTime(t.AddDate(year, int(month), day).Add(clock))
}

// DateTimeStringFloor
/*
 * Usage: DateTimeStringFloor("2025-01-23 09:39:31", Hour)
 * Output: "2025-01-23 09:00:00"
 */
func DateTimeStringFloor(s *string, field uint) *String {
	if s == nil {
		return NullString()
	}
	if *CharLength(s).Value == 0 {
		return NonNullString(*s)
	}
	if field < 0 || field > 5 {
		panic("invalid datetime field. e.g: Year(0), Month(1), Hour(3), Minute(4)")
	}
	dt, err := dateparse.ParseAny(*s)
	if err != nil {
		panic("dateparse.ParseAny(" + (*s) + ") failed. Error: " + err.Error())
	}
	datetime := DateTimeFloor(&dt, field)
	if datetime.Valid {
		return NonNullString(datetime.Value.String())
	}
	return NullString()
}

// ShiftStringPosition
/*
 * Usage: ShiftStringPosition("abcdefg", 3)
 * Output: efgabcd
 */
func ShiftStringPosition(s *string, pos int) *String {
	if s == nil {
		return NullString()
	}
	runes := []rune(*s)
	rl := len(runes)
	if rl == 0 {
		return NonNullString(EmptyString)
	}
	if pos == 0 || pos%rl == 0 {
		return NonNullString(*s)
	}
	if pos > 0 {
		// shift to right
		pos = rl - (pos % rl)
	} else {
		// shift to left
		pos = (pos * -1) % rl
	}
	runes = append(runes[pos:], runes[:pos]...)
	return NonNullString(string(runes))
}

// MaskAll
/*
 * Usage: MaskAll("ABCabc123", "*")
 * Output: ***********
 */
func MaskAll(s *string, r string) *String {
	if s == nil {
		return NullString()
	}
	var cl = CharLength(s)
	if !cl.Valid || *cl.Value == 0 {
		return NonNullString(EmptyString)
	}
	return NonNullString(strings.Repeat(r, *cl.Value))
}

// MaskChars
/*
 * Replace upper case with first argument X;
 * Replace lower case with second argument x;
 * Replace digit with third argument n;
 * Usage: MaskChars("ABCabc123", "X", "x", "n")
 * Output: XXXxxxnnn
 */
func MaskChars(s *string, ur int32, lr int32, dr int32) *String {
	if s == nil {
		return NullString()
	}
	if *CharLength(s).Value == 0 {
		return NonNullString(*s)
	}
	var sb strings.Builder
	sb.Grow(len(*s))
	for _, ch := range *s {
		if unicode.IsUpper(ch) {
			sb.WriteRune(ur) // upper case replacer
		} else if unicode.IsLower(ch) {
			sb.WriteRune(lr) // lower case replacer
		} else if unicode.IsDigit(ch) {
			sb.WriteRune(dr) // digit replacer
		} else {
			sb.WriteRune(ch)
		}
	}
	return NonNullString(sb.String())
}

// MaskRangeN
/*
 * Usage: MaskRangeN("Hello World", 2, 8, "*")
 * Output: He*******ld
 */
func MaskRangeN(s *string, from uint, to uint, r string) *String {
	if s == nil {
		return NullString()
	}
	from = from - 1
	runes := []rune(*s)
	var cl = len(runes)
	if to > uint(cl) {
		to = uint(cl)
	}
	if from > to {
		return NonNullString(*s)
	}
	var sb = &strings.Builder{}
	sb.Grow(cl)
	for i, j := 0, int(from); i < j; i = i + 1 {
		sb.WriteRune(runes[i])
	}
	for i, j := from, to; i < j; i = i + 1 {
		sb.WriteString(r)
	}
	for i, j := to, cl; i < uint(j); i = i + 1 {
		sb.WriteRune(runes[i])
	}
	return NonNullString(sb.String())
}

// MaskSubString
/*
 * Usage: MaskSubString('Hello World', "Wo", "**")
 * Output: Hello **rld
 */
func MaskSubString(s *string, sub string, r string) *String {
	if s == nil {
		return NullString()
	}
	return NonNullString(strings.ReplaceAll(*s, sub, r))
}

// MaskPrefix
/*
 *
 * Usage: MaskPrefix("username@example.com", "@", "abc")
 * Output: abcbbca@example.com
 */
func MaskPrefix(s *string, del string, r string) *String {
	if s == nil {
		return NullString()
	}
	runes := []rune(*s)
	rl := len(runes)
	if rl == 0 {
		return NonNullString(*s)
	}
	idx := strings.Index(*s, del)
	if idx == -1 {
		return NonNullString(*s)
	}
	prefix := runes[0:idx]
	Shuffle(&prefix, r)
	return NonNullString(string(append(prefix, runes[idx:]...)))
}

// Shuffle
func Shuffle(s *[]rune, r string) {
	rrunes := []rune(r)
	srl, rrl := len(*s), len(rrunes)
	for i := 0; i < srl; i++ {
		(*s)[i] = rrunes[rand2.Int()%rrl]
	}
}

// Mask
/*
 * Replace upper case with first element X
 * Replace lower case with second element x
 * Replace digit with last element n
 * Examples:
 * SQL: select mask('ABCabc123', 'X', 'x', 'n') from dual;
 * Output: XXXxxxnnn
 */
func Mask(s *string, ur int32, lr int32, dr int32) string {
	if s == nil || len(Trim(s)) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.Grow(len(*s))
	for _, ch := range *s {
		if unicode.IsUpper(ch) {
			sb.WriteRune(ur) // upper case replacer
		} else if unicode.IsLower(ch) {
			sb.WriteRune(lr) // lower case replacer
		} else if unicode.IsDigit(ch) {
			sb.WriteRune(dr) // digit replacer
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

// MaskFirstN
/*
  SQL: select mask_first_n('ABCabc123', 'X', 'x', 'n', 3);
  Output: XXXabc123
*/
func MaskFirstN(s *string, ur int32, lr int32, dr int32, n uint32) string {
	//TODO
	return ""
}

// MaskLastN
/*

 */
func MaskLastN(s *string) string {
	//TODO
	return ""
}

// MaskShowFirstN
/*
  this this
*/
func MaskShowFirstN(s *string) string {
	//TODO
	return ""
}

// MaskShowLastN
/*
	this
*/
func MaskShowLastN(s *string) string {
	//TODO
	return ""
}
