package vm

import (
	"github.com/beevik/etree"
	"mirror"
	"os"
)

func Loadapp() *etree.Document {
	var apppath = os.Args[1]

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(apppath); err != nil {
		panic(err)
	}
	return doc
}

func TestExpression(r *Memoryspace, start mirror.Atom) {
	r.Space[start.Point_x][start.Point_y] = mirror.Atom{Type: "op", Operator: "null"}
	r.Space[start.Point_x+1][start.Point_y] = mirror.Atom{Type: "int", V_int: 10086}
}
