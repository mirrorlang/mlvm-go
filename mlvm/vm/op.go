package vm

import (
	"mlvm_go/mirror"
)

func op(cpu *Runner) {
	atom := cpu.At(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y).(*mirror.OpAtom)
	switch atom.Op {
	// 单目运算
	case "nil":
		switch t := cpu.OpRight().Type(); t {
		case "rect":
			right := cpu.OpRight().(*mirror.RectAtom)
			for i := right.Y; i < right.Size_y; i++ {
				for j := right.X; j < right.Size_x; j++ {
					cpu.Set(right.X+j, right.Y+i, nil)
				}
			}
		case "point":
			right := cpu.OpRight().(*mirror.PointAtom)
			x, y := right.GlobalAddr(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y)
			cpu.Set(x, y, nil)
		default:
			cpu.Set(cpu.Runfunc.Cpu_x+1, cpu.Runfunc.Cpu_y, nil)
		}
	case "=":
		switch t := cpu.OpLeft().Type(); t {
		case "rectpoint":
			left := cpu.OpLeft().(*mirror.CascadeAtom)
			rect := cpu.At(left.GlobalAddr(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y)).(*mirror.RectAtom)
			cpu.Set(rect.X+left.Inrect_offset_x, rect.Y+left.Inrect_offset_y, cpu.result)
		case "point":
			left := cpu.OpLeft().(*mirror.PointAtom)
			x, y := left.GlobalAddr(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y)
			cpu.Set(x, y, cpu.result)
		default:
			cpu.Set(cpu.Runfunc.Cpu_x-1, cpu.Runfunc.Cpu_y, cpu.result)
		}
	case "!":
		switch t := cpu.OpLeft().Type(); t {
		case "point":
			right := cpu.OpRight().(*mirror.PointAtom)
			updater := cpu.At(right.GlobalAddr(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y)).(*mirror.BoolAtom)
			updater.Value = !updater.Value
			x, y := right.GlobalAddr(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y)
			cpu.Set(x, y, updater)
		default:
			updater := cpu.At(cpu.Runfunc.Cpu_x+1, cpu.Runfunc.Cpu_y).(*mirror.BoolAtom)
			updater.Value = !updater.Value
			cpu.Set(cpu.Runfunc.Cpu_x+1, cpu.Runfunc.Cpu_y, updater)
		}
		//双运算
	case "+":
		fallthrough
	case "-":
		var left, right mirror.Atom
		if cpu.OpLeft().Type() == "point" {
			left = cpu.At(cpu.OpLeft().(*mirror.PointAtom).GlobalAddr(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y))
		} else {
			left = cpu.OpLeft()
		}
		if cpu.OpRight().Type() == "point" {
			right = cpu.At(cpu.OpRight().(*mirror.PointAtom).GlobalAddr(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y))
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
			left = cpu.At(cpu.OpLeft().(*mirror.PointAtom).GlobalAddr(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y))
		} else {
			left = cpu.OpLeft()
		}
		if cpu.OpRight().Type() == "point" {
			right = cpu.At(cpu.OpRight().(*mirror.PointAtom).GlobalAddr(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y))
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
	}
	cpu.Next()
}
