package feistel_test

import (
	"io"
	"strings"
	"tehrelt/bpid/cipher/pkg/feistel"
	"testing"
)

func TestFeistelCipher(t *testing.T) {
	tests := []struct {
		input string
		key   string
	}{
		{"hello", "qwerty"},
		{"world", "hello"},
	}

	for _, test := range tests {

		key := feistel.GenerateKeyFromString(test.key)
		cipher := feistel.New(key)

		encrypted, err := cipher.Encrypt(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Encrypt(%q) throw error: %v", test.input, err)
		}

		decrypted, err := cipher.Decrypt(encrypted)
		if err != nil {
			t.Errorf("Decrypt(%q) throw error: %v", test.input, err)
		}

		dec, err := io.ReadAll(decrypted)
		if err != nil {
			t.Errorf("Ошибка при чтении: %v", err)
		}

		if string(dec) != test.input {
			t.Errorf(test.input + " не совпадает с " + string(dec))
		}

		t.Logf("%q -> %q", test.input, string(dec))
	}
}
