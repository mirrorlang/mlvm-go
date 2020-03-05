package vm

import (
	"mirror"
)

type Runner struct {
	status string
	pause  chan int
}

func NewRunner() *Runner {
	return &Runner{
		status: "ready",
	}
}
func (r *Runner) Pause() {
	r.status = "pause"
}
func (r *Runner) Goon() {
	<-r.pause
}
func (r *Runner) Do(memoryspace *Memoryspace, startpoint mirror.Atom) {
	for {
		switch r.status {
		case "exit":
			return
		case "pause":
			r.pause <- 0
		}
		cmd := memoryspace.Space[startpoint.Point_x][startpoint.Point_y]
		switch cmd.Operator {
		case "null":
			right := memoryspace.Space[startpoint.Point_x+1][startpoint.Point_y]
			switch right.Type {
			case "point":
				memoryspace.Space[right.Point_x][right.Point_y] = mirror.Atom{Type: "null"}
			default:
				memoryspace.Space[startpoint.Point_x+1][startpoint.Point_y] = mirror.Atom{Type: "null"}
			}

		case "!":
			right := memoryspace.Space[startpoint.Point_x+1][startpoint.Point_y]
			switch right.Type {
			case "point":
				memoryspace.Space[right.Point_x][right.Point_y].V_bool = !memoryspace.Space[right.Point_x][right.Point_y].V_bool
			default:
				memoryspace.Space[startpoint.Point_x+1][startpoint.Point_y].V_bool = !memoryspace.Space[startpoint.Point_x+1][startpoint.Point_y].V_bool
			}
		case "==":
			left := memoryspace.Space[startpoint.Point_x-1][startpoint.Point_y]
			var leftv mirror.Atom
			switch left.Type {
			case "point":
				leftv = memoryspace.Space[left.Point_x][left.Point_y]
			default:
				leftv = left
			}
			right := memoryspace.Space[startpoint.Point_x+1][startpoint.Point_y]
			var rightv mirror.Atom
			switch right.Type {
			case "point":
				rightv = memoryspace.Space[right.Point_x][right.Point_y]
			default:
				rightv = right
			}
			if leftv.Type == rightv.Type {
				switch left.Type {
				case "bool":
				}
			}
		}
	}
}
