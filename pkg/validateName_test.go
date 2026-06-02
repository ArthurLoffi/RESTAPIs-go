package pkg

import "testing"

func TestValidateName(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected bool
	}{
		{"nome simples", "Arthur", true},
		{"nome com espaço", "Arthur Loffi", true},
		{"nome com acento", "João", true},
		{"nome com cedilha", "François", true},
		{"nome com espaços múltiplos", "Maria José Silva", true},

		{"string vazia", "", false},
		{"número", "Arthur123", false},
		{"caractere especial arroba", "Arthur@", false},
		{"caractere especial hífen", "Arthur-Loffi", false},
		{"underscore", "Arthur_Loffi", false},
		{"ponto", "Arthur.", false},
		{"apenas números", "12345", false},
		{"injeção SQL", "'; DROP TABLE users; --", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateName(tt.input)
			if got != tt.expected {
				t.Errorf("ValidateName(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}