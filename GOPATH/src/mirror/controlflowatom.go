package mirror

import (
	"encoding/json"
	"fmt"
)

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

func (b GotoAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}

type IfAtom struct {
	Op   string //if,ifelse
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
func (b IfAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
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
func (b SwitchAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}
