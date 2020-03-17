package main

import (
	"fmt"
	"mirror"
	"mlvm/monitor"
	"mlvm/test"
	"mlvm/vm"
	"time"
)

var mem = vm.NewMemory()
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
	test.TestExpression_not(mem, mirror.Point{X: 4, Y: 4})
	//vm.TestExpression_goto(mem, mirror.PointAtom{Point: mirror.Point{X: 0, Y: 3}})
	cpu.Goon()
}

func test_1(cpu *vm.Runner) {
	go func() {
		cpu.Pause()
	}()
	time.Sleep(time.Second)

	//vm.TestFunc(mem)
	//vm.TestCallfunc(mem)

	cpu.Goon()
	time.Sleep(time.Second * 2)

	fmt.Scanln()

}

func main() {
	var cpu0 = vm.NewRunner(mem)
	var cpu1 = vm.NewRunner(mem)
	cpus = append(cpus, cpu0)
	cpus = append(cpus, cpu1)
	cpu0.Computecycle = time.Millisecond * 3001
	cpu1.Computecycle = time.Millisecond * 4001
	go cpu0.Do(0, 0)

	go test_nil(cpu0)

	go cpu1.Do(4, 4)
	go test_not(cpu1)
	//go test_1()
	monitor.Run(cpus, mem)
}
