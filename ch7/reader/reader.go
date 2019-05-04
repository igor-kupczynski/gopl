package reader

import "io"

// LimitReader returns a wrapped io.Reader instance. The wrapper limits what
// can be read to the first n-bytes. It sends an EOF after that.
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{inner: r, limit: n}
}

type limitReader struct {
	inner io.Reader
	limit int64
}

func (r *limitReader) Read(p []byte) (int, error) {
	if int64(len(p)) > r.limit {
		p = p[:r.limit]
	}
	n, err := r.inner.Read(p)
	r.limit -= int64(n)
	return n, err
}
