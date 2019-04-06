package main

import "testing"

func TestCountDifferentBits(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "same stuff has 0 bit difference",
			args: args{
				[]byte("foo"),
				[]byte("foo"),
			},
			want: 0,
		},
		{
			name: "count the different bits",
			args: args{
				[]byte("foo"),
				[]byte("bar"),
			},
			want: 150,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountDifferentBits(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CountDifferentBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
