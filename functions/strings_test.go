package functions

import "testing"

// TestReverse
/**
 *
 */
func TestReverse(t *testing.T) {
	var input = "我1爱a中Y国"
	var expected = "国Y中a爱1我"
	if actual := Reverse(&input); actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

// BenchmarkReverse
/*
 * - strings.Builder:
 * - BenchmarkReverse-12    	 1447245	       788.6 ns/op
 * - BenchmarkReverse-12    	 1508229	       795.8 ns/op
 * - BenchmarkReverse-12    	 1507929	       786.1 ns/op
 */
func BenchmarkReverse(b *testing.B) {
	var str = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < b.N; i++ {
		Reverse(&str)
	}
}

/*
 *
 */
func TestCharLength(t *testing.T) {
	var input = "我😁1爱☺a中🇨🇳Y国"
	var expected = 11
	if actual := CharLength(&input); actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

/*
 *
 */
func BenchmarkCharLength(b *testing.B) {
	var str = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < b.N; i++ {
		CharLength(&str)
	}
}
