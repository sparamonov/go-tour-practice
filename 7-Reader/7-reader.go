package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (rdr MyReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	
	cnt := 0
	
	for i := range p {
		cnt++
		p[i] = 'A'		
	}
	return cnt, nil
}

func main() {
	reader.Validate(MyReader{})
}
