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
	r.Space[start.Point_y][start.Point_x] = &mirror.Atom{Type: "op", Operator: "nil"}
	r.Space[start.Point_y][start.Point_x+1] = &mirror.Atom{Type: "int", V_int: 10086}
	r.Space[start.Point_y+1][start.Point_x] = &mirror.Atom{Type: "op", Operator: "empty"}
	r.Space[start.Point_y+1][start.Point_x+1] = &mirror.Atom{Type: "point", Point_y: 1, Point_x: 6}
	r.Space[1][6] = &mirror.Atom{Type: "int", V_int: 10086}
}

func TestExpression_not(r *Memoryspace, start mirror.Atom) {
	r.Space[start.Point_y][start.Point_x] = &mirror.Atom{Type: "op", Operator: "!"}
	r.Space[start.Point_y][start.Point_x+1] = &mirror.Atom{Type: "bool", V_bool: true}

}

func TestExpression_goto(r *Memoryspace, start mirror.Atom) {
	r.Space[start.Point_y][start.Point_x] = &mirror.Atom{Type: "op", Operator: "go"}
	r.Space[start.Point_y][start.Point_x+1] = &mirror.Atom{Type: "point", Offset_y: 1, Offset_x: 3}
	r.Space[start.Point_y+1][start.Point_x+3] = &mirror.Atom{Type: "op", Operator: "*"}
	r.Space[start.Point_y+1][start.Point_x+2] = &mirror.Atom{Type: "int", V_int: 1321}
	r.Space[start.Point_y+1][start.Point_x+4] = &mirror.Atom{Type: "int", V_int: 124}
	r.Space[start.Point_y+1][start.Point_x+1] = &mirror.Atom{Type: "op", Operator: "="}
	r.Space[start.Point_y+2][start.Point_x+3] = &mirror.Atom{Type: "op", Operator: "go"}
	r.Space[start.Point_y+2][start.Point_x+4] = &mirror.Atom{Type: "point", Offset_y: -1, Offset_x: -2}
}

// func add(a,b):return a+b,a-b
func TestCallfunc(r *Memoryspace) {

	r.Space[0][0] = &mirror.Atom{Type: "op", Operator: "rect"}
	r.Space[0][1] = &mirror.Atom{Type: "rect", Point_y: 0, Point_x: 0, Size_y: 15, Size_x: 5}

	r.Space[1][0] = &mirror.Atom{Type: "op", Operator: "go"} //go
	r.Space[1][1] = &mirror.Atom{Type: "point", Offset_y: 9, Offset_x: 2}

	r.Space[2][0] = &mirror.Atom{Type: "func", Size_y: 8, Size_x: 4, Name: "addrex"} //func addr
	r.Space[4][0] = &mirror.Atom{Type: "point", Offset_y: 3, Offset_x: 2}            //first op pisition

	r.Space[5][1] = &mirror.Atom{Type: "point", Offset_y: -3, Offset_x: -1} //arg 0
	r.Space[5][2] = &mirror.Atom{Type: "op", Operator: "+"}
	r.Space[5][3] = &mirror.Atom{Type: "point", Offset_y: -2, Offset_x: -1} //arg 1
	r.Space[6][2] = &mirror.Atom{Type: "op", Operator: "="}
	r.Space[6][1] = &mirror.Atom{Type: "rectdata", Offset_x: -2, Offset_y: -1, Rectoffset_x: 0, Rectoffset_y: 0}

	r.Space[7][1] = &mirror.Atom{Type: "point", Offset_y: -5, Offset_x: -1} //arg 0
	r.Space[7][2] = &mirror.Atom{Type: "op", Operator: "-"}
	r.Space[7][3] = &mirror.Atom{Type: "point", Offset_y: -4, Offset_x: -1} //arg 1
	r.Space[8][2] = &mirror.Atom{Type: "op", Operator: "="}
	r.Space[8][1] = &mirror.Atom{Type: "rectdata", Offset_x: -2, Offset_y: -3, Rectoffset_x: 0, Rectoffset_y: 1}

	r.Space[9][2] = &mirror.Atom{Type: "op", Operator: "return"}

	r.Space[10][3] = &mirror.Atom{Type: "int", V_int: 1023}
	r.Space[11][3] = &mirror.Atom{Type: "int", V_int: 231}

	r.Space[10][2] = &mirror.Atom{Type: "op", Operator: "call"}                                  //call
	r.Space[11][2] = &mirror.Atom{Type: "point", Point_y: 2, Point_x: 0}                         //func point
	r.Space[12][2] = &mirror.Atom{Type: "rect", Offset_y: 0, Offset_x: 1, Size_x: 1, Size_y: 2}  //arg rect
	r.Space[13][2] = &mirror.Atom{Type: "rect", Offset_y: 0, Offset_x: -1, Size_x: 1, Size_y: 2} //result rect

}
