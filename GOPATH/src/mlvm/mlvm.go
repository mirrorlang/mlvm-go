package main

import (
	"fmt"
	"mirror"
	"mlvm/monitor"
	"mlvm/test"
	"mlvm/vm"
	mem2 "mlvm/vm/mem"
	"time"
)

var mem = mem2.NewMemory()
var cpus []*vm.Runner

func test_nil(cpu *vm.Runner) {
	go func() {
		cpu.Pause()
	}()
	time.Sleep(time.Second)
	test.TestExpression_nil(mem, mirror.Point{X: 0, Y: 0})
	cpu.Goon()
}
func test_not(cpu *vm.Runner) {
	go func() {
		cpu.Pause()
	}()
	time.Sleep(time.Second)

	//vm.TestExpression_goto(mem, mirror.PointAtom{Point: mirror.Point{X: 0, Y: 3}})
	cpu.Goon()
}
func test_goto(cpu *vm.Runner) {
	go func() {
		cpu.Pause()
	}()
	time.Sleep(time.Second)
	test.TestExpression_not(mem, mirror.Point{X: 4, Y: 4})
	test.TestExpression_goto(mem, mirror.Point{X: 0, Y: 6})
	cpu.Goon()
}

func test_funccall(cpu *vm.Runner) {
	go func() {
		cpu.Pause()
	}()
	time.Sleep(time.Second)

	test.TestFunc(mem, mirror.Point{X: 0, Y: 10})
	test.TestCallfunc(mem, mirror.Point{X: 0, Y: 0})

	cpu.Goon()
	time.Sleep(time.Second * 2)

	fmt.Scanln()

}

func main() {
	var cpu0 = vm.NewRunner(mem)

	var cpu2 = vm.NewRunner(mem)
	cpus = append(cpus, cpu0)

	cpus = append(cpus, cpu2)
	cpu0.Computecycle = time.Millisecond * 3001

	cpu2.Computecycle = time.Millisecond * 3001
	go cpu0.Do(0, 0)
	go test_nil(cpu0)

	go cpu2.Do(0, 6)
	go test_goto(cpu2)
	monitor.Run(cpus, mem)
}
