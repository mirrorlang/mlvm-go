package mirror

import "fmt"

type GotoAtom struct {
	Op string //go goto
}

func (b GotoAtom) Type() string {
	return "controlflow"
}
func (b GotoAtom) ControlflowType() string {
	return "Goto"
}
func (b GotoAtom) String() string {
	return fmt.Sprint(b.Op)
}

type IfAtom struct {
	op   string //if,ifelse
	True Point
	Else Point
}

func (b IfAtom) ControlflowType() string {
	return "if"
}
func (b IfAtom) Type() string {
	return "controlflow"
}
func (b IfAtom) String() string {
	return "if"
}

type SwitchAtom struct {
	op      string //switch
	Case    map[Point]Point
	Default Point
}

func (b SwitchAtom) ControlflowType() string {
	return "switch"
}
func (b SwitchAtom) Type() string {
	return "controlflow"
}
func (b SwitchAtom) String() string {
	return "switch..."
}
