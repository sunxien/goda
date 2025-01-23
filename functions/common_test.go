package functions

import "testing"

// TestCharLength
func TestCharLength(t *testing.T) {
	var input = "我😁1爱☺a中🇨🇳Y国"
	var expected = 11
	if actual := CharLength(&input); *actual.Value != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual.Value)
	}
}

// BenchmarkCharLength
func BenchmarkCharLength(b *testing.B) {
	var str = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < b.N; i++ {
		CharLength(&str)
	}
}
