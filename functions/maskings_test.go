package functions

import "testing"

import (
	"fmt"
	"testing"
)

// TestMask_nil
func TestMask_nil(t *testing.T) {
	output := MaskAll(nil, "*")
	if output.Valid {
		t.Errorf("expected: false, actual: %t", output.Valid)
	} else {
		if output.Value != nil {
			t.Errorf("expected: nil, actual: %s", *output.Value)
		}
	}
}

// TestMask_empty
func TestMask_empty(t *testing.T) {
	input := ""
	output := MaskAll(&input, "*")
	if output.Valid {
		if *output.Value != "" {
			t.Errorf("expected: [\"\"], actual: %s", *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// TestMask_blank
func TestMask_blank(t *testing.T) {
	input := " "
	output := MaskAll(&input, "*")
	if output.Valid {
		if *output.Value != "*" {
			t.Errorf("expected: [\"\"], actual: %s", *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: [true], actual: [%t]", output.Valid)
	}
}

// TestMask_string
func TestMask_string(t *testing.T) {
	input := "我a是B中1国人"
	actual := MaskAll(&input, "*")
	expected := "********"
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: [%s], actual: [%s]", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: [true], actual: [%t]", actual.Valid)
	}
}

// BenchmarkMask
func BenchmarkMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		Mask(&input, "*")
	}
}

// MaskRangeString
func TestMaskRangeN(t *testing.T) {
	input := "我a是B中1国人" // [0, 7]
	actual := MaskRangeString(&input, 2, 7, "*")
	expected := "我******人"
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// BenchmarkMaskRangeString
func BenchmarkMaskRangeN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人" // [0, 7]
		MaskRangeString(&input, 2, 7, "*")
	}
}

// TestMaskPrefix
func TestMaskPrefix(t *testing.T) {
	input := "username@gmail.com" // [0, 7]
	actual := MaskPrefix(&input, "@", "ab")
	// expected := "我******人"
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %s, actual: nil", *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// BenchmarkMaskPrefix
func BenchmarkMaskPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "username@gmail.com" // [0, 7]
		MaskPrefix(&input, "@", "abc")
	}
}

// TestFloat32Floor
func TestFloat32Floor(t *testing.T) {
	input := float32(1234567.34)
	actual := Float32Floor(input, 3)
	expected := 1234000
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %d, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %d, actual: %d", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%f], Output: [%d]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// BenchmarkFloat32Floor
// BenchmarkFloat32Floor-10    	1000000000	         0.9142 ns/op
func BenchmarkFloat32Rounding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := float32(1234567.34)
		Float32Floor(input, 3)
	}
}

// TestFloat64Floor
func TestFloat64Floor(t *testing.T) {
	input := float64(1234567.34)
	actual := Float64Floor(input, 3)
	expected := 1234000
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %d, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %d, actual: %d", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%f], Output: [%d]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// BenchmarkFloat64Floor
// BenchmarkFloat64Floor-10    	1000000000	         0.9181 ns/op
func BenchmarkFloat64Floor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := float64(1234567.34)
		Float64Floor(input, 3)
	}
}

// TestDecimalFloor
func TestDecimalFloor(t *testing.T) {
	input := "1234567.34"
	actual := DecimalFloor(&input, 3)
	expected := 1234000
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %d, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %d, actual: %d", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%d]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// BenchmarkDecimalFloor
func BenchmarkDecimalFloor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "1234567.34"
		DecimalFloor(&input, 3)
	}
}

// TestSwapStringPosition
func TestShiftStringPosition1(t *testing.T) {
	input := "abcdefg"
	actual := ShiftStringPosition(&input, 3)
	expected := "efgabcd"
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// TestSwapStringPosition
func TestShiftStringPosition2(t *testing.T) {
	input := "abcdefg"
	actual := ShiftStringPosition(&input, -3)
	expected := "defgabc"
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// BenchmarkShiftStringPosition
func BenchmarkShiftStringPosition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "abcdefg"
		ShiftStringPosition(&input, 3)
	}
}

// TestDateTimeStringFloorYear
func TestDateTimeStringFloorYear(t *testing.T) {
	input := "2025-01-23 12:17:22"
	actual := DateTimeStringFloor(&input, Year)
	expected := "2025-01-01 00:00:00 +0000 UTC"
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// TestDateTimeStringFloorMonth
func TestDateTimeStringFloorMonth(t *testing.T) {
	input := "2025-01-23 12:17:22"
	actual := DateTimeStringFloor(&input, Month)
	expected := "2024-01-01 00:00:00 +0000 UTC"
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// TestDateTimeStringFloorHour
func TestDateTimeStringFloorHour(t *testing.T) {
	input := "2025-01-23 12:17:22"
	actual := DateTimeStringFloor(&input, Hour)
	expected := "2025-01-23 12:00:00 +0000 UTC"
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// TestDateTimeStringFloorMinute
func TestDateTimeStringFloorMinute(t *testing.T) {
	input := "2025-01-23 12:17:22"
	actual := DateTimeStringFloor(&input, Minute)
	expected := "2025-01-23 12:17:00 +0000 UTC"
	if actual.Valid {
		if actual.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *actual.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *actual.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *actual.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", actual.Valid)
	}
}

// TestMask
/*

 */
func TestMaskAll(t *testing.T) {
	var input = "ABCabc123"
	var expected = "XXXxxxnnn"
	if actual := Mask(&input, 'X', 'x', 'n'); actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

// BenchmarkMaskAll
/*

 */
func BenchmarkMaskAll(b *testing.B) {
	var input = "ABCabc123"
	for i := 0; i < b.N; i++ {
		Mask(&input, 'X', 'x', 'n')
	}
}
