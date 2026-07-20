package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(s []byte) (n int, err error) {
	n, err = r.r.Read(s)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		if (s[i] >= 'A' && s[i] <= 'M') || (s[i] >= 'a' && s[i] <= 'm') {
			s[i] += 13
		} else if (s[i] >= 'N' && s[i] <= 'Z') || (s[i] >= 'n' && s[i] <= 'z') {
			s[i] -= 13
		}
	}
	return n, nil

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
