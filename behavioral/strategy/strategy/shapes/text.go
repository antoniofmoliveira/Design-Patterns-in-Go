package shapes

import (
	"github.com/antoniofmoliveira/patterns/behavioral/strategy/strategy"
)

type TextSquare struct {
	strategy.DrawOutput
}

// func (t *TextSquare) Print() error {
// 	r := bytes.NewReader([]byte("Square"))
// 	io.Copy(t.Writer, r)
// 	return nil
// }

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Square"))
	return nil
}
