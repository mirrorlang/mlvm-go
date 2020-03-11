package mirror

import (
	"fmt"
	"strconv"
)

type Atom struct {
	Type               string //null,bool,int,string,point,rectpoint,operator
	V_bool             bool
	Name               string
	V_int              int
	V_string           string
	X, Y               int
	Offset_x, Offset_y int

	Size_x, Size_y             int
	Rectoffset_x, Rectoffset_y int
	Operator                   string
}

func (a *Atom) Clone() Atom {
	if a == nil {
		return Atom{}
	}
	return *a
}
func (a *Atom) Value() interface{} {
	if a == nil {
		return nil
	}
	switch a.Type {
	case "null":
		return ""
	case "int":
		return a.V_int
	case "string":
		return a.V_string
	case "bool":
		return a.V_bool
	case "func":
		return a.Name + "()"
	case "op":
		return a.Operator
	case "point":
		fallthrough
	case "rect":
		fallthrough
	case "rectdata":
		return a.String()
	default:
		return "unknow"
	}
}

func (a *Atom) String() string {
	switch a.Type {
	case "point":
		if a.X == 0 && a.Y == 0 {
			s := "·(" + " "
			if a.Offset_x >= 0 {
				s += "+" + strconv.Itoa(a.Offset_x)
			} else {
				s += strconv.Itoa(a.Offset_x)
			}
			s += ","
			s += " "
			if a.Offset_y >= 0 {
				s += "+" + strconv.Itoa(a.Offset_y)
			} else {
				s += strconv.Itoa(a.Offset_y)
			}
			s += ")"
			return s
		} else {
			s := "·(" + strconv.Itoa(a.X)
			if a.Offset_x >= 0 {
				s += "+" + strconv.Itoa(a.Offset_x)
			} else {
				s += strconv.Itoa(a.Offset_x)
			}
			s += ","
			s += strconv.Itoa(a.Y)
			if a.Offset_y >= 0 {
				s += "+" + strconv.Itoa(a.Offset_y)
			} else {
				s += strconv.Itoa(a.Offset_y)
			}
			s += ")"
			return s
		}

	case "rectdata":
		s := "↴("
		if a.Offset_x >= 0 {
			s += "+" + strconv.Itoa(a.Offset_x)
		} else {
			s += strconv.Itoa(a.Offset_x)
		}
		s += ","
		if a.Offset_y >= 0 {
			s += "+" + strconv.Itoa(a.Offset_y)
		} else {
			s += strconv.Itoa(a.Offset_y)
		}
		s += ").("
		if a.Rectoffset_x >= 0 {
			s += "+" + strconv.Itoa(a.Rectoffset_x)
		} else {
			s += strconv.Itoa(a.Rectoffset_x)
		}
		s += ","
		if a.Rectoffset_y >= 0 {
			s += "+" + strconv.Itoa(a.Rectoffset_y)
		} else {
			s += strconv.Itoa(a.Rectoffset_y)
		}
		s += ")"
		return s

	case "rect":
		s := "□["
		if a.Y == 0 && a.X == 0 {
			s += " "
		} else {
			s += strconv.Itoa(a.X)
		}

		if a.Offset_x >= 0 {
			s += "+" + strconv.Itoa(a.Offset_x)
		} else {
			s += strconv.Itoa(a.Offset_x)
		}
		if a.Size_x >= 0 {
			s += "_" + strconv.Itoa(a.Size_x)
		} else {
			panic("size <0")
		}
		s += ","
		if a.Y == 0 && a.X == 0 {
			s += " "
		} else {
			s += strconv.Itoa(a.Y)
		}
		if a.Offset_y >= 0 {
			s += "+" + strconv.Itoa(a.Offset_y)
		} else {
			s += strconv.Itoa(a.Offset_y)
		}
		if a.Size_y >= 0 {
			s += "|" + strconv.Itoa(a.Size_y)
		} else {
			panic("size <0")
		}
		s += "]"
		return s
	case "func":
		s := a.Name + "()"
		s += "{_" + strconv.Itoa(a.Size_x) + "|" + strconv.Itoa(a.Size_y) + "}"
		return s
	case "":
		return ""
	default:
		return fmt.Sprint(a.Value())
	}
}
