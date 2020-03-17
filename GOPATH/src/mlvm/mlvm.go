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
var cpu0 = vm.NewRunner(mem)

func test_0() {
	go func() {
		cpu0.Pause()
	}()
	time.Sleep(time.Second)
	test.TestExpression_null(mem, mirror.PointAtom{Point: mirror.Point{X: 0, Y: 0}})
	//vm.TestExpression_not(mem, mirror.PointAtom{Point: mirror.Point{X: 0, Y: 2}})
	//vm.TestExpression_goto(mem, mirror.PointAtom{Point: mirror.Point{X: 0, Y: 3}})

	fmt.Println()
	cpu0.Goon()
	time.Sleep(time.Second * 5)

}

func test_1() {
	go func() {
		cpu0.Pause()
	}()
	time.Sleep(time.Second)

	//vm.TestFunc(mem)
	//vm.TestCallfunc(mem)

	cpu0.Goon()
	time.Sleep(time.Second * 2)

	fmt.Scanln()

}

func main() {
	cpu0.Computecycle = time.Millisecond * 3001
	go cpu0.Do(0, 0)
	test_0()
	//go test_1()
	monitor.Run([]*vm.Runner{cpu0}, mem)
}
