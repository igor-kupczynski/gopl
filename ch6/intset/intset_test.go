package intset

import (
	"reflect"
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	type fields struct {
		words []uint
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []uint
	}{
		{
			"add to empty set",
			fields{},
			args{0},
			[]uint{1},
		},
		{
			"add new word if needed",
			fields{make([]uint, 1)},
			args{wsize},
			[]uint{0, 1},
		},
		{
			"no change if argument is present",
			fields{[]uint{1}},
			args{0},
			[]uint{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			s.Add(tt.args.x)
			if got := s.words; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntSet.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_AddAll(t *testing.T) {
	type fields struct {
		words []uint
	}
	tests := []struct {
		name   string
		fields fields
		args   []int
		want   []uint
	}{
		{
			"add to empty set",
			fields{},
			[]int{0, 1, 2, 3, 4},
			[]uint{0x1f},
		},
		{
			"add new word if needed",
			fields{make([]uint, 1)},
			[]int{wsize*2 + 1, wsize},
			[]uint{0, 1, 2},
		},
		{
			"no change if argument is present",
			fields{[]uint{3}},
			[]int{0, 1},
			[]uint{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			s.AddAll(tt.args...)
			if got := s.words; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntSet.AddAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Clear(t *testing.T) {
	type fields struct {
		words []uint
	}
	tests := []struct {
		name   string
		fields fields
		want   []uint
	}{
		{
			"clears the set",
			fields{},
			nil,
		},
		{
			"clears the set",
			fields{[]uint{0, 1}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			s.Clear()
			if got := s.words; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntSet.Clear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Elems(t *testing.T) {
	type fields struct {
		words []uint
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			"empty set yields empty slice",
			fields{},
			nil,
		},
		{
			"empty set yields empty slice",
			fields{[]uint{0}},
			nil,
		},
		{
			"return elements present in the set",
			fields{[]uint{0x1248}},
			[]int{3, 6, 9, 12},
		},
		{
			"return elements present in the set",
			fields{[]uint{0x1248, 0, 0xf0}},
			[]int{3, 6, 9, 12, wsize*2 + 4, wsize*2 + 5, wsize*2 + 6, wsize*2 + 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			got := s.Elems()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntSet.Elems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Copy(t *testing.T) {
	type fields struct {
		words []uint
	}
	tests := []struct {
		name   string
		fields fields
		want   []uint
	}{
		{
			"copies the set",
			fields{},
			nil,
		},
		{
			"copies the set",
			fields{[]uint{0, 1}},
			[]uint{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			got := s.Copy().words
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntSet.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Has(t *testing.T) {
	type fields struct {
		words []uint
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"empty set has no integers",
			fields{},
			args{0},
			false,
		},
		{
			"empty set has no integers",
			fields{},
			args{1},
			false,
		},
		{
			"empty set has no integers",
			fields{},
			args{15},
			false,
		},
		{
			"included integer is reported",
			fields{[]uint{0xf0}},
			args{4},
			true,
		},
		{
			"not include integer is not reported",
			fields{[]uint{0xf0}},
			args{3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			if got := s.Has(tt.args.x); got != tt.want {
				t.Errorf("IntSet.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Len(t *testing.T) {
	type fields struct {
		words []uint
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"empty set has zero length",
			fields{},
			0,
		},
		{
			"empty set has zero length",
			fields{[]uint{0}},
			0,
		},
		{
			"count the total number of members",
			fields{[]uint{0xf}},
			4,
		},
		{
			"count the total number of members in all words",
			fields{[]uint{0xf, 0x0, 0xf}},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("IntSet.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Remove(t *testing.T) {
	type fields struct {
		words []uint
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []uint
	}{
		{
			"do nothing on empty set",
			fields{},
			args{0},
			nil,
		},
		{
			"do nothing on empty set",
			fields{[]uint{0}},
			args{2},
			[]uint{0},
		},
		{
			"do nothing if the element doesn't exist",
			fields{[]uint{1}},
			args{2},
			[]uint{1},
		},
		{
			"remove the element from the set",
			fields{[]uint{0x10}},
			args{4},
			[]uint{0},
		},
		{
			"remove the element from the set when multiple words",
			fields{[]uint{0xf, 0xf0}},
			args{wsize + 4},
			[]uint{0xf, 0xe0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			s.Remove(tt.args.x)
			if !reflect.DeepEqual(s.words, tt.want) {
				t.Errorf("IntSet.Remove() = %v, want %v", s, tt.want)
			}
		})
	}
}
