package vm

import (
	"mirror"
)

func op(cpu *Runner) {
	atom := cpu.mem.At(cpu.X, cpu.Y).(*mirror.OpAtom)
	switch atom.Op {
	// 单目运算
	case "nil":
		switch t := cpu.OpRight().Type(); t {
		case "rect":
			right := cpu.OpRight().(*mirror.RectAtom)
			for i := right.Y; i < right.Size_y; i++ {
				for j := right.X; j < right.Size_x; j++ {
					cpu.mem.Set(right.X+j, right.Y+i, nil)
				}
			}
		case "point":
			right := cpu.OpRight().(*mirror.PointAtom)
			x, y := right.GlobalAddr(cpu.X, cpu.Y)
			cpu.mem.Set(x, y, nil)
		default:
			cpu.mem.Set(cpu.X+1, cpu.Y, nil)
		}
		cpu.Next()
	case "=":
		switch t := cpu.OpLeft().Type(); t {
		case "rectpoint":
			left := cpu.OpLeft().(*mirror.RectPointAtom)
			rect := cpu.mem.At(left.GlobalAddr(cpu.X, cpu.Y)).(*mirror.RectAtom)
			cpu.mem.Set(rect.X+left.Inrect_offset_x, rect.Y+left.Inrect_offset_y, cpu.result)
		case "point":
			left := cpu.OpLeft().(*mirror.PointAtom)
			x, y := left.GlobalAddr(cpu.X, cpu.Y)
			cpu.mem.Set(x, y, cpu.result)
		default:
			cpu.mem.Set(cpu.X-1, cpu.Y, cpu.result)
		}
	case "!":
		switch t := cpu.OpLeft().Type(); t {
		case "point":
			right := cpu.OpRight().(*mirror.PointAtom)
			updater := cpu.mem.At(right.GlobalAddr(cpu.X, cpu.Y)).(mirror.BoolAtom)
			updater.Value = !updater.Value
			x, y := right.GlobalAddr(cpu.X, cpu.Y)
			cpu.mem.Set(x, y, updater)
		default:
			updater := cpu.mem.At(cpu.X+1, cpu.Y).(mirror.BoolAtom)
			updater.Value = !updater.Value
			cpu.mem.Set(cpu.X+1, cpu.Y, updater)
		}
		//双运算
	case "+":
		fallthrough
	case "-":
		var left, right mirror.Atom
		if cpu.OpLeft().Type() == "point" {
			left = cpu.mem.At(cpu.OpLeft().(*mirror.PointAtom).GlobalAddr(cpu.X, cpu.Y))
		} else {
			left = cpu.OpLeft()
		}
		if cpu.OpRight().Type() == "point" {
			left = cpu.mem.At(cpu.OpRight().(*mirror.PointAtom).GlobalAddr(cpu.X, cpu.Y))
		}
		result := 0
		//todo 指针 运算
		switch atom.Op {
		case "+":
			result = (left.(*mirror.NumAtom)).IntValue + (right.(*mirror.NumAtom)).IntValue
		case "-":
			result = (left.(*mirror.NumAtom)).IntValue - (right.(*mirror.NumAtom)).IntValue
		}
		cpu.result = mirror.NumAtom{IntValue: result}
	case "*":
		fallthrough
	case "/":
		var left, right mirror.Atom
		if cpu.OpLeft().Type() == "point" {
			left = cpu.mem.At(cpu.OpLeft().(*mirror.PointAtom).GlobalAddr(cpu.X, cpu.Y))
		} else {
			left = cpu.OpLeft()
		}
		if cpu.OpRight().Type() == "point" {
			left = cpu.mem.At(cpu.OpRight().(*mirror.PointAtom).GlobalAddr(cpu.X, cpu.Y))
		}
		result := 0

		switch atom.Op {
		case "*":
			result = (left.(*mirror.NumAtom)).IntValue * (right.(*mirror.NumAtom)).IntValue
		case "/":
			result = (left.(*mirror.NumAtom)).IntValue / (right.(*mirror.NumAtom)).IntValue
		}
		cpu.result = &mirror.NumAtom{IntValue: result}
	case "==":
		//todo
	case "rect":
		//var left mirror.Atom
		//if cpu.X > 1 {
		//	left = cpu.mem.At(cpu.X-1, cpu.Y)
		//}
		//right := cpu.mem.At(cpu.X+1, cpu.Y)
		//if left.Type == "" {
		//	cpu.Funcrect.Size_x = right.Size_x
		//	cpu.Funcrect.Size_y = right.Size_y
		//} else {
		//	left.Size_x = right.Size_x
		//	left.Size_y = right.Size_y
		//	cpu.mem.Set(cpu.X-1, cpu.Y, left)
		//}
		//cpu.Y++

	case "call": //函数调用,这里必须要知道函数体的rect
		funcp := cpu.mem.At(cpu.X, cpu.Y+1).(*mirror.PointAtom)

		funcrect := cpu.mem.At(funcp.X, funcp.Y).(*mirror.FuncAtom) //方法区的地址
		nextrunningfuncrect := funcrect                             //新执行现场
		funcrect.Y, funcrect.X = funcp.Y, funcp.X                   //方法区的绝对地址,一开始这个地址是不知道的

		nextrunningfuncrect.Y, nextrunningfuncrect.X = cpu.Y, cpu.Funcrect.X+cpu.Funcrect.Size_x //新执行现场的地址

		for i := 0; i < funcrect.Size_y; i++ {
			for j := 0; j < funcrect.Size_x; j++ {
				sourceatom := cpu.mem.At(funcrect.X+j, funcrect.Y+i)
				cpu.mem.Set(nextrunningfuncrect.X+j, nextrunningfuncrect.Y+i, sourceatom)
			}
		}
		argsrect := cpu.mem.At(cpu.X, cpu.Y+2).(*mirror.RectAtom)
		//参数rect的start.X,Y 使用相对地址
		for i := 0; i < argsrect.Size_y; i++ {
			for j := 0; j < argsrect.Size_x; j++ {
				sourceatom := cpu.mem.At(cpu.X+argsrect.X+j, cpu.Y+argsrect.Y+i)
				cpu.mem.Set(nextrunningfuncrect.X+1+j, nextrunningfuncrect.Y+i, sourceatom) //参数区是在函数体的x+1位置
			}
		}

		//缓存函数体的调用者caller地址
		oldfuncrect := cpu.Funcrect
		cpu.mem.Set(cpu.Funcrect.X+cpu.Funcrect.Size_x, cpu.Y+1, oldfuncrect)

		//修改函数体的返回值地址,这里必须是绝对地址
		returnrect := cpu.mem.At(cpu.X, cpu.Y+3).(*mirror.FuncAtom)
		cpu.mem.Set(cpu.Funcrect.X+cpu.Funcrect.Size_x, cpu.Y+3, returnrect)

		//缓存函数体的调用者caller地址
		cpu.mem.Set(cpu.Funcrect.X+cpu.Funcrect.Size_x, cpu.Y+4, &mirror.PointAtom{Point: mirror.Point{X: cpu.X, Y: cpu.Y}})

		//移动至函数
		//这里的funcrect pointxy必须是绝对坐标

		cpu.Funcrect = *nextrunningfuncrect
		//进入新的funcrect执行体，平移过去
		cpu.X = nextrunningfuncrect.X

		//进入函数第一个操作
		cpu.Next()
	case "return":
		//回收执行现场

		caller := cpu.mem.At(cpu.Funcrect.X, cpu.Funcrect.Y+4).(*mirror.PointAtom)
		Y, X := caller.Y+4, caller.X

		CallerRect := cpu.mem.At(cpu.Funcrect.X, cpu.Funcrect.Y+1).(*mirror.FuncAtom)
		for i := 0; i < cpu.Funcrect.Size_y; i++ {
			for j := 0; j < cpu.Funcrect.Size_x; j++ {
				cpu.mem.Set(cpu.Funcrect.X+j, cpu.Funcrect.Y+i, nil)
			}
		}
		//return
		cpu.Y, cpu.X = Y, X
		cpu.Funcrect = *CallerRect
	}
}
