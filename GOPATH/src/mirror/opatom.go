package mirror

import "fmt"

type OpAtom struct {
	Op     string //+ - * / nil
	Nextop Point
}

func (b OpAtom) Type() string {
	return "Op"
}
func (b OpAtom) String() string {
	return fmt.Sprint(b.Op)
}
func (b OpAtom) Name() string {
	return fmt.Sprint(b.Op)
}
