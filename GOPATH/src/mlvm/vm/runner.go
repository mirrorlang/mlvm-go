package vm

import (
	"fmt"
	"mirror"
	"time"
)

type Runner struct {
	status       string
	X, Y         int
	Funcrect     mirror.Atom
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
func (c *Runner) Op() mirror.Atom {
	return c.mem.At(c.X, c.Y)
}
func (c *Runner) Right() mirror.Atom {
	return c.mem.At(c.X+1, c.Y)
}
func (c *Runner) Left() mirror.Atom {
	return c.mem.At(c.X-1, c.Y)
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

//相对地址，转换绝对地址
func (m *Runner) Address(atom mirror.Atom) mirror.Atom {
	if atom.Type != "point" && atom.Type != "rect" && atom.Type != "rectdata" {
		panic("type !=point|rect|rectdata")
	}
	//有绝对地址
	if atom.X != 0 || atom.Y != 0 {
		return atom
	} else {
		atom.X = m.X
		atom.Y = m.Y
		return atom
	}
}
func (cpu *Runner) Do(startpoint mirror.Atom) {

	for {
		if cpu.Computecycle > 0 {
			time.Sleep(cpu.Computecycle)
		}
		switch cpu.status {
		case "exit":
			return
		case "pause":
			cpu.pause <- 0
		}
		if cpu.Y < cpu.mem.Y() {
			cmd := cpu.Op()
			if cmd.Type == "" {
				continue
			}
			switch cmd.Operator {
			// 单目运算
			case "nil":
				right := cpu.Right()
				t := right.Type
				switch t {
				case "rect":
					for i := right.Y; i < right.Size_y; i++ {
						for j := right.X; j < right.Size_x; j++ {
							cpu.mem.Set(right.X+right.Offset_x+j, right.Y+right.Offset_y+i, mirror.Atom{Type: "null"})
						}
					}
				case "point":
					right = cpu.Address(right)
					cpu.mem.Set(right.X+right.Offset_x, right.Y+right.Offset_y, mirror.Atom{Type: "null"})
				default:
					cpu.mem.Set(cpu.X+1, cpu.Y, mirror.Atom{Type: "null"})
				}
				cpu.Y++
			case "=":
				left := cpu.Left()
				v := cpu.result

				t := left.Type
				switch t {
				case "rectdata":
					left = cpu.Address(left)
					rect := cpu.mem.At(left.X+left.Offset_x, left.Y+left.Offset_y)
					cpu.mem.Set(rect.X+left.Rectoffset_x, rect.Y+left.Rectoffset_y, v)
				case "point":
					left = cpu.Address(left)
					cpu.mem.Set(left.X+left.Offset_x, left.Y+left.Offset_y, v)
				default:
					cpu.mem.Set(cpu.X-1, cpu.Y, v)
				}
				cpu.Y++
			case "!":
				right := cpu.Right()
				t := right.Type
				switch t {
				case "point":
					right = cpu.Address(right)
					updater := cpu.mem.At(right.X+right.Offset_x, right.Y+right.Offset_y)
					updater.V_bool = !updater.V_bool
					cpu.mem.Set(right.X+right.Offset_x, right.Y+right.Offset_y, updater)
				default:
					updater := cpu.mem.At(cpu.X+1, cpu.Y)
					updater.V_bool = !updater.V_bool
					cpu.mem.Set(cpu.X+1, cpu.Y, updater)
				}
				cpu.Y++
				//双运算
			case "+":
				fallthrough
			case "-":
				fallthrough
			case "*":
				fallthrough
			case "/":
				left := cpu.Left()
				if left.Type == "point" {
					left = cpu.Address(left)
					left = cpu.mem.space[left.Y+left.Offset_y][left.X+left.Offset_x].Clone()
				}
				right := cpu.mem.space[cpu.Y][cpu.X+1].Clone()
				if right.Type == "point" {
					right = cpu.Address(right)
					right = cpu.mem.space[right.Y+right.Offset_y][right.X+right.Offset_x].Clone()
				}
				result := 0
				if cmd.Operator == "+" {
					result = left.V_int + right.V_int
				}
				if cmd.Operator == "-" {
					result = left.V_int - right.V_int
				}
				if cmd.Operator == "*" {
					result = left.V_int * right.V_int
				}
				if cmd.Operator == "/" {
					result = left.V_int / right.V_int
				}
				cpu.result = mirror.Atom{Type: "int", V_int: result}
				cpu.Y++
			case "==":

				//todo
				cpu.Y++
			case "rect":
				var left mirror.Atom
				if cpu.X > 1 {
					left = cpu.mem.space[cpu.Y][cpu.X-1].Clone()
				}
				right := cpu.mem.space[cpu.Y][cpu.X+1].Clone()
				if left.Type == "" {
					cpu.Funcrect.Size_x = right.Size_x
					cpu.Funcrect.Size_y = right.Size_y
				} else {
					cpu.mem.space[cpu.Y][cpu.X-1].Size_x = right.Size_x
					cpu.mem.space[cpu.Y][cpu.X-1].Size_y = right.Size_y
				}
				cpu.Y++
			case "go":

				right := cpu.mem.space[cpu.Y][cpu.X+1].Clone()

				cpu.Y += right.Offset_y
				cpu.X += right.Offset_x

			case "goto":
				left := cpu.mem.space[cpu.Y][cpu.X-1].Clone()
				right := cpu.mem.space[cpu.Y][cpu.X+1].Clone()
				if left.Type == "" {
					cpu.Y, cpu.X = right.Y, right.X
				} else {
					cpu.mem.space[cpu.Y][cpu.X-1].Y, cpu.mem.space[cpu.Y][cpu.X-1].X = right.Y, right.X
					cpu.Y++
				}

			case "call": //函数调用,这里必须要知道函数体的rect
				nextfuncp := cpu.mem.space[cpu.Y+1][cpu.X].Clone()
				nextfuncrect := cpu.mem.space[nextfuncp.Y][nextfuncp.X].Clone()
				nextfuncrect.Y, nextfuncrect.X = nextfuncp.Y, nextfuncp.X
				for i := 0; i < nextfuncrect.Size_y; i++ {
					for j := 0; j < nextfuncrect.Size_x; j++ {
						sourceatom := cpu.mem.space[nextfuncrect.Y+nextfuncrect.Offset_y+i][nextfuncrect.X+nextfuncrect.Offset_x+j].Clone()
						cpu.mem.space[cpu.Y+i][cpu.Funcrect.X+cpu.Funcrect.Size_x+j] = &sourceatom
					}
				}
				argsrect := cpu.mem.space[cpu.Y+2][cpu.X].Clone()
				//参数rect的start.X,Y 使用相对地址
				for i := 0; i < argsrect.Size_y; i++ {
					for j := 0; j < argsrect.Size_x; j++ {
						sourceatom := cpu.mem.space[cpu.Y+argsrect.Y+argsrect.Offset_y+i][cpu.X+argsrect.X+argsrect.Offset_x+j].Clone()
						cpu.mem.space[cpu.Y+i][cpu.Funcrect.X+cpu.Funcrect.Size_x+1+j] = &sourceatom
					}
				}

				//缓存函数体的调用者caller地址
				oldfuncrect := cpu.Funcrect
				cpu.mem.space[cpu.Y+1][cpu.Funcrect.X+cpu.Funcrect.Size_x] = &oldfuncrect

				//修改函数体的返回值地址,这里必须是绝对地址
				returnrect := cpu.mem.space[cpu.Y+3][cpu.X].Clone()
				returnrect.X = cpu.X + returnrect.Offset_x
				returnrect.Y = cpu.Y + returnrect.Offset_y
				returnrect.Offset_x = 0
				returnrect.Offset_y = 0
				cpu.mem.space[cpu.Y+3][cpu.Funcrect.X+cpu.Funcrect.Size_x] = &returnrect

				//缓存函数体的调用者caller地址
				cpu.mem.space[cpu.Y+4][cpu.Funcrect.X+cpu.Funcrect.Size_x] = &mirror.Atom{Type: "point", X: cpu.X, Y: cpu.Y}

				//移动至函数
				//这里的funcrect pointxy必须是绝对坐标

				cpu.Funcrect = nextfuncrect.Clone()
				cpu.Funcrect.X = cpu.Funcrect.X + oldfuncrect.Size_x
				cpu.Funcrect.Y = cpu.Y

				//进入新的funcrect执行体，平移过去
				cpu.X = cpu.Funcrect.X

				//进入函数第一个操作

				firstop := cpu.mem.space[cpu.Y+2][cpu.X].Clone()
				cpu.X += firstop.Offset_x
				cpu.Y += firstop.Offset_y

			case "return":
				//回收执行现场
				Y, X := cpu.mem.space[cpu.Funcrect.Y+4][cpu.Funcrect.X].Y+4, cpu.mem.space[cpu.Funcrect.Y+4][cpu.Funcrect.X].X

				CallerRect := *cpu.mem.space[cpu.Funcrect.Y+1][cpu.Funcrect.X]
				for i := 0; i < cpu.Funcrect.Size_y; i++ {
					for j := 0; j < cpu.Funcrect.Size_x; j++ {
						cpu.mem.space[cpu.Funcrect.Y+i][cpu.Funcrect.X+j] = nil
					}
				}
				//return
				cpu.Y, cpu.X = Y, X
				cpu.Funcrect = CallerRect
			}
		}
	}
}
