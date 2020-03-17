package mirror

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X, Y     int
	Isoffset bool
}

func (b Point) GlobalAddr(x, y int) (int, int) {
	if b.Isoffset {
		return b.X + x, b.Y + y
	} else {
		return b.X, b.Y
	}
}

type PointAtom struct {
	Name string
	Point
}

func (b PointAtom) Type() string {
	return "point"
}
func (p PointAtom) String() string {
	return ".(" + fmt.Sprint(p.X) + "," + fmt.Sprint(p.Y) + ")"
}
func (b PointAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}

type Rect struct {
	Size_x, Size_y int
	X, Y           int
}

type RectAtom struct {
	Name string
	Rect
}

func (b RectAtom) Type() string {
	return "rect"
}
func (p RectAtom) String() string {
	return "□[" + fmt.Sprint(p.X) + "_" + fmt.Sprint(p.Size_x) + "," + fmt.Sprint(p.Y) + "|" + fmt.Sprint(p.Size_y) + "]"
}
func (b RectAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}

type RectPointAtom struct {
	Name string
	Point
	Inrect_offset_x, Inrect_offset_y int
}

func (p RectPointAtom) String() string {
	return "□(" + fmt.Sprint(p.X) + "," + fmt.Sprint(p.Y) + ").(" + fmt.Sprint(p.Inrect_offset_x) + "," + fmt.Sprint(p.Inrect_offset_y) + ")"
}

func (b RectPointAtom) Type() string {
	return "rectpoint"
}
func (b RectPointAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}

type FuncAtom struct {
	Name   string
	Nextop Point
	Rect
}

func (b FuncAtom) Type() string {
	return "func"
}
func (b FuncAtom) String() string {
	return "func " + b.Name + "()"
}
func (b FuncAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}
