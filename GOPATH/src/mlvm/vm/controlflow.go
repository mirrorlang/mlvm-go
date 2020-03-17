package vm

import "mirror"

func controlflow(cpu *Runner) {
	switch cpu.mem.At(cpu.X, cpu.Y).(type) {
	case *mirror.GotoAtom:
		right := cpu.OpRight().(*mirror.PointAtom)
		if right.Isoffset {
			cpu.X += right.X
			cpu.Y += right.Y
		} else {
			cpu.X = right.X
			cpu.Y = right.Y
		}
	case *mirror.IfAtom:
		ifat := cpu.mem.At(cpu.X, cpu.Y).(*mirror.IfAtom)
		b := cpu.OpRight().(*mirror.BoolAtom)
		if b.Value {
			if ifat.True.Isoffset {
				cpu.X += ifat.True.X
				cpu.Y += ifat.True.Y
			} else {
				cpu.X = ifat.True.X
				cpu.Y = ifat.True.Y
			}
		} else {
			if ifat.Else.Isoffset {
				cpu.X += ifat.Else.X
				cpu.Y += ifat.Else.Y
			} else {
				cpu.X = ifat.Else.X
				cpu.Y = ifat.Else.Y
			}
		}
	case *mirror.SwitchAtom:
	}
}
