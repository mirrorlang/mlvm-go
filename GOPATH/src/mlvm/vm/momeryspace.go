package vm

import (
	"fmt"
	"github.com/gosdk/log"

	"mirror"
	"strconv"
)

type Memoryspace struct {
	space [][]mirror.Atom
	n     int
}

func (m *Memoryspace) Y() int {
	return len(m.space)
}

func (m *Memoryspace) X() int {
	return len(m.space[0])
}

func (m *Memoryspace) At(x, y int) mirror.Atom {
	atom := m.space[y][x]
	if atom == nil {
		return mirror.NullAtom{}
	}
	return atom
}

func (m *Memoryspace) Set(x, y int, a mirror.Atom) {
	if a.Type() == "" {
		m.space[y][x] = nil
	} else {
		m.space[y][x] = a
	}
}
func (m *Memoryspace) Rect(point_x, point_y int, offset_x, offset_y int) (r [][]mirror.Atom) {
	r = make([][]mirror.Atom, offset_y)
	for i := 0; i < offset_y; i++ {
		r[i] = make([]mirror.Atom, offset_x)
		for j := 0; j < offset_x; j++ {
			r[i][j] = m.At(point_x+j, point_y+i)
		}
	}
	return
}
func NewMemory() (m *Memoryspace) {
	m = new(Memoryspace)
	m.n = 4
	m.Resize(m.n)
	return m
}

func (m *Memoryspace) Resize(N int) {
	m.n = N
	leng := (2 << m.n)
	if m.space != nil {
		m.space = append(m.space, make([][]mirror.Atom, leng-len(m.space))...)
		for i := 0; i < leng; i++ {
			m.space[i] = append(m.space[i], make([]mirror.Atom, leng-len(m.space[i]))...)
		}
	} else {
		m.space = make([][]mirror.Atom, leng)
		for i := 0; i < leng; i++ {
			m.space[i] = make([]mirror.Atom, leng)
		}
	}
}

const PrintWidth = 20

func (m *Memoryspace) Print() {
	fmt.Print("mem:    ")
	for j := 0; j < m.X(); j++ {
		fmt.Printf("|X=%-"+strconv.Itoa(PrintWidth-3)+"d", j)
	}
	fmt.Println()
	for i := 0; i < m.Y(); i++ {
		fmt.Printf("|Y=%-5d", i)
		for j := 0; j < m.X(); j++ {
			atom := m.At(j, i)
			var c log.Colortext
			switch atom.Type() {
			case "int":
				c = log.Cyan
			case "bool":
				c = log.LightBlue
			case "rect":
				c = log.Blue
			case "rectdata":
				c = log.Pink
			case "point":
				c = log.Gray
			case "func":
				c = log.Yellow
			case "op":
				c = log.Red
			}
			fmt.Print(fmt.Sprintf("|"))
			log.Print(c, fmt.Sprintf("%-"+strconv.Itoa(PrintWidth-1)+"s", atom.String()))
		}
		fmt.Println("|")
	}
	fmt.Println()
}
