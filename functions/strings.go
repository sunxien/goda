package functions

import (
	"strings"
)

// Reverse
/*

 */
func Reverse(s *string) string {
	if s == nil || len(Trim(s)) == 0 {
		return ""
	}
	runes := []rune(*s)
	i, j := 0, len(runes)-1
	for ; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Trim
/*

 */
func Trim(s *string) string {
	if s == nil {
		return ""
	}
	return strings.TrimSpace(*s)
}

// Ltrim
/*

 */
func Ltrim(s *string) string {
	if s == nil {
		return ""
	}
	return strings.TrimLeft(*s, " ")
}

// Rtrim
/*

 */
func Rtrim(s *string) string {
	if s == nil {
		return ""
	}
	return strings.TrimRight(*s, " ")
}

// Concat
/*

 */
func Concat(s1 *string, s2 *string) string {
	if s1 == nil || len(*s1) == 0 {
		return *s2
	}
	if s2 == nil || len(*s2) == 0 {
		return *s1
	}
	var sb strings.Builder
	sb.Grow(len(*s1) + len(*s2))
	sb.WriteString(*s1)
	sb.WriteString(*s2)
	return sb.String()
}

// Upper
/**

 */
func Upper(s *string) string {
	if s == nil {
		return ""
	}
	return strings.ToUpper(*s)
}

// Lower
/*

 */
func Lower(s *string) string {
	if s == nil {
		return ""
	}
	return strings.ToLower(*s)
}

// Replace
/*

 */
func Replace(s *string, search *string, replacement *string) string {
	if s == nil || len(Trim(s)) == 0 {
		return ""
	}
	if search == nil || len(Trim(search)) == 0 {
		return ""
	}
	if replacement == nil || len(Trim(replacement)) == 0 {
		return ""
	}
	return strings.ReplaceAll(*s, *search, *replacement)
}

// Lpad
/*

 */
func Lpad(s *string) string {
	// var cl = CharLength(s)
	// TODO
	return ""
}

// Rpad
/*

 */
func Rpad(s *string) string {
	// TODO
	return ""
}

// RawToHex
/*

 */
func RawToHex() {
	// TODO
}

// HexToRaw
/*

 */
func HexToRaw() {
	// TODO
}

// Ascii
/*
  See as:
  e.g: select ascii('a') from dual;
*/
func Ascii(s *string) int {
	if s == nil {
		return -1
	} else if *s == "" {
		return 0
	} else {
		return int((*s)[0])
	}
}

// Substr
/*

 */
func Substr(s *string, from int, length int) string {
	if length < 1 {
		return ""
	} else {
		if from == 0 || from > CharLength(s) {
			return ""
		} else if from < 0 {
			// TODO
			return ""
		} else {
			return (*s)[from:length]
		}
	}
}

// Length
/*
  SQL: select char_length('我') from dual;
  Output: 3
*/
func Length(s *string) int {
	if s == nil {
		return -1
	} else if *s == "" {
		return 0
	} else {
		return len(*s)
	}
}

// CharLength
/*
  SQL: select char_length('我') from dual;
  Output: 1
*/
func CharLength(s *string) int {
	if s == nil {
		return -1
	} else if *s == "" {
		return 0
	} else {
		return strings.Count(*s, "") - 1
	}
}
