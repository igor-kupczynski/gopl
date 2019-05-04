package writer

import (
	"bytes"
	"reflect"
	"testing"
)

func TestWordCounter_Write(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    int
		wantErr bool
	}{
		{
			"report 0 for nil input",
			nil,
			0,
			false,
		},
		{
			"report 0 for empty input",
			[]byte{},
			0,
			false,
		},
		{
			"count words",
			[]byte("foo"),
			1,
			false,
		},
		{
			"count words",
			[]byte("foo bar baz"),
			3,
			false,
		},
		{
			"count words",
			[]byte("foo bar baz\nqynx"),
			4,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w WordCounter
			_, err := w.Write(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WordCounter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if int(w) != tt.want {
				t.Errorf("WordCounter.Write() = %v, want %v", int(w), tt.want)
			}
		})
	}
}

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		name   string
		inputs [][]byte
	}{
		{
			"count the written bytes",
			[][]byte{
				[]byte("foo"),
				[]byte("bar"),
				[]byte("baza"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotOutput bytes.Buffer
			var wantOutput bytes.Buffer

			w, n := CountingWriter(&gotOutput)

			total := int64(0)
			gotCounts := []int64{total}
			wantCounts := []int64{*n}

			for _, input := range tt.inputs {
				cnt, err := w.Write(input)
				if err != nil {
					t.Errorf("CountingWriter.Write resulted in an error: %v", err)
				}

				total += int64(cnt)
				wantCounts = append(wantCounts, total)
				gotCounts = append(gotCounts, *n)

				wantOutput.Write(input)
			}

			if !reflect.DeepEqual(wantOutput.Bytes(), gotOutput.Bytes()) {
				t.Errorf("Written: want = %v, have = %v", wantOutput.String(), gotOutput.String())
			}

			if !reflect.DeepEqual(wantCounts, gotCounts) {
				t.Errorf("Counts: want = %v, have = %v", wantCounts, gotCounts)
			}

		})
	}
}
