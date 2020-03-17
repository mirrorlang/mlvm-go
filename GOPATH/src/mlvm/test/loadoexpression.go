package test

import (
	"github.com/beevik/etree"
	"mirror"
	"mlvm/vm"
	"os"
	"strings"
)

func Loadapp() *etree.Document {
	var apppath = os.Args[1]

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(apppath); err != nil {
		panic(err)
	}
	return doc
}

func Code(funcarea [][]mirror.Atom) (r string) {
	for j := 0; j < len(funcarea); j++ {
		for i := 0; i < len(funcarea[0]); i++ {
			atom := funcarea[j][i]
			r += atom.String() + "\t"
		}
		r += "\n"

	}
	return
}
func Load(fstr string) (funcarea [][]mirror.Atom) {

	for _, line := range strings.Split(fstr, "\n") {
		exp := make([]mirror.Atom, 0)
		for _, _ = range strings.Split(line, "\t") {
			//exp = append(exp, mirror.Atom{Type: "string", V_string: atom})
		}
		funcarea = append(funcarea, exp)
	}

	return
}
func TestExpression_null(r *vm.Memoryspace, start mirror.PointAtom) {
	r.Set(0, 0, &mirror.OpAtom{Op: "nil", Nextop: mirror.Point{X: 0, Y: 1, Isoffset: true}})
	r.Set(1, 0, &mirror.NumAtom{IntValue: 10086})
	r.Set(0, 1, &mirror.OpAtom{Op: "nil", Nextop: mirror.Point{X: 0, Y: 1, Isoffset: true}})
	r.Set(1, 1, &mirror.PointAtom{Point: mirror.Point{Y: 1, X: 6}})
	r.Set(6, 1, &mirror.NumAtom{IntValue: 10086})
}

//
//func TestExpression_not(r *Memoryspace, start mirror.Atom) {
//	r.space[start.Y][start.X] = &mirror.Atom{Type: "op", Operator: "!"}
//	r.space[start.Y][start.X+1] = &mirror.Atom{Type: "bool", V_bool: true}
//
//}
//
//func TestExpression_goto(r *Memoryspace, start mirror.Atom) {
//	r.space[start.Y][start.X] = &mirror.Atom{Type: "op", Operator: "go"}
//	r.space[start.Y][start.X+1] = &mirror.Atom{Type: "point", Offset_y: 1, Offset_x: 3}
//	r.space[start.Y+1][start.X+3] = &mirror.Atom{Type: "op", Operator: "*"}
//	r.space[start.Y+1][start.X+2] = &mirror.Atom{Type: "int", V_int: 1321}
//	r.space[start.Y+1][start.X+4] = &mirror.Atom{Type: "int", V_int: 124}
//	r.space[start.Y+1][start.X+1] = &mirror.Atom{Type: "op", Operator: "="}
//	r.space[start.Y+2][start.X+3] = &mirror.Atom{Type: "op", Operator: "go"}
//	r.space[start.Y+2][start.X+4] = &mirror.Atom{Type: "point", Offset_y: -1, Offset_x: -2}
//}
//
//func TestFunc(r *Memoryspace) {
//	r.space[2][0] = &mirror.Atom{Type: "func", Size_y: 8, Size_x: 4, Name: "addrex"} //func addr
//	r.space[4][0] = &mirror.Atom{Type: "point", Offset_y: 3, Offset_x: 2}            //first op pisition
//
//	r.space[5][1] = &mirror.Atom{Type: "point", Offset_y: -3, Offset_x: -1} //arg 0
//	r.space[5][2] = &mirror.Atom{Type: "op", Operator: "+"}
//	r.space[5][3] = &mirror.Atom{Type: "point", Offset_y: -2, Offset_x: -1} //arg 1
//	r.space[6][2] = &mirror.Atom{Type: "op", Operator: "="}
//	r.space[6][1] = &mirror.Atom{Type: "rectdata", Offset_x: -2, Offset_y: -1, Rectoffset_x: 0, Rectoffset_y: 0}
//
//	r.space[7][1] = &mirror.Atom{Type: "point", Offset_y: -5, Offset_x: -1} //arg 0
//	r.space[7][2] = &mirror.Atom{Type: "op", Operator: "-"}
//	r.space[7][3] = &mirror.Atom{Type: "point", Offset_y: -4, Offset_x: -1} //arg 1
//	r.space[8][2] = &mirror.Atom{Type: "op", Operator: "="}
//	r.space[8][1] = &mirror.Atom{Type: "rectdata", Offset_x: -2, Offset_y: -3, Rectoffset_x: 0, Rectoffset_y: 1}
//
//	r.space[9][2] = &mirror.Atom{Type: "op", Operator: "return"}
//}
//
//// func add(a,b):return a+b,a-b
//func TestCallfunc(r *Memoryspace) {
//
//	r.space[0][0] = &mirror.Atom{Type: "op", Operator: "rect"}
//	r.space[0][1] = &mirror.Atom{Type: "rect", Y: 0, X: 0, Size_y: 15, Size_x: 5}
//
//	r.space[1][0] = &mirror.Atom{Type: "op", Operator: "go"} //go
//	r.space[1][1] = &mirror.Atom{Type: "point", Offset_y: 9, Offset_x: 2}
//
//	r.space[10][3] = &mirror.Atom{Type: "int", V_int: 1023}
//	r.space[11][3] = &mirror.Atom{Type: "int", V_int: 231}
//
//	r.space[10][2] = &mirror.Atom{Type: "op", Operator: "call"}                                  //call
//	r.space[11][2] = &mirror.Atom{Type: "point", Y: 2, X: 0}                                     //func point
//	r.space[12][2] = &mirror.Atom{Type: "rect", Offset_y: 0, Offset_x: 1, Size_x: 1, Size_y: 2}  //arg rect
//	r.space[13][2] = &mirror.Atom{Type: "rect", Offset_y: 0, Offset_x: -1, Size_x: 1, Size_y: 2} //result rect
//
//}
