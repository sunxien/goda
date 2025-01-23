package functions

import (
	"fmt"
	"testing"
)

// TestMD5
func TestMD5(t *testing.T) {
	input := "我a是B中1国人"
	output := MD5(&input)
	expected := "41b56d64d73ed914dcd1c81c47ee95b1"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkMD5
func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		MD5(&input)
	}
}

// TestHMAC
func TestHMAC(t *testing.T) {
	input := "我a是B中1国人"
	output := HMAC(&input, "123456789")
	expected := "f2a7815dbf43233f2f0eb10171e8a82440fd175be3ac84439d4f3a278b4c8ee5"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkHMAC
func BenchmarkHMAC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		HMAC(&input, "123456789")
	}
}

// TestHEX
func TestHEX(t *testing.T) {
	input := "我a是B中1国人"
	output := HEX(&input)
	expected := "e6889161e698af42e4b8ad31e59bbde4baba"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkHEX
func BenchmarkHEX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		HEX(&input)
	}
}

// TestUNHEX
func TestUNHEX(t *testing.T) {
	input := "e6889161e698af42e4b8ad31e59bbde4baba"
	output := UNHEX(&input)
	expected := "我a是B中1国人"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkUNHEX
func BenchmarkUNHEX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "e6889161e698af42e4b8ad31e59bbde4baba"
		UNHEX(&input)
	}
}

// TestBASE64Encode
func TestBASE64Encode(t *testing.T) {
	input := "我a是B中1国人"
	output := BASE64Encode(&input)
	expected := "5oiRYeaYr0LkuK0x5Zu95Lq6"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkBase64Encode
func BenchmarkBase64Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		BASE64Encode(&input)
	}
}

// TestBase64Decode
func TestBase64Decode(t *testing.T) {
	input := "5oiRYeaYr0LkuK0x5Zu95Lq6"
	output := BASE64Decode(&input)
	expected := "我a是B中1国人"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkBase64Decode
func BenchmarkBase64Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "5oiRYeaYr0LkuK0x5Zu95Lq6"
		BASE64Decode(&input)
	}
}

// TestAESEncrypt
func TestAESEncrypt(t *testing.T) {
	input := "我a是B中1国人"
	// key: 16, 24, 32, AES-128, AES-192, AES-256
	output := AESEncrypt(&input, "0123456789abcdef")
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", *output.Value)
		} else {
			encrypt := AESDecrypt(output.Value, "0123456789abcdef")
			if *encrypt.Value != input {
				t.Errorf("expected: %s, actual: %s", *encrypt.Value, *output.Value)
			} else {
				fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
			}
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkAESEncrypt
func BenchmarkAESEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		AESEncrypt(&input, "0123456789abcdef")
	}
}

// TestAESDecrypt
func TestAESDecrypt(t *testing.T) {
	// key: 16, 24, 32, AES-128, AES-192, AES-256
	input := "93fc4b3f26d42ecaa4f9ceb7df8cda734bdbe417f7d0733fe87b99a5dba49467c58a98f2ae26a9ade9b2f2c001d478e5"
	output := AESDecrypt(&input, "0123456789abcdef")
	expected := "我a是B中1国人"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", expected)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkAESDecrypt
func BenchmarkAESDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "93fc4b3f26d42ecaa4f9ceb7df8cda734bdbe417f7d0733fe87b99a5dba49467c58a98f2ae26a9ade9b2f2c001d478e5"
		AESDecrypt(&input, "0123456789abcdef")
	}
}

// TestDESEncrypt
func TestDESEncrypt(t *testing.T) {
	input := "我a是B中1国人"
	output := DESEncrypt(&input, "01234567")
	expected := "24ed0d5e3c1d21b5b6e267c6ba8b04a02f54087673b0a0eeaf402704ed773f58"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", *output.Value)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkDESEncrypt
func BenchmarkDESEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		DESEncrypt(&input, "01234567")
	}
}

// TestDESDecrypt
func TestDESDecrypt(t *testing.T) {
	input := "24ed0d5e3c1d21b5b6e267c6ba8b04a02f54087673b0a0eeaf402704ed773f58"
	output := DESDecrypt(&input, "01234567")
	expected := "我a是B中1国人"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", *output.Value)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkDESDecrypt
func BenchmarkDESDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		DESEncrypt(&input, "01234567")
	}
}

// TestSM3
func TestSM3(t *testing.T) {
	input := "我a是B中1国人"
	output := SM3(&input)
	expected := "7f5121fe72732d88a4fac6a6061fc7f5c64105fa0e7c353150ca0e995506079f"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", *output.Value)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkSM3
func BenchmarkSM3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		SM3(&input)
	}
}

// TestSM4Encrypt
func TestSM4Encrypt(t *testing.T) {
	input := "我a是B中1国人"
	output := SM4Encrypt(&input, "0123456789abcdef")
	expected := "de22c29a0f01e1f911e4f1c06a554882a5a46aec918226d71fa903c82c17015f"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", *output.Value)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkSM4Encrypt
func BenchmarkSM4Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "我a是B中1国人"
		SM4Encrypt(&input, "0123456789abcdef")
	}
}

// TestSM4Decrypt
func TestSM4Decrypt(t *testing.T) {
	input := "de22c29a0f01e1f911e4f1c06a554882a5a46aec918226d71fa903c82c17015f"
	output := SM4Decrypt(&input, "0123456789abcdef")
	expected := "我a是B中1国人"
	if output.Valid {
		if output.Value == nil {
			t.Errorf("expected: %s, actual: nil", *output.Value)
		} else if *output.Value != expected {
			t.Errorf("expected: %s, actual: %s", expected, *output.Value)
		} else {
			fmt.Println(fmt.Sprintf("Input: [%s], Output: [%s]", input, *output.Value))
		}
	} else {
		t.Errorf("expected: true, actual: %t", output.Valid)
	}
}

// BenchmarkSM4Decrypt
func BenchmarkSM4Decrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "de22c29a0f01e1f911e4f1c06a554882a5a46aec918226d71fa903c82c17015f"
		SM4Decrypt(&input, "0123456789abcdef")
	}
}
