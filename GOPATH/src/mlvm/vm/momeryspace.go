package vm

import (
	"fmt"
	"mirror"
)

type Memoryspace struct {
	N     int
	Space [][]mirror.Atom
}

func NewMemory() (m *Memoryspace) {
	m = new(Memoryspace)
	m.Space = make([][]mirror.Atom, 10)
	for i := 0; i < 10; i++ {
		m.Space[i] = make([]mirror.Atom, 10)
	}
	return m
}
func (m *Memoryspace) Print() {
	for i := 0; i < len(m.Space); i++ {
		for j := 0; j < len(m.Space[i]); j++ {
			atom := m.Space[j][i]
			switch atom.Type {
			case "":
				fallthrough
			case "null":
				fallthrough
			case "string":
				fmt.Printf("|%-6s", atom.V_string)
			case "op":
				fmt.Printf("|%-6s", atom.Operator)
			case "int":
				fmt.Printf("|%d", atom.V_int)
			case "point":
				fmt.Printf("|(%d,%d)", atom.Point_x, atom.Point_y)
			}
		}
		fmt.Println()
	}
}
