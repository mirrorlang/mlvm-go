package vm

import "mirror"

func controlflow(cpu *Runner) {
	switch cpu.At(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y).(type) {
	case *mirror.GotoAtom:
		right := cpu.OpRight().(*mirror.PointAtom)
		if right.Isoffset {
			cpu.Runfunc.Cpu_x += right.X
			cpu.Runfunc.Cpu_y += right.Y
		} else {
			cpu.Runfunc.Cpu_x = right.X
			cpu.Runfunc.Cpu_y = right.Y
		}
	case *mirror.IfAtom:
		ifat := cpu.At(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y).(*mirror.IfAtom)
		b := cpu.OpRight().(*mirror.BoolAtom)
		if b.Value {
			if ifat.True.Isoffset {
				cpu.Runfunc.Cpu_x += ifat.True.X
				cpu.Runfunc.Cpu_y += ifat.True.Y
			} else {
				cpu.Runfunc.Cpu_x = ifat.True.X
				cpu.Runfunc.Cpu_y = ifat.True.Y
			}
		} else {
			if ifat.Else.Isoffset {
				cpu.Runfunc.Cpu_x += ifat.Else.X
				cpu.Runfunc.Cpu_y += ifat.Else.Y
			} else {
				cpu.Runfunc.Cpu_x = ifat.Else.X
				cpu.Runfunc.Cpu_y = ifat.Else.Y
			}
		}
	case *mirror.SwitchAtom:
	}
}
