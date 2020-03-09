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
	if atom.Point_x != 0 || atom.Point_y != 0 {
		return atom
	} else {
		atom.Point_x = m.X
		atom.Point_y = m.Y
		return atom
	}
}
func (cpu *Runner) Do(startpoint mirror.Atom) {
	cpu.X, cpu.Y = cpu.X, cpu.Y
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
		if cpu.Y < len(cpu.mem.Space) {
			cmd := cpu.mem.Space[cpu.Y][cpu.X].Clone()
			if cmd.Type == "" {
				continue
			}
			switch cmd.Operator {
			// 单目运算
			case "nil":
				right := cpu.mem.Space[cpu.Y][cpu.X+1].Clone()
				t := right.Type
				switch t {
				case "rect":
					for i := right.Point_y; i < right.Size_y; i++ {
						for j := right.Point_x; j < right.Size_x; j++ {
							cpu.mem.Space[right.Point_y+right.Offset_y+i][right.Point_x+right.Offset_x+j] = nil
						}
					}
				case "point":
					right = cpu.Address(right)
					cpu.mem.Space[right.Point_y+right.Offset_y][right.Point_x+right.Offset_x] = nil
				default:
					cpu.mem.Space[cpu.Y][cpu.X+1] = nil
				}
				cpu.Y++
			case "empty":
				right := cpu.mem.Space[cpu.Y][cpu.X+1].Clone()
				t := right.Type
				switch t {
				case "rect":
					for i := right.Point_y; i < right.Size_y; i++ {
						for j := right.Point_x; j < right.Size_x; j++ {
							cpu.mem.Space[right.Point_y+right.Offset_y+i][right.Point_x+right.Offset_x+j] = &mirror.Atom{Type: t}
						}
					}
				case "point":
					right = cpu.Address(right)
					cpu.mem.Space[right.Point_y+right.Offset_y][right.Point_x+right.Offset_x] = &mirror.Atom{Type: cpu.mem.Space[right.Point_y][right.Point_x].Type}
				default:
					cpu.mem.Space[cpu.Y][cpu.X+1] = nil
				}
				cpu.Y++
			case "=":
				left := cpu.mem.Space[cpu.Y][cpu.X-1].Clone()

				v := cpu.result.Clone()
				if left.Type == "" {
					cpu.mem.Space[cpu.Y][cpu.X-1] = &v
				} else {
					t := left.Type
					switch t {
					case "rectdata":
						left = cpu.Address(left)
						rect := cpu.mem.Space[left.Point_y+left.Offset_y][left.Point_x+left.Offset_x]
						cpu.mem.Space[rect.Point_y+left.Rectoffset_y][rect.Point_x+left.Rectoffset_x] = &v
					case "point":
						left = cpu.Address(left)
						cpu.mem.Space[left.Point_y+left.Offset_y][left.Point_x+left.Offset_x] = &v
					default:
						cpu.mem.Space[cpu.Y][cpu.X-1] = &v
					}
				}

				cpu.Y++
			case "!":
				right := cpu.mem.Space[cpu.Y][cpu.X+1].Clone()
				t := right.Type
				switch t {
				case "point":
					right = cpu.Address(right)
					cpu.mem.Space[right.Point_y+right.Offset_y][right.Point_x+right.Offset_x].V_bool = !cpu.mem.Space[right.Point_y+right.Offset_y][right.Point_x+right.Offset_x].V_bool
				default:
					cpu.mem.Space[cpu.Y][cpu.X+1].V_bool = !cpu.mem.Space[cpu.Y][cpu.X+1].V_bool
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
				left := cpu.mem.Space[cpu.Y][cpu.X-1].Clone()
				if left.Type == "point" {
					left = cpu.Address(left)
					left = cpu.mem.Space[left.Point_y+left.Offset_y][left.Point_x+left.Offset_x].Clone()
				}
				right := cpu.mem.Space[cpu.Y][cpu.X+1].Clone()
				if right.Type == "point" {
					right = cpu.Address(right)
					right = cpu.mem.Space[right.Point_y+right.Offset_y][right.Point_x+right.Offset_x].Clone()
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
					left = cpu.mem.Space[cpu.Y][cpu.X-1].Clone()
				}
				right := cpu.mem.Space[cpu.Y][cpu.X+1].Clone()
				if left.Type == "" {
					cpu.Funcrect.Size_x = right.Size_x
					cpu.Funcrect.Size_y = right.Size_y
				} else {
					cpu.mem.Space[cpu.Y][cpu.X-1].Size_x = right.Size_x
					cpu.mem.Space[cpu.Y][cpu.X-1].Size_y = right.Size_y
				}
				cpu.Y++
			case "go":

				right := cpu.mem.Space[cpu.Y][cpu.X+1].Clone()

				cpu.Y += right.Offset_y
				cpu.X += right.Offset_x

			case "goto":
				left := cpu.mem.Space[cpu.Y][cpu.X-1].Clone()
				right := cpu.mem.Space[cpu.Y][cpu.X+1].Clone()
				if left.Type == "" {
					cpu.Y, cpu.X = right.Point_y, right.Point_x
				} else {
					cpu.mem.Space[cpu.Y][cpu.X-1].Point_y, cpu.mem.Space[cpu.Y][cpu.X-1].Point_x = right.Point_y, right.Point_x
					cpu.Y++
				}

			case "call": //函数调用,这里必须要知道函数体的rect
				nextfuncp := cpu.mem.Space[cpu.Y+1][cpu.X].Clone()
				nextfuncrect := cpu.mem.Space[nextfuncp.Point_y][nextfuncp.Point_x].Clone()
				nextfuncrect.Point_y, nextfuncrect.Point_x = nextfuncp.Point_y, nextfuncp.Point_x
				cpu.Print()
				for i := 0; i < nextfuncrect.Size_y; i++ {
					for j := 0; j < nextfuncrect.Size_x; j++ {
						sourceatom := cpu.mem.Space[nextfuncrect.Point_y+nextfuncrect.Offset_y+i][nextfuncrect.Point_x+nextfuncrect.Offset_x+j].Clone()
						cpu.mem.Space[cpu.Y+i][cpu.Funcrect.Point_x+cpu.Funcrect.Size_x+j] = &sourceatom
					}
				}
				cpu.mem.Print()
				argsrect := cpu.mem.Space[cpu.Y+2][cpu.X].Clone()
				//参数rect的start.X,Y 使用相对地址
				for i := 0; i < argsrect.Size_y; i++ {
					for j := 0; j < argsrect.Size_x; j++ {
						sourceatom := cpu.mem.Space[cpu.Y+argsrect.Point_y+argsrect.Offset_y+i][cpu.X+argsrect.Point_x+argsrect.Offset_x+j].Clone()
						cpu.mem.Space[cpu.Y+i][cpu.Funcrect.Point_x+cpu.Funcrect.Size_x+1+j] = &sourceatom
					}
				}
				cpu.mem.Print()

				//缓存函数体的调用者caller地址
				oldfuncrect := cpu.Funcrect
				cpu.mem.Space[cpu.Y+1][cpu.Funcrect.Point_x+cpu.Funcrect.Size_x] = &oldfuncrect

				//修改函数体的返回值地址,这里必须是绝对地址
				returnrect := cpu.mem.Space[cpu.Y+3][cpu.X].Clone()
				returnrect.Point_x = cpu.X + returnrect.Offset_x
				returnrect.Point_y = cpu.Y + returnrect.Offset_y
				returnrect.Offset_x = 0
				returnrect.Offset_y = 0
				cpu.mem.Space[cpu.Y+3][cpu.Funcrect.Point_x+cpu.Funcrect.Size_x] = &returnrect

				//缓存函数体的调用者caller地址
				cpu.mem.Space[cpu.Y+4][cpu.Funcrect.Point_x+cpu.Funcrect.Size_x] = &mirror.Atom{Type: "point", Point_x: cpu.X, Point_y: cpu.Y}
				cpu.mem.Print()
				//移动至函数
				//这里的funcrect pointxy必须是绝对坐标

				cpu.Funcrect = nextfuncrect.Clone()
				cpu.Funcrect.Point_x = cpu.Funcrect.Point_x + oldfuncrect.Size_x
				cpu.Funcrect.Point_y = cpu.Y

				//进入新的funcrect执行体，平移过去
				cpu.X = cpu.Funcrect.Point_x
				cpu.Print()
				//进入函数第一个操作

				firstop := cpu.mem.Space[cpu.Y+2][cpu.X].Clone()
				cpu.X += firstop.Offset_x
				cpu.Y += firstop.Offset_y

				cpu.mem.Print()
				cpu.Print()

			case "return":
				//回收执行现场
				Y, X := cpu.mem.Space[cpu.Funcrect.Point_y+4][cpu.Funcrect.Point_x].Point_y+4, cpu.mem.Space[cpu.Funcrect.Point_y+4][cpu.Funcrect.Point_x].Point_x

				CallerRect := *cpu.mem.Space[cpu.Funcrect.Point_y+1][cpu.Funcrect.Point_x]
				for i := 0; i < cpu.Funcrect.Size_y; i++ {
					for j := 0; j < cpu.Funcrect.Size_x; j++ {
						cpu.mem.Space[cpu.Funcrect.Point_y+i][cpu.Funcrect.Point_x+j] = nil
					}
				}
				//return
				cpu.Y, cpu.X = Y, X
				cpu.Funcrect = CallerRect
			}
		}
	}
}
