package anagram

import "testing"

func TestAnagram(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"detects anagrams", args{"igor", "rogi"}, true},
		{"detects short anagrams", args{"a", "a"}, true},
		{"detects not anagrams", args{"banana", "banana"}, false},
		{"anagrams have the same length", args{"banana", "aananab"}, false},
		{"anagrams can use non western script", args{"Hello, 世界", "界世 ,olleH"}, true},
		{"anagrams can have accents", args{"jeść", "ćśej"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagram(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("anagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
