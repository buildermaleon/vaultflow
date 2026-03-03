package crypto

import (
	"strings"
	"testing"
)

func TestEncryptor_EncryptDecrypt(t *testing.T) {
	e := New("test-master-key-32bytes!")
	
	plaintext := "Hello, World!"
	
	ciphertext, err := e.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encrypt() error = %v", err)
	}
	
	decrypted, err := e.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("Decrypt() error = %v", err)
	}
	
	if decrypted != plaintext {
		t.Errorf("Decrypted = %v, want %v", decrypted, plaintext)
	}
}

func TestEncryptor_EmptyString(t *testing.T) {
	e := New("test-key-32-bytes!!!")
	
	ciphertext, _ := e.Encrypt("")
	decrypted, err := e.Decrypt(ciphertext)
	
	if err != nil {
		t.Fatalf("Error = %v", err)
	}
	if decrypted != "" {
		t.Error("Should handle empty string")
	}
}

func TestEncryptor_LongString(t *testing.T) {
	e := New("test-key-32-bytes!!!")
	
	plaintext := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
	Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`
	
	ciphertext, _ := e.Encrypt(plaintext)
	decrypted, err := e.Decrypt(ciphertext)
	
	if err != nil {
		t.Fatalf("Error = %v", err)
	}
	if decrypted != plaintext {
		t.Error("Long string encryption failed")
	}
}

func TestEncryptor_SpecialChars(t *testing.T) {
	e := New("test-key-32-bytes!!!")
	
	testCases := []string{
		"!@#$%^&*()",
		"中文测试",
		"🎉 emojis",
		"new\nline\ttab",
	}
	
	for _, tc := range testCases {
		ciphertext, _ := e.Encrypt(tc)
		decrypted, err := e.Decrypt(ciphertext)
		if err != nil {
			t.Fatalf("Failed for %q: %v", tc, err)
		}
		if decrypted != tc {
			t.Errorf("Failed for %q", tc)
		}
	}
}

func TestEncryptor_InvalidCiphertext(t *testing.T) {
	e := New("test-key-32-bytes!!!")
	
	_, err := e.Decrypt("invalid-base64!!!")
	if err == nil {
		t.Error("Should error on invalid base64")
	}
}

func TestEncryptor_ShortCiphertext(t *testing.T) {
	e := New("test-key-32-bytes!!!")
	
	_, err := e.Decrypt("YWJjZA==")
	if err == nil {
		t.Error("Should error on short ciphertext")
	}
}

func TestEncryptor_PanicWithEmptyKey(t *testing.T) {
	e := New("")
	
	ciphertext, err := e.Encrypt("test")
	if err != nil {
		t.Fatalf("Encrypt error = %v", err)
	}
	
	decrypted, err := e.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("Decrypt error = %v", err)
	}
	
	if decrypted != "test" {
		t.Errorf("Decrypted = %v, want test", decrypted)
	}
}

func TestEncryptor_KeyPadding(t *testing.T) {
	e := New("short")
	
	ciphertext, _ := e.Encrypt("test")
	decrypted, _ := e.Decrypt(ciphertext)
	
	if decrypted != "test" {
		t.Error("Short key encryption failed")
	}
}

func TestEncryptor_Exactly32Bytes(t *testing.T) {
	e := New("12345678901234567890123456789012")
	
	ciphertext, _ := e.Encrypt("test")
	decrypted, _ := e.Decrypt(ciphertext)
	
	if decrypted != "test" {
		t.Error("32-byte key encryption failed")
	}
}

func TestEncryptor_VeryLongKey(t *testing.T) {
	e := New(strings.Repeat("a", 100))
	
	ciphertext, _ := e.Encrypt("test")
	decrypted, _ := e.Decrypt(ciphertext)
	
	if decrypted != "test" {
		t.Error("Long key encryption failed")
	}
}

func BenchmarkEncrypt(b *testing.B) {
	e := New("benchmark-key-32bytes!!")
	data := "benchmark test data"
	
	for i := 0; i < b.N; i++ {
		e.Encrypt(data)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	e := New("benchmark-key-32bytes!!")
	data := "benchmark test data"
	ciphertext, _ := e.Encrypt(data)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Decrypt(ciphertext)
	}
}
