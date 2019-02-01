package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	bigA := byte('A')
	bigZ := byte('Z')
	smA := byte('a')
	smZ := byte('z')
	
	for i, v := range b[:n] {
		switch {
		case (v + 13) > smZ:
				b[i] = v - 13
		case (v + 13) >= smA && (v + 13) <= smZ:
				b[i] = v + 13
		case (v + 13) > bigZ:
				b[i] = v - 13
		case (v + 13) >= bigA && (v + 13) <= bigZ:
				b[i] = v + 13
		}
	}
	
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
