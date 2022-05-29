package test

import (
	"github.com/beevik/etree"
	"mlvm_go/mirror"
	"mlvm_go/mlvm/vm/mem"
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

func TestFunc(r *mem.Memoryspace, s mirror.Point) {
	r.Set(s.X, s.Y, &mirror.FuncAtom{
		Funcbody: mirror.Rect{Size_y: 5, Size_x: 6, Point: s},
		Args:     mirror.Rect{Point: mirror.Point{X: 1, Isoffset: true}, Size_x: 2, Size_y: 1},
		Value:    mirror.Rect{Point: mirror.Point{X: 3, Isoffset: true}, Size_x: 2, Size_y: 1},
		Name:     "addrex",
		Nextop:   mirror.Point{Isoffset: true, X: 4, Y: 2},
	}) //func addr
	r.Set(s.X+1, s.Y, &mirror.NumAtom{Name: "a"})
	r.Set(s.X+2, s.Y, &mirror.NumAtom{Name: "b"})

	r.Set(s.X+3, s.Y, &mirror.NumAtom{Name: "R1"})
	r.Set(s.X+4, s.Y, &mirror.NumAtom{Name: "R2"})

	r.Set(s.X+1, s.Y+2, &mirror.PointAtom{Point: mirror.Point{Isoffset: true, X: +1, Y: -2}})
	r.Set(s.X+2, s.Y+2, &mirror.OpAtom{Op: "=", Nextop: mirror.Point{Isoffset: true, X: 2, Y: 1}})
	r.Set(s.X+3, s.Y+2, &mirror.PointAtom{Point: mirror.Point{Isoffset: true, X: -3, Y: -2}})
	r.Set(s.X+4, s.Y+2, &mirror.OpAtom{Op: "-", Nextop: mirror.Point{Isoffset: true, X: -2}})
	r.Set(s.X+5, s.Y+2, &mirror.PointAtom{Point: mirror.Point{Isoffset: true, X: -2, Y: -2}})

	r.Set(s.X+1, s.Y+3, &mirror.PointAtom{Point: mirror.Point{Isoffset: true, X: +2, Y: -3}})
	r.Set(s.X+2, s.Y+3, &mirror.OpAtom{Op: "=", Nextop: mirror.Point{Isoffset: true, Y: 1}})
	r.Set(s.X+3, s.Y+3, &mirror.PointAtom{Point: mirror.Point{Isoffset: true, X: -3, Y: -3}})
	r.Set(s.X+4, s.Y+3, &mirror.OpAtom{Op: "+", Nextop: mirror.Point{Isoffset: true, X: -2}})
	r.Set(s.X+5, s.Y+3, &mirror.PointAtom{Point: mirror.Point{Isoffset: true, X: -2, Y: -3}})

	r.Set(s.X+2, s.Y+4, &mirror.ReturnAtom{})
}

func TestCallfunc(r *mem.Memoryspace, s mirror.Point) {

	r.Set(s.X, s.Y, &mirror.FuncAtom{
		Name:     "main",
		Nextop:   mirror.Point{Isoffset: true, X: 3, Y: 2},
		Funcbody: mirror.Rect{Point: mirror.Point{X: 0, Y: 0}, Size_y: 4, Size_x: 6},
	})

	r.Set(s.X, s.Y+1, &mirror.NumAtom{IntValue: 3})
	r.Set(s.X+1, s.Y+1, &mirror.NumAtom{IntValue: 5})

	// 第一种调用方式
	r.Set(s.X, s.Y+2, &mirror.NumAtom{Name: "r1"})
	r.Set(s.X+1, s.Y+2, &mirror.NumAtom{Name: "r2"})
	//func point
	r.Set(s.X+2, s.Y+2, &mirror.OpAtom{Op: "=", Nextop: mirror.Point{
		X:        0,
		Y:        1,
		Isoffset: true,
	}})
	r.Set(s.X+3, s.Y+2, &mirror.CallAtom{
		Nextop: mirror.Point{Isoffset: true, X: -1, Y: 0},
		Func:   mirror.GPoint{Y: 10, X: 0},
		Args:   &mirror.Rect{Point: mirror.Point{X: -3, Y: -1, Isoffset: true}, Size_x: 2, Size_y: 1},
	})
}
