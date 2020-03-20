package vm

import (
	"fmt"
	"mirror"
	"mlvm/vm/mem"
	"time"
)

type Runner struct {
	status string

	Runfunc      *mirror.FuncAtom
	resulttype   string //atom、array、rect
	result       mirror.Atom
	pause        chan int
	Computecycle time.Duration
	mem          *mem.Memoryspace
}

func (c *Runner) Exception() {
	c.status = "exception"
}
func (c *Runner) At(x, y int) mirror.Atom {
	return c.mem.At(x, y)
}
func (c *Runner) Set(x, y int, a mirror.Atom) {
	c.mem.Set(x, y, a)
}

func (c *Runner) InFunc(x, y int) mirror.Atom {
	if x > c.Runfunc.Funcbody.Size_x || y > c.Runfunc.Funcbody.Size_y {
		c.Exception()
	}
	return c.At(c.Runfunc.Funcbody.X+x, c.Runfunc.Funcbody.Y+y)
}
func NewRunner(m *mem.Memoryspace) *Runner {
	return &Runner{
		status: "ready",
		pause:  make(chan int),
		mem:    m,
	}
}
func (c *Runner) Print() {
	fmt.Printf("cpu %+v", c)
	fmt.Println()
}
func (r *Runner) Pause() {
	r.status = "pause"
}
func (r *Runner) Goon() {
	<-r.pause
	r.status = "run"
}

func (c *Runner) OpRight() mirror.Atom {
	if c.At(c.Runfunc.Cpu_x+1, c.Runfunc.Cpu_y) != nil {
		return c.At(c.Runfunc.Cpu_x+1, c.Runfunc.Cpu_y)
	} else {
		return &mirror.NullAtom{}
	}

}
func (c *Runner) OpLeft() mirror.Atom {
	if c.At(c.Runfunc.Cpu_x-1, c.Runfunc.Cpu_y) != nil {
		return c.At(c.Runfunc.Cpu_x-1, c.Runfunc.Cpu_y)
	} else {
		return &mirror.NullAtom{}
	}
}
func (c *Runner) Next() {
	atom := c.At(c.Runfunc.Cpu_x, c.Runfunc.Cpu_y)
	var next = mirror.Point{X: 0, Y: 1, Isoffset: true}
	switch atom.(type) {
	case *mirror.OpAtom:
		next = atom.(*mirror.OpAtom).Nextop
	case *mirror.FuncAtom:
		next = atom.(*mirror.FuncAtom).Nextop
	case *mirror.CallAtom:
		next = atom.(*mirror.CallAtom).Nextop
	}
	if next.Isoffset {
		c.Runfunc.Cpu_x += next.X
		c.Runfunc.Cpu_y += next.Y
	} else {
		c.Runfunc.Cpu_x = next.X
		c.Runfunc.Cpu_y = next.Y
	}
}

func (c *Runner) Do(x, y int) {
	switch c.At(x, y).(type) {
	case *mirror.FuncAtom:
		c.Runfunc = c.At(x, y).(*mirror.FuncAtom)
	default:
		c.Runfunc = &mirror.FuncAtom{
			Cpu_x: x,
			Cpu_y: y,
		}
	}
	c.Runfunc.Funcbody.X, c.Runfunc.Funcbody.Y = x, y

	for {
		if c.Computecycle > 0 {
			time.Sleep(c.Computecycle)
			if c.status != "idle" {
				c.mem.Print()
			}
		}
		switch c.status {
		case "exit":
			return
		case "idle":
		case "pause":
		case "exception":
			c.pause <- 0
		}

		atom := c.At(c.Runfunc.Cpu_x, c.Runfunc.Cpu_y)
		if atom != nil {
			switch atom.Type() {
			case "op":
				op(c)
			case "func":
				funcc(c)
			case "controlflow":
				controlflow(c)
			default:
				c.status = "idle"
				continue
			}
		} else {
			c.status = "idle"
			continue
		}

	}
}
