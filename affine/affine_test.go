package affine

import (
	"testing"
)

func TestTableEncrypt(t *testing.T) {
	var tests = []struct {
		plaintext string
		a, b      int
		expected  string
	}{
		{"affinecipher", 5, 8, "IHHWVCSWFRCP"},
		{"affine cipher", 5, 8, "IHHWVCXSPACEXSWFRCP"},
		{"affine 3 cip2her1", 5, 8, "IHHWVCXSPACEXXTHREEXXSPACEXSWFXTWOXRCPXONEX"},
		{"ěščřžýáí", 5, 8, "CUSPDYIW"},
		{"asf", 5, 8, "IUH"},
	}
	for _, test := range tests {
		if output := Encrypt(test.plaintext, test.a, test.b); output != test.expected {
			t.Error("expected ciphertext to be", test.expected, "but got", output)
		}
	}
}

func TestTableDecrypt(t *testing.T) {
	var tests = []struct {
		ciphertext string
		a, b       int
		expected   string
	}{
		{"IHHWVCSWFRCP", 5, 8, "AFFINECIPHER"},
		{"IHHWVCXSPACEXSWFRCP", 5, 8, "AFFINE CIPHER"},
		{"IHHWVCXSPACEXXTHREEXXSPACEXSWFXTWOXRCPXONEX", 5, 8, "AFFINE 3 CIP2HER1"},
		{"CUSPDYIW", 5, 8, "ESCRZYAI"},
	}
	for _, test := range tests {
		if output := Decrypt(test.ciphertext, test.a, test.b); output != test.expected {
			t.Error("expected plaintext to be ", test.expected, "but got", output)
		}
	}
}
