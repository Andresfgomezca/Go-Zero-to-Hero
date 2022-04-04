package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	//reader r of the struct rot
	n, err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		char := p[i]
		//alphabet divided in 2, if the letter is on of the first thirteen letters it will
		//add 13 to the ascii char
		if (char >= 'A' && char < 'N') || (char >= 'a' && char < 'n') {
			p[i] += 13
			//if the letter is one of the last 13 letters it will substract 13 to the ascii code
		} else if (char > 'M' && char <= 'Z') || (char > 'm' && char <= 'z') {
			p[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
