package vm

import (
	"mirror"
)

type Runner struct {
	status string
	x, y   int
	pause  chan int
}

func NewRunner() *Runner {
	return &Runner{
		status: "ready",
		pause:  make(chan int),
	}
}
func (r *Runner) Pause() {
	r.status = "pause"
}
func (r *Runner) Goon() {
	<-r.pause
	r.status = "run"
}
func (r *Runner) Do(memoryspace *Memoryspace, startpoint mirror.Atom) {
	r.x, r.y = r.x, r.y
	for {
		switch r.status {
		case "exit":
			return
		case "pause":
			r.pause <- 0
		}
		cmd := memoryspace.Space[r.y][r.x]
		if cmd == nil {
			r.x, r.y = 0, 0
			continue
		}
		switch cmd.Operator {
		case "null":
			right := memoryspace.Space[r.y][r.x+1]
			if right != nil {
				if right.Type == "point" {
					memoryspace.Space[right.Point_y][right.Point_x] = nil
					break
				}
				memoryspace.Space[r.y][r.x+1] = nil
			}

		case "!":
			right := memoryspace.Space[r.y][r.x+1]
			switch right.Type {
			case "point":
				memoryspace.Space[right.Point_y][right.Point_x].V_bool = !memoryspace.Space[right.Point_y][right.Point_x].V_bool
			default:
				memoryspace.Space[r.y][r.x+1].V_bool = !memoryspace.Space[r.y][r.x+1].V_bool
			}
		case "==":
			left := memoryspace.Space[r.y][r.x-1]
			var leftv *mirror.Atom
			switch left.Type {
			case "point":
				leftv = memoryspace.Space[left.Point_y][left.Point_x]
			default:
				leftv = left
			}
			right := memoryspace.Space[r.y][r.x+1]
			var rightv *mirror.Atom
			switch right.Type {
			case "point":
				rightv = memoryspace.Space[right.Point_y][right.Point_x]
			default:
				rightv = right
			}
			if leftv.Type == rightv.Type {
				switch left.Type {
				case "bool":
				}
			}
		}
		r.y++
	}
}
