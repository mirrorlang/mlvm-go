package vm

import (
	"fmt"
	"mirror"
	"time"
)

type Runner struct {
	status       string
	X, Y         int
	Funcrect     mirror.FuncAtom
	result       mirror.Atom
	pause        chan int
	Computecycle time.Duration
	mem          *Memoryspace
}

func NewRunner(m *Memoryspace) *Runner {
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
	return c.mem.At(c.X+1, c.Y)
}
func (c *Runner) OpLeft() mirror.Atom {
	return c.mem.At(c.X-1, c.Y)
}
func (cpu *Runner) Next() {
	atom := cpu.mem.At(cpu.X, cpu.Y)
	var next mirror.Point
	switch atom.(type) {
	case mirror.OpAtom:
		next = atom.(mirror.OpAtom).Nextop
	case mirror.FuncAtom:
		next = atom.(mirror.FuncAtom).Nextop
	}
	if next.Isoffset {
		cpu.X += next.X
		cpu.Y += next.Y
	} else {
		cpu.X = next.X
		cpu.Y = next.Y
	}
}
func (cpu *Runner) Do() {

	for {
		if cpu.Computecycle > 0 {
			time.Sleep(cpu.Computecycle)
			if cpu.status != "idle" {
				cpu.mem.Print()
			}
		}
		switch cpu.status {
		case "exit":
			return
		case "idle":
		case "pause":
			cpu.pause <- 0
		}
		if cpu.Y < cpu.mem.Y() {
			atom := cpu.mem.At(cpu.X, cpu.Y)
			switch atom.Type() {
			case "op":
				op(cpu)
			case "controlflow":
				controlflow(cpu)
			default:
				cpu.status = "idle"
				continue
			}

		}
	}
}
