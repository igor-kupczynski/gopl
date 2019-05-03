package bytecounter

import (
	"bufio"
	"bytes"
	"io"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// WordCounter is an io.Writer which ignores the input but counts the words
// the input contains
type WordCounter int

var one = WordCounter(1)

func (w *WordCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		*w += one
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return len(p), nil
}

// CountingWriter wraps an io.Writer and returns the wrapped instance
// and a pointer to an int64 value with the current count of the written bytes
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{inner: w}
	return cw, &cw.total
}

type countingWriter struct {
	inner io.Writer
	total int64
}

func (w *countingWriter) Write(p []byte) (int, error) {
	n, err := w.inner.Write(p)
	w.total += int64(n)
	return n, err
}
