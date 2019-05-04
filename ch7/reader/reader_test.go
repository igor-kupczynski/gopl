package reader

import (
	"bytes"
	"io"
	"math"
	"reflect"
	"testing"
)

const bufSize = 3

func TestLimitReader(t *testing.T) {

	tests := []struct {
		name  string
		input []byte
		limit int64
		want  [][]byte
	}{
		{
			"should read the input normally if under limit",
			[]byte("foobar"),
			10,
			[][]byte{[]byte("foo"), []byte("bar"), {}},
		},
		{
			"should limit the read",
			[]byte("foobar"),
			2,
			[][]byte{[]byte("fo"), {}},
		},
		{
			"should limit the read across multiple reads",
			[]byte("foobar"),
			5,
			[][]byte{[]byte("foo"), []byte("ba"), {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := LimitReader(bytes.NewReader(tt.input), tt.limit)

			wantTotal := 0
			gotTotal := 0
			got := make([][]byte, 0, len(tt.want))

			for _, wantRound := range tt.want {
				wantTotal += len(wantRound)
				gotRound := make([]byte, bufSize)

				n, err := r.Read(gotRound)
				if n < bufSize {
					gotRound = gotRound[:n]
				}
				gotTotal += n

				if err != nil && err != io.EOF {
					t.Errorf("Unexpected error while reading: %v", err)
				}

				got = append(got, gotRound)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LimitReader.Read: got %v, want %v", got, tt.want)
			}

			if gotTotal != int(math.Min(float64(tt.limit), float64(len(tt.input)))) {
				t.Errorf("LimitReader.Read: got %v, want %v", got, tt.want)
			}
		})
	}
}
