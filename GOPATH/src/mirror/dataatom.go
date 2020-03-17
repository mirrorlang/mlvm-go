package mirror

import (
	"fmt"
)

type NullAtom struct {
	name string
}

func (b NullAtom) Type() string {
	return "null"
}
func (b NullAtom) String() string {
	return ""
}
func (b NullAtom) Name() string {
	return b.name
}

type BoolAtom struct {
	Value bool
	name  string
}

func (b BoolAtom) Type() string {
	return "bool"
}
func (b BoolAtom) String() string {
	return fmt.Sprint(b.Value)
}

func (b BoolAtom) Name() string {
	return b.name
}

type StringAtom struct {
	value string
	name  string
}

func (b StringAtom) Type() string {
	return "string"
}
func (b StringAtom) String() string {
	return b.value
}
func (b StringAtom) Name() string {
	return b.name
}

type NumAtom struct {
	IntValue int
	name     string
}

func (b NumAtom) Type() string {
	return "num"
}
func (b NumAtom) String() string {
	return fmt.Sprint(b.IntValue)
}
func (b NumAtom) Name() string {
	return b.name
}

//type Atom struct {
//	Type               string //null,bool,int,string,point,rectpoint,operator
//	V_bool             bool
//	Name               string
//	V_int              int
//	V_string           string
//	X, Y               int
//	Offset_x, Offset_y int
//
//	Size_x, Size_y             int
//	Rectoffset_x, Rectoffset_y int
//	Operator                   string
//}
