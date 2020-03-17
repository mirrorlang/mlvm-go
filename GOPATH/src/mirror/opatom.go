package mirror

import (
	"encoding/json"
	"fmt"
)

type OpAtom struct {
	Op     string //+ - * / nil
	Nextop Point
}

func (b OpAtom) Type() string {
	return "op"
}
func (b OpAtom) String() string {
	return fmt.Sprint(b.Op)
}
func (b OpAtom) Name() string {
	return fmt.Sprint(b.Op)
}
func (b OpAtom) Tomap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	bs, _ := json.Marshal(b)
	json.Unmarshal(bs, &m)
	m["Type"] = b.Type()
	return
}
