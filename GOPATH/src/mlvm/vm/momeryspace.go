package vm

import (
	"fmt"
	"github.com/gosdk/log"

	"mirror"
	"strconv"
)

type Memoryspace struct {
	Space [][]*mirror.Atom
	n     int
}

func (m *Memoryspace) Y() int {
	return len(m.Space)
}

func (m *Memoryspace) X() int {
	return len(m.Space[0])
}

func (m *Memoryspace) At(x, y int) mirror.Atom {
	atom := m.Space[y][x]
	if atom == nil {
		return mirror.Atom{Type: "null"}
	}
	return *atom
}

func (m *Memoryspace) Set(x, y int, a mirror.Atom) {
	if a.Type == "" {
		m.Space[y][x] = nil
	} else {
		m.Space[y][x] = &a
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

const PrintWidth = 16

func (m *Memoryspace) Print() {
	fmt.Print("mem:    ")
	for j := 0; j < m.X(); j++ {
		fmt.Printf("|X=%-"+strconv.Itoa(PrintWidth-2)+"d", j)
	}
	fmt.Println()
	for i := 0; i < m.Y(); i++ {
		fmt.Printf("|Y=%-5d", i)
		for j := 0; j < m.X(); j++ {
			atom := m.At(j, i)
			var c log.Colortext
			switch atom.Type {
			case "int":
				c = log.Cyan
			case "bool":
				c = log.LightBlue
			case "rect":
				c = log.Blue
			case "point":
				c = log.Gray
			case "func":
				c = log.Yellow
			case "op":
				c = log.Red
			}
			log.Print(c, fmt.Sprintf("|%-"+strconv.Itoa(PrintWidth)+"s", atom.String()))
		}
		fmt.Println("|")
	}
	fmt.Println()
}
