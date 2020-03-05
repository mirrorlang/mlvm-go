package vm

import (
	"fmt"
	"mirror"
	"strconv"
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

const Width = 10

func (m *Memoryspace) Print() {
	fmt.Print("mem:")
	for j := 0; j < len(m.Space[0]); j++ {
		fmt.Printf("|x=%-"+strconv.Itoa(Width-2)+"d", j)
	}
	fmt.Println()
	for i := 0; i < len(m.Space); i++ {
		fmt.Print("|y=" + strconv.Itoa(i))
		for j := 0; j < len(m.Space[i]); j++ {
			atom := m.Space[i][j]
			if atom == nil {
				fmt.Printf("|%"+strconv.Itoa(Width)+"s", "")
				continue
			}
			switch atom.Type {
			case "":
				fallthrough
			case "string":
				fmt.Printf("|%-"+strconv.Itoa(Width)+"s", atom.V_string)
			case "bool":
				fmt.Printf("|%-"+strconv.Itoa(Width)+"t", atom.V_bool)
			case "op":
				fmt.Printf("|%-"+strconv.Itoa(Width)+"s", atom.Operator)
			case "int":
				fmt.Printf("|%-"+strconv.Itoa(Width)+"d", atom.V_int)
			case "point":
				fmt.Printf("|pr(%-"+strconv.Itoa((Width-4)/2)+"d%-"+strconv.Itoa((Width-4)/2)+"d)", atom.Point_x, atom.Point_y)
			}
		}
		fmt.Println("|")
	}
}
