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

const Width = 16

func (m *Memoryspace) Print() {
	fmt.Print("mem:    ")
	for j := 0; j < len(m.Space[0]); j++ {
		fmt.Printf("|x=%-"+strconv.Itoa(Width-2)+"d", j)
	}
	fmt.Println()
	for i := 0; i < len(m.Space); i++ {
		fmt.Printf("|y=%-5d", i)
		for j := 0; j < len(m.Space[i]); j++ {
			atom := m.Space[i][j]
			if atom == nil {
				fmt.Printf("|%"+strconv.Itoa(Width)+"s", "")
				continue
			}
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
			log.Print(c, fmt.Sprintf("|%-"+strconv.Itoa(Width)+"s", atom.String()))
		}
		fmt.Println("|")
	}
	fmt.Println()
}
