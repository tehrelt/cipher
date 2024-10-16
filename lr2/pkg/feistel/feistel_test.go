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
		keys  []string
	}{
		{"hello", []string{"qwerty", "gahdamn", "rewkash"}},
		{"world", []string{"lorem", "ipsum", "rewkash"}},
	}

	for _, test := range tests {

		keys := feistel.GenerateKeysFromString(test.keys)
		cipher := feistel.New(keys)

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
			t.Errorf("%q не совпадает %q", test.input, string(dec))
		}

		t.Logf("%q -> %q", test.input, string(dec))
	}
}
