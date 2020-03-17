package mirror

import (
	"encoding/json"
	"fmt"
)

type NullAtom struct {
	Name string
}

func (b NullAtom) Type() string {
	return "null"
}
func (b NullAtom) String() string {
	return ""
}
func (b NullAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}

type BoolAtom struct {
	Value bool
	Name  string
}

func (b BoolAtom) Type() string {
	return "bool"
}
func (b BoolAtom) String() string {
	return fmt.Sprint(b.Value)
}
func (b BoolAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}

type StringAtom struct {
	value string
	Name  string
}

func (b StringAtom) Type() string {
	return "string"
}
func (b StringAtom) String() string {
	return b.value
}
func (b StringAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}

type NumAtom struct {
	IntValue int
	Name     string
}

func (b NumAtom) Type() string {
	return "num"
}
func (b NumAtom) String() string {
	return fmt.Sprint(b.IntValue)
}
func (b NumAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
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
