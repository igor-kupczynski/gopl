package comma

import "testing"

func TestComma(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"should add comma to numbers", "12345", "12,345"},
		{"should not add comma to short numbers", "123", "123"},
		{"should add multiples commas to long numbers", "12345678", "12,345,678"},
		{"should handle negative numbers", "-12345", "-12,345"},
		{"should handle positive numbers", "+123", "+123"},
		{"should handle decimals", "12345.6789", "12,345.6789"},
		{"should handle short decimals", "123.4567", "123.4567"},
		{"should handle everything", "-12345678.9876", "-12,345,678.9876"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.in); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}
		})
	}
}
