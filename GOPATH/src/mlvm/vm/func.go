package vm

import "mirror"

func funcc(cpu *Runner) {

	switch cpu.At(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y).(type) {

	case *mirror.FuncAtom: //函数调用,这里必须要知道函数体的rect
		funcrect := cpu.At(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y).(*mirror.FuncAtom)
		cpu.Runfunc = funcrect
		cpu.Next()
	case *mirror.CallAtom: //函数调用,这里必须要知道函数体的rect
		caller := cpu.At(cpu.Runfunc.Cpu_x, cpu.Runfunc.Cpu_y).(*mirror.CallAtom)

		if caller.Value == nil {
			//创建 新执行现场
			funcrect := cpu.At(caller.Func.X, caller.Func.Y).(*mirror.FuncAtom) //方法区的地址

			NewFuncX, NewFuncY := cpu.Runfunc.Funcbody.X+cpu.Runfunc.Funcbody.Size_x, cpu.Runfunc.Cpu_y

			for i := 0; i < funcrect.Funcbody.Size_y; i++ {
				for j := 0; j < funcrect.Funcbody.Size_x; j++ {
					sourceatom := cpu.At(funcrect.Funcbody.X+j, funcrect.Funcbody.Y+i)
					cpu.Set(NewFuncX+j, NewFuncY+i, sourceatom)
				}
			}
			cpu.mem.Print()
			// 参数检查，传递参数

			if caller.Args != nil {
				Args := *caller.Args
				if Args.Isoffset {
					Args.X = cpu.Runfunc.Cpu_x + Args.X
					Args.Y = cpu.Runfunc.Cpu_y + Args.Y
					Args.Isoffset = false
				}

				//传参原则，不要求两个矩形完全一致func.args==args，但要求对应位置的类型一致
				for i := 0; i < Args.Size_y && i < funcrect.Args.Size_y; i++ {
					for j := 0; j < Args.Size_x && j < funcrect.Args.Size_x; j++ {
						targetatom := cpu.At(NewFuncX+funcrect.Args.X+j, NewFuncY+funcrect.Args.Y+i)
						sourceatom := cpu.At(Args.X+j, Args.Y+i)
						if targetatom != nil {
							//todo 类型检查
							cpu.Set(NewFuncX+funcrect.Args.X+j, NewFuncY+funcrect.Args.Y+i, sourceatom) //参数区是在函数体的x+1位置
						}

					}
				}
			}
			cpu.mem.Print()
			//写入新funcrect的funcbody的X,Y
			newfunc := cpu.At(NewFuncX, NewFuncY).(*mirror.FuncAtom)
			newfunc.Funcbody.X, newfunc.Funcbody.Y = NewFuncX, NewFuncY
			newfunc.CallerX = cpu.Runfunc.Cpu_x - cpu.Runfunc.Funcbody.Size_x
			newfunc.Cpu_x, newfunc.Cpu_y = NewFuncX, NewFuncY
			cpu.Set(NewFuncX, NewFuncY, newfunc)

			//移动至函数
			cpu.Runfunc = newfunc
		} else {
			cpu.Next()
		}

	case *mirror.ReturnAtom:

		caller := cpu.At(cpu.Runfunc.Funcbody.X+cpu.Runfunc.CallerX, cpu.Runfunc.Funcbody.Y).(*mirror.CallAtom)
		//存储函数的计算结果
		caller.Value = make([][]mirror.Atom, cpu.Runfunc.Value.Size_y)
		for i := 0; i < cpu.Runfunc.Value.Size_y; i++ {
			caller.Value[i] = make([]mirror.Atom, cpu.Runfunc.Value.Size_x)
			for j := 0; j < cpu.Runfunc.Value.Size_x; j++ {
				caller.Value[i][j] = cpu.InFunc(cpu.Runfunc.Value.X+j, cpu.Runfunc.Value.Y+i)
			}
		}
		//回收执行现场
		cpu.Set(cpu.Runfunc.Funcbody.X+cpu.Runfunc.CallerX, cpu.Runfunc.Funcbody.Y, caller)
		for i := 0; i < cpu.Runfunc.Funcbody.Size_y; i++ {
			for j := 0; j < cpu.Runfunc.Funcbody.Size_x; j++ {
				cpu.Set(cpu.Runfunc.Funcbody.X+j, cpu.Runfunc.Funcbody.Y+i, nil)
			}
		}
		//return

		cpu.Runfunc = cpu.At(caller.Myfunc.X, caller.Myfunc.Y).(*mirror.FuncAtom)
	}
}
