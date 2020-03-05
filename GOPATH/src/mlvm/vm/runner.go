package vm

import (
	"mirror"
	"time"
)

type Runner struct {
	status       string
	x, y         int
	result       mirror.Atom
	pause        chan int
	Computecycle time.Duration
	memoryspace  *Memoryspace
}

func NewRunner(m *Memoryspace) *Runner {
	return &Runner{
		status:      "ready",
		pause:       make(chan int),
		memoryspace: m,
	}
}
func (r *Runner) Pause() {
	r.status = "pause"
}
func (r *Runner) Goon() {
	<-r.pause
	r.status = "run"
}
func (r *Runner) real(y, x int) (int, int) {
	a := r.memoryspace.Space[y][x]
	if a != nil {
		if a.Type == "point" {
			return a.Point_y, a.Point_x
		} else {
			return y, x
		}
	}
	return y, x
}

func (r *Runner) Do(startpoint mirror.Atom) {
	r.x, r.y = r.x, r.y
	for {
		if r.Computecycle > 0 {
			time.Sleep(r.Computecycle)
		}
		switch r.status {
		case "exit":
			return
		case "pause":
			r.pause <- 0
		}
		if r.y < len(r.memoryspace.Space) {
			cmd := r.memoryspace.Space[r.y][r.x]
			if cmd == nil {
				continue
			}
			switch cmd.Operator {
			// 单目运算
			case "free":
				y, x := r.real(r.y, r.x+1)
				r.memoryspace.Space[y][x] = nil
				r.y++
			case "=":
				y, x := r.real(r.y, r.x-1)
				r.memoryspace.Space[y][x] = &r.result
				y++
			case "!":
				y, x := r.real(r.y, r.x+1)
				r.memoryspace.Space[y][x].V_bool = !r.memoryspace.Space[y][x].V_bool
				r.y++
				//双运算

			case "+":
				fallthrough
			case "-":
				fallthrough
			case "*":
				fallthrough
			case "/":
				right_y, right_x := r.real(r.y, r.x+1)
				left_y, left_x := r.real(r.y, r.x-1)
				result := 0
				if cmd.Operator == "+" {
					result = r.memoryspace.Space[left_y][left_x].V_int - r.memoryspace.Space[right_y][right_x].V_int
				}
				if cmd.Operator == "-" {
					result = r.memoryspace.Space[left_y][left_x].V_int - r.memoryspace.Space[right_y][right_x].V_int
				}
				if cmd.Operator == "*" {
					result = r.memoryspace.Space[left_y][left_x].V_int * r.memoryspace.Space[right_y][right_x].V_int
				}
				if cmd.Operator == "/" {
					result = r.memoryspace.Space[left_y][left_x].V_int / r.memoryspace.Space[right_y][right_x].V_int
				}
				r.result = mirror.Atom{Type: "int", V_int: result}
			case "==":
				left := r.memoryspace.Space[r.y][r.x-1]
				var leftv *mirror.Atom
				switch left.Type {
				case "point":
					leftv = r.memoryspace.Space[left.Point_y][left.Point_x]
				default:
					leftv = left
				}
				right := r.memoryspace.Space[r.y][r.x+1]
				var rightv *mirror.Atom
				switch right.Type {
				case "point":
					rightv = r.memoryspace.Space[right.Point_y][right.Point_x]
				default:
					rightv = right
				}
				if leftv.Type == rightv.Type {
					switch left.Type {
					case "bool":
					}
				}
				r.y++
			case "go":
				right := r.memoryspace.Space[r.y][r.x+1]
				r.y += right.Point_y
				r.x += right.Point_x
			case "goto":
				right := r.memoryspace.Space[r.y][r.x+1]
				r.y, r.x = right.Point_y, right.Point_x
			}

		}

	}
}
