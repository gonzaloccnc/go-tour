package methods

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func ExerciceRotReader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, (&r).r)
}
