package mirror

import (
	"encoding/json"
	"fmt"
)

type CallAtom struct {
	Func   GPoint //+ - * / nil
	Myfunc Point  //+ - * / nil
	Args   *Rect  //可内
	// ，可外
	//可内，直接执行参数范围
	Value  [][]Atom //存放 函数的计算结果
	Nextop Point
}

func (b CallAtom) Type() string {
	return "func"
}
func (b CallAtom) String() string {
	return "call" + fmt.Sprint(b.Func)
}

func (b CallAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	m["Op"] = "call"
	return
}

type FuncAtom struct {
	CallerX      int //调用者的X偏移量
	Funcbody     Rect
	Name         string
	Nextop       Point
	Cpu_x, Cpu_y int
	Args         Rect
	Value        Rect
}

func (b FuncAtom) Type() string {
	return "func"
}
func (b FuncAtom) String() string {
	return "call  " + b.Name + "()"
}
func (b FuncAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	m["Op"] = "func"
	return
}

type ReturnAtom struct {
}

func (b ReturnAtom) Type() string {
	return "func"
}
func (b ReturnAtom) String() string {
	return "return"
}
func (b ReturnAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	m["Op"] = "return"
	return
}
