package functions

import (
	"strings"
	"unicode"
)

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

// MaskRangeN
/*

 */
func MaskRangeN(s *string) string {
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
