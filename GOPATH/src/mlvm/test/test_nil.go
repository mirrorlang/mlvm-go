package test

import (
	"mirror"
	"mlvm/vm/mem"
)

func TestExpression_nil(r *mem.Memoryspace, s mirror.Point) {
	r.Set(s.X+0, s.Y+0, &mirror.OpAtom{Op: "nil", Nextop: mirror.Point{X: 0, Y: 1, Isoffset: true}})
	r.Set(s.X+1, s.Y+0, &mirror.NumAtom{IntValue: 10386})

	r.Set(s.X+0, s.Y+1, &mirror.OpAtom{Op: "nil", Nextop: mirror.Point{X: 0, Y: 1, Isoffset: true}})
	r.Set(s.X+1, s.Y+1, &mirror.PointAtom{Point: mirror.Point{Y: 1, X: 6, Isoffset: false}})
	r.Set(s.X+6, s.Y+1, &mirror.NumAtom{IntValue: 106})

	r.Set(s.X+0, s.Y+2, &mirror.OpAtom{Op: "nil", Nextop: mirror.Point{X: 0, Y: 1, Isoffset: true}})
	r.Set(s.X+1, s.Y+2, &mirror.PointAtom{Point: mirror.Point{Y: 0, X: 6, Isoffset: true}})
	r.Set(s.X+6, s.Y+2, &mirror.NumAtom{IntValue: 10})
}

func TestExpression_not(r *mem.Memoryspace, s mirror.Point) {
	r.Set(s.X+0, s.Y+0, &mirror.OpAtom{Op: "!", Nextop: mirror.Point{X: 0, Y: 1, Isoffset: true}})
	r.Set(s.X+1, s.Y, &mirror.BoolAtom{Value: true})

}

func TestExpression_goto(r *mem.Memoryspace, s mirror.Point) {
	r.Set(s.X, s.Y, &mirror.GotoAtom{Op: "goto"})
	r.Set(s.X+1, s.Y, &mirror.PointAtom{Point: mirror.Point{4, 4, false}})
	r.Set(4, 5, &mirror.GotoAtom{Op: "goto"})
	r.Set(5, 5, &mirror.PointAtom{Point: mirror.Point{0, 6, false}})
}
