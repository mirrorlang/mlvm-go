package vm

import (
	"fmt"
	"mirror"
)

type Memoryspace struct {
	n     int
	Space [][]*mirror.Atom
}

func NewMemory() (m *Memoryspace) {
	m = new(Memoryspace)
	m.n = 3
	m.Resize(m.n)
	return m
}
func (m *Memoryspace) Resize(N int) {
	m.n = N
	leng := (2 << m.n)
	if m.Space != nil {
		m.Space = append(m.Space, make([][]*mirror.Atom, leng-len(m.Space))...)
		for i := 0; i < leng; i++ {
			m.Space[i] = append(m.Space[i], make([]*mirror.Atom, leng-len(m.Space[i]))...)
		}
	} else {
		m.Space = make([][]*mirror.Atom, leng)
		for i := 0; i < leng; i++ {
			m.Space[i] = make([]*mirror.Atom, leng)
		}
	}

}
func (m *Memoryspace) Print() {
	for i := 0; i < len(m.Space); i++ {
		for j := 0; j < len(m.Space[i]); j++ {
			atom := m.Space[i][j]
			if atom == nil {
				fmt.Printf("|%6s", "")
				continue
			}
			switch atom.Type {
			case "":
				fallthrough
			case "null":
				fallthrough
			case "string":
				fmt.Printf("|%6s", atom.V_string)
			case "bool":
				fmt.Printf("|%6t", atom.V_bool)
			case "op":
				fmt.Printf("|%-6s", atom.Operator)
			case "int":
				fmt.Printf("|%d", atom.V_int)
			case "point":
				fmt.Printf("|(y=%d,x=%d)", atom.Point_y, atom.Point_x)
			}
		}
		fmt.Println()
	}
}
