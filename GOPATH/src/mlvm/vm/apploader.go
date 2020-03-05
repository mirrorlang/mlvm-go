package vm

import (
	"fmt"
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

func TestExpression_null(r *Memoryspace, start mirror.Atom) {
	r.Space[start.Point_y][start.Point_x] = &mirror.Atom{Type: "op", Operator: "null"}
	r.Space[start.Point_y][start.Point_x+1] = &mirror.Atom{Type: "int", V_int: 10086}
	r.Space[start.Point_y+1][start.Point_x] = &mirror.Atom{Type: "op", Operator: "null"}
	r.Space[start.Point_y+1][start.Point_x+1] = &mirror.Atom{Type: "point", Point_y: 1, Point_x: 6}
	r.Space[1][6] = &mirror.Atom{Type: "int", V_int: 10086}
	r.Print()
	fmt.Println()
}

func TestExpression_not(r *Memoryspace, start mirror.Atom) {
	r.Space[start.Point_y][start.Point_x] = &mirror.Atom{Type: "op", Operator: "!"}
	r.Space[start.Point_y][start.Point_x+1] = &mirror.Atom{Type: "bool", V_bool: true}
	r.Print()
	fmt.Println()
}
