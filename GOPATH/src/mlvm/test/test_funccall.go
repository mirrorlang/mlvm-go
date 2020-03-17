package test

import (
	"github.com/beevik/etree"
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
