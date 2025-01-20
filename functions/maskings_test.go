package functions

import "testing"

// TestMask
/*

 */
func TestMask(t *testing.T) {
	var input = "ABCabc123"
	var expected = "XXXxxxnnn"
	if actual := Mask(&input, 'X', 'x', 'n'); actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func BenchmarkMask(b *testing.B) {

}
