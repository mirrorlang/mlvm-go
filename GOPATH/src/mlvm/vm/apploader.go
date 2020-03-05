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

func TestExpression_null(r *Memoryspace, start mirror.Atom) {
	r.Space[start.Point_y][start.Point_x] = &mirror.Atom{Type: "op", Operator: "free"}
	r.Space[start.Point_y][start.Point_x+1] = &mirror.Atom{Type: "int", V_int: 10086}
	r.Space[start.Point_y+1][start.Point_x] = &mirror.Atom{Type: "op", Operator: "free"}
	r.Space[start.Point_y+1][start.Point_x+1] = &mirror.Atom{Type: "point", Point_y: 1, Point_x: 6}
	r.Space[1][6] = &mirror.Atom{Type: "int", V_int: 10086}
}

func TestExpression_not(r *Memoryspace, start mirror.Atom) {
	r.Space[start.Point_y][start.Point_x] = &mirror.Atom{Type: "op", Operator: "!"}
	r.Space[start.Point_y][start.Point_x+1] = &mirror.Atom{Type: "bool", V_bool: true}

}

func TestExpression_goto(r *Memoryspace, start mirror.Atom) {
	r.Space[start.Point_y][start.Point_x] = &mirror.Atom{Type: "op", Operator: "go"}
	r.Space[start.Point_y][start.Point_x+1] = &mirror.Atom{Type: "point", Point_y: 1, Point_x: 3}
	r.Space[start.Point_y+1][start.Point_x+3] = &mirror.Atom{Type: "op", Operator: "*"}
	r.Space[start.Point_y+1][start.Point_x+2] = &mirror.Atom{Type: "int", V_int: 1321}
	r.Space[start.Point_y+1][start.Point_x+4] = &mirror.Atom{Type: "int", V_int: 124}
	r.Space[start.Point_y+1][start.Point_x+1] = &mirror.Atom{Type: "op", Operator: "="}
	r.Space[start.Point_y+2][start.Point_x] = &mirror.Atom{Type: "op", Operator: "go"}
	r.Space[start.Point_y+2][start.Point_x+1] = &mirror.Atom{Type: "point", Point_y: -1, Point_x: 2}
}
