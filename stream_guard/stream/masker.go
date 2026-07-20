package stream

import (
	"fmt"
	"io"
)

type ErrSecurityAlert struct {
	Offset int64
	Reason string
}

func (p *ErrSecurityAlert) Error() string {
	return fmt.Sprintf("[ Error ] Reason : %s | Offset : %v ", p.Reason, p.Offset)
}

type MaskingReader struct {
	Source io.Reader
	Offset int64
}

func (r *MaskingReader) Read(p []byte) (n int, err error) {
	n, err = r.Source.Read(p)
	r.Offset += int64(n)
	for i := 0; i < n; i++ {
		if p[i] == '=' {
			p[i] = '*'
		}
	}
	return n, err
}
